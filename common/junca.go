package common

func BuildInTx(to Address) bool {
	return to == HexToAddress(BlockSignerSMC) ||
		to == HexToAddress(MasternodeVotingSMC) ||
		to == HexToAddress(RandomizeSMC) ||
		to == HexToAddress(JuncaJRC21Issuer)
}
