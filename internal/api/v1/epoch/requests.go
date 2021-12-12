package epoch

// URI params for GetEpoch
type GetEpochUriParams struct {
	Number uint32 `uri:"number" binding:"required"`
}
