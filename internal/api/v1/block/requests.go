package block

// URI params for GetBlock
type GetBlockUriParams struct {
	Number uint32 `uri:"number" binding:"required"`
}
