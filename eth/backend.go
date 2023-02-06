// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package eth implements the Ethereum protocol.
package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juncachain/juncachain/accounts"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/consensus"
	"github.com/juncachain/juncachain/consensus/beacon"
	"github.com/juncachain/juncachain/consensus/clique"
	"github.com/juncachain/juncachain/consensus/posv"
	"github.com/juncachain/juncachain/contracts"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/core/bloombits"
	"github.com/juncachain/juncachain/core/rawdb"
	"github.com/juncachain/juncachain/core/state"
	"github.com/juncachain/juncachain/core/state/pruner"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/core/vm"
	"github.com/juncachain/juncachain/eth/downloader"
	"github.com/juncachain/juncachain/eth/ethconfig"
	"github.com/juncachain/juncachain/eth/gasprice"
	"github.com/juncachain/juncachain/eth/protocols/eth"
	"github.com/juncachain/juncachain/eth/protocols/snap"
	"github.com/juncachain/juncachain/ethclient"
	"github.com/juncachain/juncachain/ethdb"
	"github.com/juncachain/juncachain/event"
	"github.com/juncachain/juncachain/internal/ethapi"
	"github.com/juncachain/juncachain/internal/shutdowncheck"
	"github.com/juncachain/juncachain/log"
	"github.com/juncachain/juncachain/miner"
	"github.com/juncachain/juncachain/node"
	"github.com/juncachain/juncachain/p2p"
	"github.com/juncachain/juncachain/p2p/dnsdisc"
	"github.com/juncachain/juncachain/p2p/enode"
	"github.com/juncachain/juncachain/params"
	"github.com/juncachain/juncachain/rlp"
	"github.com/juncachain/juncachain/rpc"
	"github.com/pkg/errors"
	"github.com/umpc/go-sortedmap"
	"math/big"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Config contains the configuration options of the ETH protocol.
// Deprecated: use ethconfig.Config instead.
type Config = ethconfig.Config

// Ethereum implements the Ethereum full node service.
type Ethereum struct {
	config *ethconfig.Config

	// Handlers
	txPool             *core.TxPool
	blockchain         *core.BlockChain
	handler            *handler
	ethDialCandidates  enode.Iterator
	snapDialCandidates enode.Iterator
	merger             *consensus.Merger

	// DB interfaces
	chainDb ethdb.Database // Block chain database

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests     chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer      *core.ChainIndexer             // Bloom indexer operating during block imports
	closeBloomHandler chan struct{}

	APIBackend *EthAPIBackend

	miner     *miner.Miner
	gasPrice  *big.Int
	etherbase common.Address

	networkID     uint64
	netRPCService *ethapi.NetAPI

	p2pServer *p2p.Server

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)

	shutdownTracker *shutdowncheck.ShutdownTracker // Tracks if and when the node has shutdown ungracefully

	IPCEndpoint string
	once        sync.Once
	client      *ethclient.Client // Global ipc client instance.
}

// New creates a new Ethereum object (including the
// initialisation of the common Ethereum object)
func New(stack *node.Node, config *ethconfig.Config) (*Ethereum, error) {
	// Ensure configuration values are compatible and sane
	if config.SyncMode == downloader.LightSync {
		return nil, errors.New("can't run eth.Ethereum in light sync mode, use les.LightEthereum")
	}
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	if config.Miner.GasPrice == nil || config.Miner.GasPrice.Cmp(common.Big0) <= 0 {
		log.Warn("Sanitizing invalid miner gas price", "provided", config.Miner.GasPrice, "updated", ethconfig.Defaults.Miner.GasPrice)
		config.Miner.GasPrice = new(big.Int).Set(ethconfig.Defaults.Miner.GasPrice)
	}
	if config.NoPruning && config.TrieDirtyCache > 0 {
		if config.SnapshotCache > 0 {
			config.TrieCleanCache += config.TrieDirtyCache * 3 / 5
			config.SnapshotCache += config.TrieDirtyCache * 2 / 5
		} else {
			config.TrieCleanCache += config.TrieDirtyCache
		}
		config.TrieDirtyCache = 0
	}
	log.Info("Allocated trie memory caches", "clean", common.StorageSize(config.TrieCleanCache)*1024*1024, "dirty", common.StorageSize(config.TrieDirtyCache)*1024*1024)

	// Transfer mining-related config to the ethash config.
	ethashConfig := config.Ethash
	ethashConfig.NotifyFull = config.Miner.NotifyFull

	// Assemble the Ethereum object
	chainDb, err := stack.OpenDatabaseWithFreezer("chaindata", config.DatabaseCache, config.DatabaseHandles, config.DatabaseFreezer, "eth/db/chaindata/", false)
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlockWithOverride(chainDb, config.Genesis, config.OverrideTerminalTotalDifficulty, config.OverrideTerminalTotalDifficultyPassed)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("")
	log.Info(strings.Repeat("-", 153))
	for _, line := range strings.Split(chainConfig.String(), "\n") {
		log.Info(line)
	}
	log.Info(strings.Repeat("-", 153))
	log.Info("")

	if err := pruner.RecoverPruning(stack.ResolvePath(""), chainDb, stack.ResolvePath(config.TrieCleanCacheJournal)); err != nil {
		log.Error("Failed to recover state", "error", err)
	}
	merger := consensus.NewMerger(chainDb)
	eth := &Ethereum{
		config:            config,
		merger:            merger,
		chainDb:           chainDb,
		eventMux:          stack.EventMux(),
		accountManager:    stack.AccountManager(),
		engine:            ethconfig.CreateConsensusEngine(stack, chainConfig, &ethashConfig, config.Miner.Notify, config.Miner.Noverify, chainDb),
		closeBloomHandler: make(chan struct{}),
		networkID:         config.NetworkId,
		gasPrice:          config.Miner.GasPrice,
		etherbase:         config.Miner.Etherbase,
		bloomRequests:     make(chan chan *bloombits.Retrieval),
		bloomIndexer:      core.NewBloomIndexer(chainDb, params.BloomBitsBlocks, params.BloomConfirms),
		p2pServer:         stack.Server(),
		shutdownTracker:   shutdowncheck.NewShutdownTracker(chainDb),
		IPCEndpoint:       stack.IPCEndpoint(),
	}

	bcVersion := rawdb.ReadDatabaseVersion(chainDb)
	var dbVer = "<nil>"
	if bcVersion != nil {
		dbVer = fmt.Sprintf("%d", *bcVersion)
	}
	log.Info("Initialising Ethereum protocol", "network", config.NetworkId, "dbversion", dbVer)

	if !config.SkipBcVersionCheck {
		if bcVersion != nil && *bcVersion > core.BlockChainVersion {
			return nil, fmt.Errorf("database version is v%d, Geth %s only supports v%d", *bcVersion, params.VersionWithMeta, core.BlockChainVersion)
		} else if bcVersion == nil || *bcVersion < core.BlockChainVersion {
			if bcVersion != nil { // only print warning on upgrade, not on init
				log.Warn("Upgrade blockchain database version", "from", dbVer, "to", core.BlockChainVersion)
			}
			rawdb.WriteDatabaseVersion(chainDb, core.BlockChainVersion)
		}
	}
	var (
		vmConfig = vm.Config{
			EnablePreimageRecording: config.EnablePreimageRecording,
		}
		cacheConfig = &core.CacheConfig{
			TrieCleanLimit:      config.TrieCleanCache,
			TrieCleanJournal:    stack.ResolvePath(config.TrieCleanCacheJournal),
			TrieCleanRejournal:  config.TrieCleanCacheRejournal,
			TrieCleanNoPrefetch: config.NoPrefetch,
			TrieDirtyLimit:      config.TrieDirtyCache,
			TrieDirtyDisabled:   config.NoPruning,
			TrieTimeLimit:       config.TrieTimeout,
			SnapshotLimit:       config.SnapshotCache,
			Preimages:           config.Preimages,
		}
	)
	eth.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, chainConfig, eth.engine, vmConfig, eth.shouldPreserve, &config.TxLookupLimit)
	if err != nil {
		return nil, err
	}
	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		eth.blockchain.SetHead(compat.RewindTo)
		rawdb.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}
	eth.bloomIndexer.Start(eth.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = stack.ResolvePath(config.TxPool.Journal)
	}
	eth.txPool = core.NewTxPool(config.TxPool, chainConfig, eth.blockchain)

	// Permit the downloader to use the trie cache allowance during fast sync
	cacheLimit := cacheConfig.TrieCleanLimit + cacheConfig.TrieDirtyLimit + cacheConfig.SnapshotLimit
	checkpoint := config.Checkpoint
	if checkpoint == nil {
		checkpoint = params.TrustedCheckpoints[genesisHash]
	}
	if eth.handler, err = newHandler(&handlerConfig{
		Database:       chainDb,
		Chain:          eth.blockchain,
		TxPool:         eth.txPool,
		Merger:         merger,
		Network:        config.NetworkId,
		Sync:           config.SyncMode,
		BloomCache:     uint64(cacheLimit),
		EventMux:       eth.eventMux,
		Checkpoint:     checkpoint,
		RequiredBlocks: config.RequiredBlocks,
	}); err != nil {
		return nil, err
	}
	// for PoSV
	if chainConfig.Posv != nil {
		var c *posv.PoSV
		if cl, ok := eth.engine.(*beacon.Beacon); ok {
			if e, ok := cl.InnerEngine().(*posv.PoSV); ok {
				c = e
			}
		}

		eth.TxPool().IsSigner = func(address common.Address) bool {
			currentHeader := eth.blockchain.CurrentHeader()
			header := currentHeader
			// Sometimes, the latest block hasn't been inserted to chain yet
			// getSnapshot from parent block if it exists
			parentHeader := eth.blockchain.GetHeader(currentHeader.ParentHash, currentHeader.Number.Uint64()-1)
			if parentHeader != nil {
				// not genesis block
				header = parentHeader
			}

			lastCheckpoint := c.LastEpoch(header.Number.Uint64())
			var extra posv.Extra
			if err := extra.FromBytes(eth.blockchain.GetHeaderByNumber(lastCheckpoint).Extra); err != nil {
				return false
			}

			for _, v := range extra.Epoch.M1Slice() {
				if v == address {
					return true
				}
			}
			return false
		}

		c.HookValidator = func(number *big.Int, masters []common.Address) ([]common.Address, error) {
			start := time.Now()
			client, err := eth.GetClient()
			if err != nil {
				return nil, err
			}
			validators, err := contracts.GetValidators(client, masters)
			if err != nil {
				return nil, err
			}
			log.Debug("Time Calculated HookValidator ", "block", number.Uint64(), "time", common.PrettyDuration(time.Since(start)))
			return validators, nil
		}

		c.HookGetMasterNodes = func(number *big.Int) (posv.MasterNodes, error) {
			client, err := eth.GetClient()
			if err != nil {
				return nil, err
			}

			candidates, err := contracts.GetCandidates(client, number)
			if err != nil {
				return nil, err
			}

			var (
				masternodes posv.MasterNodes
			)
			for k, v := range candidates {
				masternodes = append(masternodes, posv.MasterNode{
					Address: k,
					Stake:   v,
				})
			}
			sort.Sort(masternodes)
			return masternodes, nil
		}

		c.HookBlockSign = func(signer common.Address, toSign, current *types.Header) error {
			if err := contracts.CreateTransactionSign(chainConfig, eth.txPool, eth.accountManager, signer, toSign, current); err != nil {
				return errors.Errorf("Fail to create tx block sign for importing block: %v", err)
			}
			return nil
		}

		c.HookGetBlockSigners = func(chain consensus.ChainHeaderReader, stateBlock *state.StateDB, header *types.Header) (map[common.Address]int, error) {
			if header.Number.Uint64()%c.Config().Epoch != 0 {
				return nil, nil
			}
			begin := c.LastEpoch(header.Number.Uint64())
			end := header.Number.Uint64() - 1

			signCount := make(map[common.Address]int)
			for i := begin; i <= end; i++ {
				hd := chain.GetHeaderByNumber(i)
				if hd != nil {
					signers, err := contracts.GetSignersFromContract(stateBlock, hd.Hash())
					if err != nil {
						log.Crit("HookQueryBlockSigners", "err", err)
					}
					for _, v := range signers {
						signCount[v] = signCount[v] + 1
					}
				}
			}

			return signCount, nil
		}

		c.HookSetRandomizeSecret = func(signer common.Address, header *types.Header) error {
			checkNumber := header.Number.Uint64() % c.Config().Epoch
			if c.Config().Epoch-checkNumber != common.EpochBlockSecret {
				return nil
			}
			log.Info("[PoSV]Is check number to set secret", "number", header.Number.Uint64())
			if err := contracts.CreateTransactionSetSecret(chainConfig, eth.txPool, eth.accountManager, header, chainDb, signer); err != nil {
				return errors.Errorf("Fail to create tx set secret for importing block: %v", err)
			}
			return nil
		}

		c.HookSetRandomizeOpening = func(signer common.Address, header *types.Header) error {
			checkNumber := header.Number.Uint64() % c.Config().Epoch
			if c.Config().Epoch-checkNumber != common.EpochBlockOpening {
				return nil
			}
			log.Info("[PoSV]Is check number to set opening", "number", header.Number.Uint64())
			if err := contracts.CreateTransactionSetOpening(chainConfig, eth.txPool, eth.accountManager, header, chainDb, signer); err != nil {
				return errors.Errorf("Fail to create tx set opening for importing block: %v", err)
			}
			return nil
		}

		c.HookReward = func(chain consensus.ChainHeaderReader, stateBlock *state.StateDB, header *types.Header) (map[string]interface{}, error) {
			number := header.Number.Uint64()
			if number == 0 || number%c.Config().Epoch != 0 {
				return nil, nil
			}
			type Stat struct {
				Address common.Address `json:"address"`
				Cap     *big.Int       `json:"cap"`
				Reward  *big.Int       `json:"reward"`
			}
			var (
				epoch          = make(map[string]interface{})
				rewards        = make(map[string]interface{})
				m1Stat         = make(map[common.Address]*Stat)
				m2Stat         = make(map[common.Address]*Stat)
				voterStat      = make(map[common.Address]*Stat)
				totalCap       = new(big.Int)
				totalSignCount int
			)

			// list sealers as M1(valid M1 list)
			lastNumber := number - c.Config().Epoch
			if lastNumber == 0 {
				lastNumber = 1
			}
			var sealers = make(map[common.Address]int)
			for i := number - 1; i >= lastNumber; i-- {
				hd := chain.GetHeaderByNumber(i)
				if hd == nil {
					return nil, errors.Errorf("header %d is nil", lastNumber+i)
				}
				sealer, err := c.Author(hd)
				if err != nil {
					return nil, err
				}
				sealers[sealer] = sealers[sealer] + 1
			}
			rewards["sealers"] = sealers

			// list block signers as M2(all M2)
			m2s, err := c.HookGetBlockSigners(chain, stateBlock, header)
			if err != nil {
				return nil, err
			}
			rewards["signmiss"] = m2s

			for sealer, _ := range sealers {
				m1Stat[sealer] = &Stat{Address: sealer, Cap: new(big.Int), Reward: new(big.Int)}
				voters := contracts.GetVotersFromState(stateBlock, sealer)
				for _, v := range voters {
					if voterStat[v] == nil {
						voterStat[v] = &Stat{Address: v, Cap: new(big.Int), Reward: new(big.Int)}
					}
					cap := contracts.GetVoterCapFromState(stateBlock, sealer, v)
					totalCap = new(big.Int).Add(totalCap, cap)
					if st, ok := voterStat[v]; ok {
						st.Cap = new(big.Int).Add(st.Cap, cap)
					}
					if st, ok := m1Stat[sealer]; ok {
						st.Cap = new(big.Int).Add(st.Cap, cap)
					}
				}
			}

			for signer, signCount := range m2s {
				totalSignCount = totalSignCount + signCount
				if st, ok := m2Stat[signer]; ok {
					st.Cap = new(big.Int).Add(st.Cap, big.NewInt(int64(signCount)))
				} else {
					m2Stat[signer] = &Stat{Address: signer, Cap: big.NewInt(int64(signCount)), Reward: new(big.Int)}
				}
			}

			rewardM1 := new(big.Int).Mul(c.Config().Reward, new(big.Int).SetInt64(common.RewardM1Percent))
			rewardM1 = new(big.Int).Div(rewardM1, new(big.Int).SetInt64(100))
			for _, v := range m1Stat {
				v.Reward = new(big.Int).Div(new(big.Int).Mul(rewardM1, v.Cap), totalCap)
				if v.Reward.Cmp(new(big.Int)) > 0 {
					stateBlock.AddBalance(v.Address, v.Reward)
				}
			}
			rewards["m1"] = m1Stat

			rewardM2 := new(big.Int).Mul(c.Config().Reward, new(big.Int).SetInt64(common.RewardM2Percent))
			rewardM2 = new(big.Int).Div(rewardM2, new(big.Int).SetInt64(100))
			for _, v := range m2Stat {
				v.Reward = new(big.Int).Div(new(big.Int).Mul(rewardM2, v.Cap), new(big.Int).SetInt64(int64(totalSignCount)))
				if v.Reward.Cmp(new(big.Int)) > 0 {
					stateBlock.AddBalance(v.Address, v.Reward)
				}
			}
			rewards["m2"] = m2Stat

			rewardVoter := new(big.Int).Mul(c.Config().Reward, new(big.Int).SetInt64(common.RewardVoterPercent))
			rewardVoter = new(big.Int).Div(rewardVoter, new(big.Int).SetInt64(100))
			for _, v := range voterStat {
				v.Reward = new(big.Int).Div(new(big.Int).Mul(rewardVoter, v.Cap), totalCap)
				if v.Reward.Cmp(new(big.Int)) > 0 {
					stateBlock.AddBalance(v.Address, v.Reward)
				}
			}
			rewards["voters"] = voterStat

			rewardFoundation := new(big.Int).Mul(c.Config().Reward, new(big.Int).SetInt64(common.RewardFoundationPercent))
			rewardFoundation = new(big.Int).Div(rewardFoundation, new(big.Int).SetInt64(100))
			rewards["foundation"] = struct {
				Address common.Address `json:"address"`
				Reward  *big.Int       `json:"reward"`
			}{c.Config().Foundation, rewardFoundation}
			if rewardFoundation.Cmp(new(big.Int)) > 0 {
				stateBlock.AddBalance(c.Config().Foundation, rewardFoundation)
			}

			var extra posv.Extra
			if err := extra.FromBytes(header.Extra); err != nil {
				return nil, err
			}
			epoch["epoch"] = extra.Epoch
			epoch["rewards"] = rewards
			bz, _ := json.Marshal(epoch)
			_ = chainDb.Put([]byte(fmt.Sprintf("%s-%d", common.EpochKeyPrefix, number/c.Config().Epoch)), bz)
			return epoch, nil
		}

		c.HookPenalty = func(chain consensus.ChainHeaderReader, header *types.Header) ([]common.Address, error) {
			number := header.Number.Uint64()
			if number == 0 || number%c.Config().Epoch != 0 {
				return nil, nil
			}
			lastNumber := number - c.Config().Epoch
			if lastNumber == 0 {
				lastNumber = 1
			}
			epoch, err := c.GetEpoch(chain, c.LastEpoch(number))
			if err != nil {
				return nil, err
			}

			// Count M1 miss seal block by turn
			sealermiss := sortedmap.New(0, func(i, j interface{}) bool {
				if i.(int) > j.(int) {
					return true
				}
				return false
			})
			for i := number - 1; i >= lastNumber; i-- {
				hd := chain.GetHeaderByNumber(i)
				if hd == nil {
					return nil, errors.Errorf("header %d is nil", i)
				}
				signer, err := c.Author(hd)
				if err != nil {
					return nil, err
				}
				turn := epoch.M1(c.Config().Epoch, i)
				if signer != turn {
					if v, ok := sealermiss.Get(turn); !ok {
						sealermiss.Insert(turn, 1)
					} else {
						sealermiss.Replace(turn, v.(int)+1)
					}
				}
			}

			var penalties = make([]common.Address, 0)
			if it, err := sealermiss.IterCh(); err == nil {
				for rec := range it.Records() {
					if rec.Val.(int) > common.EpochBlockSealMissAllow {
						penalties = append(penalties, rec.Key.(common.Address))
					}
				}
			}
			return penalties, nil
		}
	}

	eth.miner = miner.New(eth, &config.Miner, chainConfig, eth.EventMux(), eth.engine, eth.isLocalBlock)
	eth.miner.SetExtra(makeExtraData(config.Miner.ExtraData))

	eth.APIBackend = &EthAPIBackend{stack.Config().ExtRPCEnabled(), stack.Config().AllowUnprotectedTxs, eth, nil}
	if eth.APIBackend.allowUnprotectedTxs {
		log.Info("Unprotected transactions allowed")
	}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.Miner.GasPrice
	}
	eth.APIBackend.gpo = gasprice.NewOracle(eth.APIBackend, gpoParams)

	// Setup DNS discovery iterators.
	dnsclient := dnsdisc.NewClient(dnsdisc.Config{})
	eth.ethDialCandidates, err = dnsclient.NewIterator(eth.config.EthDiscoveryURLs...)
	if err != nil {
		return nil, err
	}
	eth.snapDialCandidates, err = dnsclient.NewIterator(eth.config.SnapDiscoveryURLs...)
	if err != nil {
		return nil, err
	}

	// Start the RPC service
	eth.netRPCService = ethapi.NewNetAPI(eth.p2pServer, config.NetworkId)

	// Register the backend on the node
	stack.RegisterAPIs(eth.APIs())
	stack.RegisterProtocols(eth.Protocols())
	stack.RegisterLifecycle(eth)

	// Successful startup; push a marker and check previous unclean shutdowns.
	eth.shutdownTracker.MarkStartup()

	return eth, nil
}

func makeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		// create default extradata
		extra, _ = rlp.EncodeToBytes([]interface{}{
			uint(params.VersionMajor<<16 | params.VersionMinor<<8 | params.VersionPatch),
			"geth",
			runtime.Version(),
			runtime.GOOS,
		})
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", params.MaximumExtraDataSize)
		extra = nil
	}
	return extra
}

// APIs return the collection of RPC services the ethereum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Ethereum) APIs() []rpc.API {
	apis := ethapi.GetAPIs(s.APIBackend)

	// Append any APIs exposed explicitly by the consensus engine
	apis = append(apis, s.engine.APIs(s.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "eth",
			Service:   NewEthereumAPI(s),
		}, {
			Namespace: "miner",
			Service:   NewMinerAPI(s),
		}, {
			Namespace: "eth",
			Service:   downloader.NewDownloaderAPI(s.handler.downloader, s.eventMux),
		}, {
			Namespace: "admin",
			Service:   NewAdminAPI(s),
		}, {
			Namespace: "debug",
			Service:   NewDebugAPI(s),
		}, {
			Namespace: "net",
			Service:   s.netRPCService,
		},
	}...)
}

func (s *Ethereum) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Ethereum) Etherbase() (eb common.Address, err error) {
	s.lock.RLock()
	etherbase := s.etherbase
	s.lock.RUnlock()

	if etherbase != (common.Address{}) {
		return etherbase, nil
	}
	if wallets := s.AccountManager().Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			etherbase := accounts[0].Address

			s.lock.Lock()
			s.etherbase = etherbase
			s.lock.Unlock()

			log.Info("Etherbase automatically configured", "address", etherbase)
			return etherbase, nil
		}
	}
	return common.Address{}, fmt.Errorf("etherbase must be explicitly specified")
}

// isLocalBlock checks whether the specified block is mined
// by local miner accounts.
//
// We regard two types of accounts as local miner account: etherbase
// and accounts specified via `txpool.locals` flag.
func (s *Ethereum) isLocalBlock(header *types.Header) bool {
	author, err := s.engine.Author(header)
	if err != nil {
		log.Warn("Failed to retrieve block author", "number", header.Number.Uint64(), "hash", header.Hash(), "err", err)
		return false
	}
	// Check whether the given address is etherbase.
	s.lock.RLock()
	etherbase := s.etherbase
	s.lock.RUnlock()
	if author == etherbase {
		return true
	}
	// Check whether the given address is specified by `txpool.local`
	// CLI flag.
	for _, account := range s.config.TxPool.Locals {
		if account == author {
			return true
		}
	}
	return false
}

// shouldPreserve checks whether we should preserve the given block
// during the chain reorg depending on whether the author of block
// is a local account.
func (s *Ethereum) shouldPreserve(header *types.Header) bool {
	// The reason we need to disable the self-reorg preserving for clique
	// is it can be probable to introduce a deadlock.
	//
	// e.g. If there are 7 available signers
	//
	// r1   A
	// r2     B
	// r3       C
	// r4         D
	// r5   A      [X] F G
	// r6    [X]
	//
	// In the round5, the inturn signer E is offline, so the worst case
	// is A, F and G sign the block of round5 and reject the block of opponents
	// and in the round6, the last available signer B is offline, the whole
	// network is stuck.
	if _, ok := s.engine.(*clique.Clique); ok {
		return false
	}
	return s.isLocalBlock(header)
}

// SetEtherbase sets the mining reward address.
func (s *Ethereum) SetEtherbase(etherbase common.Address) {
	s.lock.Lock()
	s.etherbase = etherbase
	s.lock.Unlock()

	s.miner.SetEtherbase(etherbase)
}

// StartMining starts the miner with the given number of CPU threads. If mining
// is already running, this method adjust the number of threads allowed to use
// and updates the minimum price required by the transaction pool.
func (s *Ethereum) StartMining(threads int) error {
	// Update the thread count within the consensus engine
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := s.engine.(threaded); ok {
		log.Info("Updated mining threads", "threads", threads)
		if threads == 0 {
			threads = -1 // Disable the miner from within
		}
		th.SetThreads(threads)
	}
	// If the miner was not running, initialize it
	if !s.IsMining() {
		// Propagate the initial price point to the transaction pool
		s.lock.RLock()
		price := s.gasPrice
		s.lock.RUnlock()
		s.txPool.SetGasPrice(price)

		// Configure the local mining address
		eb, err := s.Etherbase()
		if err != nil {
			log.Error("Cannot start mining without etherbase", "err", err)
			return fmt.Errorf("etherbase missing: %v", err)
		}

		getWallet := func() (accounts.Wallet, error) {
			wallet, err := s.accountManager.Find(accounts.Account{Address: eb})
			if wallet == nil || err != nil {
				log.Error("Etherbase account unavailable locally", "err", err)
				return wallet, fmt.Errorf("signer missing: %v", err)
			}
			return wallet, nil
		}

		if c, ok := s.engine.(*clique.Clique); ok {
			w, err := getWallet()
			if err != nil {
				return fmt.Errorf("signer missing: %v", err)
			}
			c.Authorize(eb, w.SignData)
		} else if cl, ok := s.engine.(*beacon.Beacon); ok {
			if c, ok := cl.InnerEngine().(*clique.Clique); ok {
				w, err := getWallet()
				if err != nil {
					return fmt.Errorf("signer missing: %v", err)
				}
				c.Authorize(eb, w.SignData)
			} else if c, ok := cl.InnerEngine().(*posv.PoSV); ok {
				w, err := getWallet()
				if err != nil {
					return fmt.Errorf("signer missing: %v", err)
				}
				c.Authorize(eb, w.SignData)
			}
		}
		// If mining is started, we can disable the transaction rejection mechanism
		// introduced to speed sync times.
		atomic.StoreUint32(&s.handler.acceptTxs, 1)

		go s.miner.Start(eb)
	}
	return nil
}

// StopMining terminates the miner, both at the consensus engine level as well as
// at the block creation level.
func (s *Ethereum) StopMining() {
	// Update the thread count within the consensus engine
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := s.engine.(threaded); ok {
		th.SetThreads(-1)
	}
	// Stop the block creating itself
	s.miner.Stop()
}

func (s *Ethereum) IsMining() bool      { return s.miner.Mining() }
func (s *Ethereum) Miner() *miner.Miner { return s.miner }

func (s *Ethereum) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Ethereum) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Ethereum) TxPool() *core.TxPool               { return s.txPool }
func (s *Ethereum) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Ethereum) Engine() consensus.Engine           { return s.engine }
func (s *Ethereum) ChainDb() ethdb.Database            { return s.chainDb }
func (s *Ethereum) IsListening() bool                  { return true } // Always listening
func (s *Ethereum) Downloader() *downloader.Downloader { return s.handler.downloader }
func (s *Ethereum) Synced() bool                       { return atomic.LoadUint32(&s.handler.acceptTxs) == 1 }
func (s *Ethereum) SetSynced()                         { atomic.StoreUint32(&s.handler.acceptTxs, 1) }
func (s *Ethereum) ArchiveMode() bool                  { return s.config.NoPruning }
func (s *Ethereum) BloomIndexer() *core.ChainIndexer   { return s.bloomIndexer }
func (s *Ethereum) Merger() *consensus.Merger          { return s.merger }
func (s *Ethereum) SyncMode() downloader.SyncMode {
	mode, _ := s.handler.chainSync.modeAndLocalHead()
	return mode
}

// Protocols returns all the currently configured
// network protocols to start.
func (s *Ethereum) Protocols() []p2p.Protocol {
	protos := eth.MakeProtocols((*ethHandler)(s.handler), s.networkID, s.ethDialCandidates)
	if s.config.SnapshotCache > 0 {
		protos = append(protos, snap.MakeProtocols((*snapHandler)(s.handler), s.snapDialCandidates)...)
	}
	return protos
}

// Start implements node.Lifecycle, starting all internal goroutines needed by the
// Ethereum protocol implementation.
func (s *Ethereum) Start() error {
	eth.StartENRUpdater(s.blockchain, s.p2pServer.LocalNode())

	// Start the bloom bits servicing goroutines
	s.startBloomHandlers(params.BloomBitsBlocks)

	// Regularly update shutdown marker
	s.shutdownTracker.Start()

	// Figure out a max peers count based on the server limits
	maxPeers := s.p2pServer.MaxPeers
	if s.config.LightServ > 0 {
		if s.config.LightPeers >= s.p2pServer.MaxPeers {
			return fmt.Errorf("invalid peer config: light peer count (%d) >= total peer count (%d)", s.config.LightPeers, s.p2pServer.MaxPeers)
		}
		maxPeers -= s.config.LightPeers
	}
	// Start the networking layer and the light server if requested
	s.handler.Start(maxPeers)
	return nil
}

// Stop implements node.Lifecycle, terminating all internal goroutines used by the
// Ethereum protocol.
func (s *Ethereum) Stop() error {
	// Stop all the peer-related stuff first.
	s.ethDialCandidates.Close()
	s.snapDialCandidates.Close()
	s.handler.Stop()

	// Then stop everything else.
	s.bloomIndexer.Close()
	close(s.closeBloomHandler)
	s.txPool.Stop()
	s.miner.Close()
	s.blockchain.Stop()
	s.engine.Close()

	// Clean shutdown marker as the last thing before closing db
	s.shutdownTracker.Stop()

	s.chainDb.Close()
	s.eventMux.Stop()

	return nil
}

func (s *Ethereum) GetClient() (*ethclient.Client, error) {
	s.once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				log.Crit("PoSV init ethclient timeout")
			default:
				cli, err := ethclient.Dial(s.IPCEndpoint)
				if err != nil {
					log.Warn("PoSV ethclient.Dial", "err", err)
					continue
				}
				s.client = cli
				return
			}
		}
	})
	return s.client, nil
}
