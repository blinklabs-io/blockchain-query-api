// TODO: migrate this to subdir like 'block'
package v1

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesEpoch(g *gin.RouterGroup) {
	groupEpoch := g.Group("/epoch")
	groupEpoch.GET("/:number", HandleGetEpoch)
}

// URI params for GetEpoch
type GetEpochUriParams struct {
	Number uint32 `uri:"number" binding:"required"`
}

// XXX: this doesn't actually do anything useful since the underlying 'epoch'
// table in the DB is empty
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
	db.Where(&models.Epoch{EpochNumber: uriParams.Number}).First(&epoch)
	// TODO: create struct for API response and populate/return that
	c.JSON(200, epoch)
}
