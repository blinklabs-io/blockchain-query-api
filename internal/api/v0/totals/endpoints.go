package totals

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/totals", HandleGetTotals)
	g.POST("/totals", HandleGetTotals)
}

func HandleGetTotals(c *gin.Context) {
	epochNumber := c.DefaultPostForm("_epoch_no", "")

	var adaPots []*models.AdaPots
	db := cardano_db_sync.GetHandle()
	if epochNumber != "" {
		epoch, err := strconv.ParseUint(epochNumber, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"msg": "could not parse _epoch_no"})
			return
		}
		result := db.Model(&adaPots).Where(&models.AdaPots{
			EpochNumber: uint32(epoch)}).Order("epoch_no desc").Find(&adaPots)
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
	} else {
		// TODO: make this paginate, don't limit specifically
		result := db.Order("epoch_no desc").Limit(100).Find(&adaPots)
		if result.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(result.Error) {
				c.JSON(404, gin.H{"msg": "records not found"})
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
	}
	// Create response from returned item
	var r []*TotalsResponse
	for _, v := range adaPots {
		r = append(r, NewTotalsResponse(v))
	}
	c.JSON(200, r)
}
