package tip

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/tip", HandleGetTip)
}

type Tip struct {
	Hash            []byte     `gorm:"column:hash"`
	EpochNumber     uint16     `gorm:"column:epoch_no"`
	SlotNumber      uint32     `gorm:"column:slot_no"`
	EpochSlotNumber uint32     `gorm:"column:epoch_slot_no"`
	BlockNumber     uint32     `gorm:"column:block_no"`
	Time            *time.Time `gorm:"column:time"`
}

func HandleGetTip(c *gin.Context) {
	var tip Tip
	db := cardano_db_sync.GetHandle()
	result := db.Order("id desc").Limit(1).Table("block").Find(&tip)
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
	r := NewTipResponse(&tip)
	c.JSON(200, r)
}
