package pool

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	// Blocks
	g.GET("/pool_blocks", HandleGetPoolBlocks)
	g.POST("/pool_blocks", HandleGetPoolBlocks)
	// Delegators
	g.GET("/pool_delegators", HandleGetPoolDelegators)
	g.POST("/pool_delegators", HandleGetPoolDelegators)
	// History
	g.GET("/pool_history", HandleGetPoolHistory)
	g.POST("/pool_history", HandleGetPoolHistory)
	// Info
	g.GET("/pool_info", HandleGetPoolInfo)
	g.POST("/pool_info", HandleGetPoolInfo)
	// List
	g.GET("/pool_list", HandleGetPoolList)
	// Metadata
	g.GET("/pool_metadata", HandleGetPoolMetadata)
	// Relays
	g.GET("/pool_relays", HandleGetPoolRelays)
	// Updates
	g.GET("/pool_updates", HandleGetPoolUpdates)
	g.POST("/pool_updates", HandleGetPoolUpdates)
}

type Block struct {
	EpochNumber uint64     `gorm:"column:epoch_no"`
	EpochSlot   uint64     `gorm:"column:epoch_slot"`
	AbsSlot     uint64     `gorm:"column:abs_slot"`
	BlockHeight uint64     `gorm:"column:block_height"`
	BlockHash   []byte     `gorm:"column:block_hash"`
	BlockTime   *time.Time `gorm:"column:block_time"`
}

type Delegator struct {
	StakeAddress string  `gorm:"column:stake_address"`
	TotalBalance float64 `gorm:"column:total_balance"`
	EpochNumber  uint64  `gorm:"column:epoch_no"`
}

type History struct {
	EpochNumber        uint64  `gorm:"column:epoch_no"`
	ActiveStake        uint64  `gorm:"column:active_stake"`
	ActiveStakePercent float32 `gorm:"column:active_stake_pct"`
	SaturationPercent  float32 `gorm:"column:saturation_pct"`
	BlockCount         uint64  `gorm:"column:block_cnt"`
	DelegatorCount     uint64  `gorm:"column:delegator_cnt"`
	Margin             float32 `gorm:"column:margin"`
	FixedCost          uint64  `gorm:"column:fixed_cost"`
	PoolFees           float32 `gorm:"column:pool_fees"`
	DelegRewards       float64 `gorm:"column:deleg_rewards"`
	EpochRos           float64 `gorm:"column:epoch_ros"`
}



















func HandleGetPool(c *gin.Context) {
	var uriParams GetPoolUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		// TODO: provide a more useful error message
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Retrieve pool from DB
	var pool models.PoolUpdate
	db := cardano_db_sync.GetHandle()
	// Retrieve latest registration ID per pool
	max_txn_ids := db.Model(&models.PoolUpdate{}).Select("max(registered_tx_id)").Group("hash_id")
	poolIdQuery := db.Select("id").Where("view = ?", uriParams.Address).Table("pool_hash")
	result := db.Model(&models.PoolUpdate{}).Where("registered_tx_id in (?) AND id in (?)", max_txn_ids, poolIdQuery).First(&pool)
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
	r := NewPoolResponse(&pool, uriParams.Address)
	c.JSON(200, r)
}

func HandleGetPoolList(c *gin.Context) {
	db := cardano_db_sync.GetHandle()
	var pools []*string

	// Retrieve latest registration ID per pool
	max_txn_ids := db.Model(&models.PoolUpdate{}).Select("max(registered_tx_id)").Group("hash_id")
	result := db.Model(&models.PoolUpdate{}).Select("view").Joins("inner join pool_hash on pool_update.hash_id = pool_hash.id").Where("registered_tx_id in (?)", max_txn_ids).Limit(100).Find(&pools)
	if result.Error != nil {
		// Not found
		if cardano_db_sync.IsRecordNotFoundError(result.Error) {
			c.JSON(404, gin.H{"msg": "pool not found"})
			return
		}
		// Some other database error
		// TODO: log this failure
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}
	// Create response from returned item
	c.JSON(200, pools)
}
