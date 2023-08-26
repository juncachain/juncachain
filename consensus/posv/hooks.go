package posv

import (
	"encoding/json"
	"fmt"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/consensus"
	"github.com/juncachain/juncachain/contracts"
	"github.com/juncachain/juncachain/core/state"
	"github.com/juncachain/juncachain/core/types"
	"github.com/pkg/errors"
	"github.com/umpc/go-sortedmap"
	"math/big"
)

// HookGetCandidates
// Read the candidate list and candidate votes from the contract,
// sort and return according to the rules;
// sorting rules: the one with the larger votes is first, and the
// candidate with the smaller address is the first when the votes are the same
func (c *PoSV) HookGetCandidates(state *state.StateDB) MasterNodes {
	var ms MasterNodes
	candidates := contracts.GetCandidatesFromState(state)
	for _, v := range candidates {
		cap := contracts.GetCandidateCapFromState(state, v)
		ms = append(ms, MasterNode{
			Address: v,
			Stake:   cap,
		})
	}
	return ms
}

func (c *PoSV) HookPenalty(chain consensus.ChainHeaderReader, number uint64) ([]common.Address, error) {
	lastCheckpoint := c.LastCheckpoint(number)
	lastEpoch, err := c.GetEpoch(chain, lastCheckpoint)
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
	for i := lastCheckpoint + 1; i < number; i++ {
		hd := chain.GetHeaderByNumber(i)
		if hd != nil {
			signer, err := c.Author(hd)
			if err != nil {
				return nil, err
			}
			turn := lastEpoch.M1(c.Config().Epoch, i)
			if signer != turn {
				if v, ok := sealermiss.Get(turn); !ok {
					sealermiss.Insert(turn, 1)
				} else {
					sealermiss.Replace(turn, v.(int)+1)
				}
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

func (c *PoSV) HookRandomizeSigners(statedb *state.StateDB, masternodes []common.Address) ([]common.Address, error) {
	var randomizes []int64
	for _, addr := range masternodes {
		random, err := contracts.GetRandomizeFromStateDB(statedb, addr)
		if err != nil {
			return nil, err
		}
		randomizes = append(randomizes, random)
	}
	// Get randomize m2 list.
	return contracts.GenerateM2FromRandomize(randomizes, masternodes)
}

func (c *PoSV) HookReward(chain consensus.ChainHeaderReader, stateBlock *state.StateDB, header *types.Header, nextEpoch *Epoch) (map[string]interface{}, error) {
	number := header.Number.Uint64()
	if number == 0 || number%c.Config().Epoch != 0 {
		return nil, nil
	}
	type Stat struct {
		Address common.Address `json:"address"`
		Owner   common.Address `json:"owner"` // for candidate
		Cap     *big.Int       `json:"cap"`
		Reward  *big.Int       `json:"reward"`
	}
	var (
		epochStat         = make(map[string]interface{})
		rewards           = make(map[string]interface{})
		m1Stat            = make(map[common.Address]*Stat)
		m2Stat            = make(map[common.Address]*Stat)
		voterStat         = make(map[common.Address]*Stat)
		totalCap          = new(big.Int)
		totalSignedBlocks int
		totalSealedBlocks int
	)

	// Get last Epoch
	lastCheckpoint := c.LastCheckpoint(number)
	lastEpoch, err := c.GetEpoch(chain, lastCheckpoint)
	if err != nil {
		return nil, errors.Errorf("GetEpoch %d err %v", lastCheckpoint, err)
	}

	// List sealers as M1(valid M1 list)
	var sealers = make(map[common.Address]int)
	for i := lastCheckpoint + 1; i < number; i++ {
		hd := chain.GetHeaderByNumber(i)
		if hd != nil {
			sealer, err := c.Author(hd)
			if err != nil {
				return nil, err
			}
			sealers[sealer] = sealers[sealer] + 1
		}
	}

	// List block signers as M2(all M2)
	blockSigners, err := c.HookGetBlockSigners(chain, stateBlock, header)
	if err != nil {
		return nil, err
	}

	// if the block signer isn't M2,remove it
	for m2, _ := range blockSigners {
		if !lastEpoch.IsM2(m2) {
			delete(blockSigners, m2)
		}
	}

	// Generate m1 stat and voters stat
	for sealer, count := range sealers {
		totalSealedBlocks += count
		owner := contracts.GetCandidateOwnerFromState(stateBlock, sealer)
		m1Stat[sealer] = &Stat{Address: sealer, Owner: owner, Cap: new(big.Int).SetInt64(int64(count)), Reward: new(big.Int)}
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
		}
	}

	// Generate m2 stat
	for signer, signCount := range blockSigners {
		totalSignedBlocks = totalSignedBlocks + signCount
		if st, ok := m2Stat[signer]; ok {
			st.Cap = new(big.Int).Add(st.Cap, big.NewInt(int64(signCount)))
		} else {
			owner := contracts.GetCandidateOwnerFromState(stateBlock, signer)
			m2Stat[signer] = &Stat{Address: signer, Owner: owner, Cap: big.NewInt(int64(signCount)), Reward: new(big.Int)}
		}
	}

	// Calc current epoch rewards
	currentEpochRewards := lastEpoch.Reward
	if currentEpochRewards.Cmp(big.NewInt(0)) == 0 {
		return epochStat, nil
	}
	rewards["total"] = currentEpochRewards

	// Calc and rewards to m1s
	rewardM1 := new(big.Int).Mul(currentEpochRewards, new(big.Int).SetInt64(common.RewardM1Percent))
	rewardM1 = new(big.Int).Div(rewardM1, new(big.Int).SetInt64(100))
	for _, v := range m1Stat {
		v.Reward = new(big.Int).Div(new(big.Int).Mul(rewardM1, v.Cap), new(big.Int).SetUint64(uint64(totalSealedBlocks)))
		if v.Reward.Cmp(new(big.Int)) > 0 {
			if chain.Config().ChainID.Cmp(big.NewInt(667)) == 0 || chain.Config().IsJuncatFix1(header.Number) {
				stateBlock.AddBalance(v.Owner, v.Reward)
			} else {
				stateBlock.AddBalance(v.Owner, v.Reward)
			}
		}
	}
	rewards["sealers"] = m1Stat

	// Calc and rewards to blockSigners
	rewardM2 := new(big.Int).Mul(currentEpochRewards, new(big.Int).SetInt64(common.RewardM2Percent))
	rewardM2 = new(big.Int).Div(rewardM2, new(big.Int).SetInt64(100))
	for _, v := range m2Stat {
		v.Reward = new(big.Int).Div(new(big.Int).Mul(rewardM2, v.Cap), new(big.Int).SetInt64(int64(totalSignedBlocks)))
		if v.Reward.Cmp(new(big.Int)) > 0 {
			if chain.Config().ChainID.Cmp(big.NewInt(667)) == 0 || chain.Config().IsJuncatFix1(header.Number) {
				stateBlock.AddBalance(v.Owner, v.Reward)
			} else {
				stateBlock.AddBalance(v.Owner, v.Reward)
			}
		}
	}
	rewards["signers"] = m2Stat

	// Calc and rewards to voters
	rewardVoter := new(big.Int).Mul(currentEpochRewards, new(big.Int).SetInt64(common.RewardVoterPercent))
	rewardVoter = new(big.Int).Div(rewardVoter, new(big.Int).SetInt64(100))
	for _, v := range voterStat {
		v.Reward = new(big.Int).Div(new(big.Int).Mul(rewardVoter, v.Cap), totalCap)
		if v.Reward.Cmp(new(big.Int)) > 0 {
			stateBlock.AddBalance(v.Address, v.Reward)
		}
	}
	rewards["voters"] = voterStat

	// Calc and rewards to foundation
	rewardFoundation := new(big.Int).Mul(currentEpochRewards, new(big.Int).SetInt64(common.RewardFoundationPercent))
	rewardFoundation = new(big.Int).Div(rewardFoundation, new(big.Int).SetInt64(100))
	rewards["foundation"] = struct {
		Address common.Address `json:"address"`
		Reward  *big.Int       `json:"reward"`
	}{c.Config().Foundation, rewardFoundation}
	if rewardFoundation.Cmp(new(big.Int)) > 0 {
		stateBlock.AddBalance(c.Config().Foundation, rewardFoundation)
	}

	epochStat["last"] = lastEpoch
	epochStat["current"] = nextEpoch
	epochStat["rewards"] = rewards
	bz, _ := json.Marshal(epochStat)
	_ = c.db.Put([]byte(fmt.Sprintf("%s-%d", common.EpochKeyPrefix, number/c.Config().Epoch)), bz)
	return epochStat, nil
}

func (c *PoSV) HookGetBlockSigners(chain consensus.ChainHeaderReader, stateBlock *state.StateDB, header *types.Header) (map[common.Address]int, error) {
	if header.Number.Uint64()%c.Config().Epoch != 0 {
		return nil, nil
	}
	lastCheckpoint := c.LastCheckpoint(header.Number.Uint64())
	lastEpoch, err := c.GetEpoch(chain, lastCheckpoint)
	if err != nil {
		return nil, errors.Errorf("GetEpoch %d err %v", lastCheckpoint, err)
	}

	signCount := make(map[common.Address]int)
	for i := lastCheckpoint + 1; i < header.Number.Uint64(); i++ {
		allowSigners := lastEpoch.M2(c.Config().Epoch, i)
		allowMap := make(map[common.Address]bool)
		for _, v := range allowSigners {
			allowMap[v] = true
		}

		hd := chain.GetHeaderByNumber(i)
		if hd != nil {
			for _, v := range contracts.GetBlockSignersFromState(stateBlock, hd.Hash()) {
				// Exclude maliciously signed M2
				if allowMap[v] {
					signCount[v] = signCount[v] + 1
				}
			}
		}
	}

	return signCount, nil
}

func (c *PoSV) HookBlockSign(signer common.Address, toSign, current *types.Header) error {
	if c.chainID == nil || c.txPool == nil || c.acm == nil {
		return nil
	}
	if err := contracts.CreateTransactionSign(c.chainID, c.txPool, c.acm, signer, toSign, current); err != nil {
		return errors.Errorf("Fail to create tx block sign for importing block: %v", err)
	}
	return nil
}

func (c *PoSV) HookSetRandomizeSecret(signer common.Address, header *types.Header) error {
	if c.chainID == nil || c.txPool == nil || c.acm == nil {
		return nil
	}
	if err := contracts.CreateTransactionSetSecret(c.chainID, c.txPool, c.acm, header, c.db, signer, c.config.Epoch); err != nil {
		return errors.Errorf("Fail to create tx set secret for importing block: %v", err)
	}
	return nil
}

func (c *PoSV) HookSetRandomizeOpening(signer common.Address, header *types.Header) error {
	if c.chainID == nil || c.txPool == nil || c.acm == nil {
		return nil
	}
	if err := contracts.CreateTransactionSetOpening(c.chainID, c.txPool, c.acm, header, c.db, signer); err != nil {
		return errors.Errorf("Fail to create tx set opening for importing block: %v", err)
	}
	return nil
}
