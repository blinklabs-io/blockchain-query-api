package block

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/block", HandleGetBlockList)
	groupBlock := g.Group("/block")
	// Operates on block
	groupBlock.GET("/current", HandleGetBlockLatest)
	groupBlock.GET("/latest", HandleGetBlockLatest)
	// Specific block
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

func HandleGetBlockLatest(c *gin.Context) {
	// Retrieve block from DB
	var block models.Block
	db := cardano_db_sync.GetHandle()
	latestBlockQuery := db.Select("max(block_no)").Table("block")
	result := db.Where("block_no = (?)", latestBlockQuery).Find(&block)
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

func HandleGetBlockList(c *gin.Context) {
	// Retrieve blocks from DB
	db := cardano_db_sync.GetHandle()
	var blocks []*models.Block
	// TODO: make this paginate, don't limit specifically
	result := db.Model(&models.Block{}).Limit(100).Find(&blocks)
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
	r := []*BlockResponse{}
	for _, v := range blocks {
		r = append(r, NewBlockResponse(v))
	}
	c.JSON(200, r)
}
