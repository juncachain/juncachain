package common

const (
	BlockSignerSMC                    = "0x000000000000000000426C6F636b5369676E6572" // []byte("BlockSigner")
	MasternodeVotingSMC               = "0x000000004D61737465726E6F6465566f74696e67"
	RandomizeSMC                      = "0x000000000000000000000052616E646f6D697a65"
	JuncaFoudation                    = "0x0000000000004a756E6361466F75646174696f6E"
	TeamAddr                          = "0x0000000000000000000000000000000000000099"
	TomoXAddr                         = "0x0000000000000000000000000000000000000091"
	TradingStateAddr                  = "0x0000000000000000000000000000000000000092"
	TomoXLendingAddress               = "0x0000000000000000000000000000000000000093"
	TomoXLendingFinalizedTradeAddress = "0x0000000000000000000000000000000000000094"
	TomoNativeAddress                 = "0x0000000000000000000000000000000000000001"
	LendingLockAddress                = "0x0000000000000000000000000000000000000011"
)

// contract Method
const (
	VoteMethod       = "0x6dd7d8ea"
	UnvoteMethod     = "0x02aa9be2"
	ProposeMethod    = "0x01267951"
	ResignMethod     = "0xae6e43f5"
	SignMethod       = "0xe341eaa4"
	TomoXApplyMethod = "0xc6b32f34"
	TomoZApplyMethod = "0xc6b32f34"

	HexSignMethod = "e341eaa4"
	HexSetSecret  = "34d38600"
	HexSetOpening = "e11f5ba2"
)

const (
	RewardM1Percent         = int64(40)
	RewardM2Percent         = int64(10)
	RewardVoterPercent      = int64(40)
	RewardFoundationPercent = int64(10)
	MaxMasternodes          = 150
	MaxSigners              = 99
	EpochBlockSecret        = 30
	EpochBlockOpening       = 20
	EpochBlockSignWiggle    = 15
	EpochBlockSealMissAllow = 2
)

const (
	EpochKeyPrefix = "JUNCA-EPOCH-"
)
