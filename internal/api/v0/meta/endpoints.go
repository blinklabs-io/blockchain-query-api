package meta

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/meta", HandleGetMeta)
}

func HandleGetMeta(c *gin.Context) {
	// Retrieve meta from DB
	var meta models.Meta
	db := cardano_db_sync.GetHandle()
	// We just have the one row
	result := db.First(&meta)
	if result.Error != nil {
		// Not found
		if cardano_db_sync.IsRecordNotFoundError(result.Error) {
			c.JSON(404, gin.H{"msg": "meta not found"})
			return
		}
		// Some other database error
		// TODO: log this failure
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}
	// Create response from returned item
	r := NewMetaResponse(&meta)
	c.JSON(200, r)
}
