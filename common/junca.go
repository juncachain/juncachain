package common

func BuildInTx(to Address) bool {
	return to == HexToAddress(BlockSignerSMC) ||
		to == HexToAddress(ValidatorSMC) ||
		to == HexToAddress(RandomizeSMC) ||
		to == HexToAddress(JRC21IssuerSMC)
}
