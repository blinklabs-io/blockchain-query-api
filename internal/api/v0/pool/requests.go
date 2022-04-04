package pool

// URI params for GetPool
type GetPoolUriParams struct {
	Address string `uri:"address" binding:"required"`
}
