package v1

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesBlock(g *gin.RouterGroup) {
	groupBlock := g.Group("/block")
	groupBlock.GET("/:number", HandleGetBlock)
}

// URI params for GetBlock
type GetBlockUriParams struct {
	Number uint32 `uri:"number" binding:"required"`
}

func HandleGetBlock(c *gin.Context) {
	var uriParams GetBlockUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		// TODO: provide a more useful error message
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Retrieve block from DB
	var block models.Block
	db := cardano_db_sync.GetHandle()
	db.Where(&models.Block{BlockNumber: uriParams.Number}).First(&block)
	// TODO: create struct for API response and populate/return that
	c.JSON(200, block)
}
