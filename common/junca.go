package common

func BuildInTx(to Address) bool {
	return to == HexToAddress(BlockSignerSMC) ||
		to == HexToAddress(RandomizeSMC) ||
		to == HexToAddress(JuncaJRC21Issuer)
}
