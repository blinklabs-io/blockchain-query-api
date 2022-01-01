package epoch

// URI params for GetEpoch
type GetEpochUriParams struct {
	Number uint32 `uri:"number" binding:"required"`
}

// URI params for GetEpochStake
type GetEpochStakeUriParams struct {
	Number  uint32 `uri:"number" binding:"required"`
	Account string `uri:"account"`
	Pool    string `uri:"pool"`
}
