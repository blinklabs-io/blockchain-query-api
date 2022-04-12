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

type List struct {
	PoolIdBech32 string `"gorm:"column:pool_id_bech32"`
	Ticker       string `"gorm:"column:ticker"`
}

type Metadata struct {
	PoolIdBech32 string `"gorm:"column:pool_id_bech32"`
	MetaUrl      string `gorm:"column:meta_url"`
	MetaHash     string `gorm:"column:meta_hash"`
	MetaJson     jsonb  `gorm:"column:meta_json"`
}

type Pool struct {
	PoolIdBech32      string  `"gorm:"column:pool_id_bech32"`
	PoolIdHex         string  `"gorm:"column:pool_id_hex"`
	ActiveEpochNumber int64   `gorm:"column:active_epoch_no"`
	VrfHashKey        string  `gorm:"column:vrf_hash_key"`
	Margin            float32 `gorm:"column:margin"`
	FixedCost         uint64  `gorm:"column:fixed_cost"`
	Pledge            uint64  `gorm:"column:pledge"`
	RewardAddress     string  `gorm:"column:reward_address"`
	Owners            string  `gorm:"column:owners"`
	Relays            jsonb   `gorm:"column:relays"`
	MetaUrl           string  `gorm:"column:meta_url"`
	MetaHash          string  `gorm:"column:meta_hash"`
	MetaJson          jsonb   `gorm:"column:meta_json"`
	PoolStatus        string  `gorm:"column:pool_status"`
	RetiringEpoch     uint64  `gorm:"column:retiring_epoch"`
	OpCert            string  `gorm:"column:op_cert"`
	OpCertCounter     uint32  `gorm:"column:op_cert_counter"`
	ActiveStake       uint64  `gorm:"column:active_stake"`
	BlockCount        float32 `gorm:"column:block_count"`
	LivePledge        float32 `gorm:"column:live_pledge"`
	LiveStake         uint64  `gorm:"column:live_stake"`
	LiveDelegators    int64   `gorm:"column:live_delegators"`
	LiveSaturation    float32 `gorm:"column:live_saturation"`
}

type Relays struct {
	PoolIdBech32 string `"gorm:"column:pool_id_bech32"`
	Relays       jsonb  `gorm:"column:relays"`
}

type Updates struct {
	TxHash            string     `gorm:"column:tx_hash"`
	BlockTime         *time.Time `gorm:"column:block_time"`
	PoolIdBech32      string     `gorm:"column:pool_id_bech32"`
	PoolIdHex         string     `gorm:"column:pool_id_hex"`
	ActiveEpochNumber int64      `gorm:"column:active_epoch_no"`
	VrfHashKey        string     `gorm:"column:vrf_hash_key"`
	Margin            float32    `gorm:"column:margin"`
	FixedCost         uint64     `gorm:"column:fixed_cost"`
	Pledge            uint64     `gorm:"column:pledge"`
	RewardAddress     string     `gorm:"column:reward_address"`
	Owners            string     `gorm:"column:owners"`
	Relays            jsonb      `gorm:"column:relays"`
	MetaUrl           string     `gorm:"column:meta_url"`
	MetaHash          string     `gorm:"column:meta_hash"`
	PoolStatus        string     `gorm:"column:pool_status"`
	RetiringEpoch     uint64     `gorm:"column:retiring_epoch"`
}

func HandleGetPoolBlocks(c *gin.Context) {
	pool := c.DefaultPostForm("_pool_bech32", "")
	epoch := c.DefaultPostForm("_epoch_no", "NULL")

	db := cardano_db_sync.GetHandle()
	var blocks []*Block
	var poolId uint64
	r := []*BlockResponse{}
	if pool != "" {
		// Get database ID of our pool
		poolIdResult := db.Table("grest.pool_info_cache").
			Select("pool_hash_id").
			Where("pool_id_bech32 = ?", pool).
			Order("tx_id DESC").
			Limit(1).
			Find(&poolId)
		if poolIdResult.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(poolIdResult.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		result := db.Debug().
			Table("block b").
			Select("b.epoch_no, b.epoch_slot_no as epoch_slot, b.slot_no as abs_slot, b.block_no as block_height, b.hash as block_hash, b.time").
			Joins("INNER JOIN public.slot_leader AS sl ON b.slot_leader_id = sl.id").
			Where("sl.pool_hash_id = (?) AND (?)",
				poolIdResult,
				db.Raw("? IS NULL OR b.epoch_no = ?",
					epoch,
					epoch)).
			Find(&blocks)
		if result.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(result.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		for _, v := range blocks {
			r = append(r, NewBlockResponse(v))
		}
	}
	c.JSON(200, r)
}

// TODO: implement handlers
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
