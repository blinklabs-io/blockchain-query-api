package block

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	groupBlock := g.Group("/block")
	groupBlock.GET("/:number", HandleGetBlock)
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
	result := db.Where(&models.Block{BlockNumber: uriParams.Number}).First(&block)
	if result.Error != nil {
		// Not found
		if cardano_db_sync.IsRecordNotFoundError(result.Error) {
			c.JSON(404, gin.H{"msg": "block not found"})
			return
		}
		// Some other database error
		// TODO: log this failure
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}
	// Create response from returned item
	r := NewBlockResponse(&block)
	c.JSON(200, r)
}
