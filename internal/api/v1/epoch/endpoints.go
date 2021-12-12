package epoch

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	groupEpoch := g.Group("/epoch")
	groupEpoch.GET("/:number", HandleGetEpoch)
}

func HandleGetEpoch(c *gin.Context) {
	var uriParams GetEpochUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		// TODO: provide a more useful error message
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Retrieve epoch from DB
	var epoch models.Epoch
	db := cardano_db_sync.GetHandle()
	result := db.Where(&models.Epoch{EpochNumber: uriParams.Number}).First(&epoch)
	if result.Error != nil {
		// Not found
		if cardano_db_sync.IsRecordNotFoundError(result.Error) {
			c.JSON(404, gin.H{"msg": "epoch not found"})
			return
		}
		// Some other database error
		// TODO: log this failure
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}
	// Create response from returned item
	r := NewEpochResponse(&epoch)
	c.JSON(200, r)
}
