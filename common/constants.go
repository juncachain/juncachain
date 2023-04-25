package common

const (
	BlockSignerSMC   = "0x000000000000000000426C6F636b5369676E6572" // []byte("BlockSigner")
	RandomizeSMC     = "0x000000000000000000000052616E646f6D697a65" // []byte("Randomize")
	ValidatorSMC     = "0x000000000000000000000056616c696461746f72" // []byte("Validator")
	JRC21IssuerSMC   = "0x0000000000000000004A52433231497373756572" // []byte("JRC21Issuer")
	JuncaswapWJGC    = "0x000000000000004a756e636153776170574A4743" // []byte("JuncaSwapWJGC")
	JuncaswapFactory = "0x000000004a756e636173776170466163746F7279" // []byte("JuncaswapFactory")
	JuncaswapRouter1 = "0x000000004A756E636173776170526f7574657231" // []byte("JuncaswapRouter1")
	JuncaswapRouter2 = "0x000000004a756E636173776170526F7574657232" // []byte("JuncaswapRouter2")
)

// contract Method
const (
	VoteMethod    = "0x6dd7d8ea"
	UnvoteMethod  = "0x02aa9be2"
	ProposeMethod = "0x01267951"
	ResignMethod  = "0xae6e43f5"
	SignMethod    = "0xe341eaa4"

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
	EpochBlockSignWiggle    = 8       // Delay sign the block
	EpochBlockCheckWiggle   = 8       // Based on confirmed block data
	EpochBlockSealMissAllow = 2       // Allow miss seal the blocks each epoch
	BuildInTxGas            = 1 << 18 // Fixed gas 262144
)

const (
	EpochKeyPrefix = "JUNCA-EPOCH-"
)
