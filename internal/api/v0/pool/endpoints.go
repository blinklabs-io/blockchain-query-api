package pool

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	// "github.com/cloudstruct/blockchain-query-api/internal/datasource/koios/models"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/postgres/types"
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
	ActiveStakePercent float64 `gorm:"column:active_stake_pct"`
	SaturationPercent  float64 `gorm:"column:saturation_pct"`
	BlockCount         uint64  `gorm:"column:block_cnt"`
	DelegatorCount     uint64  `gorm:"column:delegator_cnt"`
	Margin             float32 `gorm:"column:margin"`
	FixedCost          uint64  `gorm:"column:fixed_cost"`
	PoolFees           float32 `gorm:"column:pool_fees"`
	DelegRewards       float64 `gorm:"column:deleg_rewards"`
	EpochRos           float64 `gorm:"column:epoch_ros"`
}

type List struct {
	PoolIdBech32 string `gorm:"column:pool_id_bech32"`
	Ticker       string `gorm:"column:ticker"`
}

type Metadata struct {
	PoolIdBech32 string      `gorm:"column:pool_id_bech32"`
	MetaUrl      string      `gorm:"column:meta_url"`
	MetaHash     string      `gorm:"column:meta_hash"`
	MetaJson     types.Jsonb `gorm:"column:meta_json"`
}

type Pool struct {
	PoolIdBech32      string      `gorm:"column:pool_id_bech32"`
	PoolIdHex         string      `gorm:"column:pool_id_hex"`
	ActiveEpochNumber int64       `gorm:"column:active_epoch_no"`
	VrfHashKey        string      `gorm:"column:vrf_hash_key"`
	Margin            float32     `gorm:"column:margin"`
	FixedCost         uint64      `gorm:"column:fixed_cost"`
	Pledge            uint64      `gorm:"column:pledge"`
	RewardAddress     string      `gorm:"column:reward_address"`
	Owners            string      `gorm:"column:owners"`
	Relays            types.Jsonb `gorm:"column:relays"`
	MetaUrl           string      `gorm:"column:meta_url"`
	MetaHash          string      `gorm:"column:meta_hash"`
	MetaJson          types.Jsonb `gorm:"column:meta_json"`
	PoolStatus        string      `gorm:"column:pool_status"`
	RetiringEpoch     uint64      `gorm:"column:retiring_epoch"`
	OpCert            string      `gorm:"column:op_cert"`
	OpCertCounter     uint32      `gorm:"column:op_cert_counter"`
	ActiveStake       uint64      `gorm:"column:active_stake"`
	BlockCount        float64     `gorm:"column:block_count"`
	LivePledge        float32     `gorm:"column:live_pledge"`
	LiveStake         uint64      `gorm:"column:live_stake"`
	LiveDelegators    int64       `gorm:"column:live_delegators"`
	LiveSaturation    float32     `gorm:"column:live_saturation"`
}

type Relays struct {
	PoolIdBech32 string      `gorm:"column:pool_id_bech32"`
	Relays       types.Jsonb `gorm:"column:relays"`
}

type Updates struct {
	TxHash            string      `gorm:"column:tx_hash"`
	BlockTime         *time.Time  `gorm:"column:block_time"`
	PoolIdBech32      string      `gorm:"column:pool_id_bech32"`
	PoolIdHex         string      `gorm:"column:pool_id_hex"`
	ActiveEpochNumber int64       `gorm:"column:active_epoch_no"`
	VrfHashKey        string      `gorm:"column:vrf_hash_key"`
	Margin            float32     `gorm:"column:margin"`
	FixedCost         uint64      `gorm:"column:fixed_cost"`
	Pledge            uint64      `gorm:"column:pledge"`
	RewardAddress     string      `gorm:"column:reward_address"`
	Owners            string      `gorm:"column:owners"`
	Relays            types.Jsonb `gorm:"column:relays"`
	MetaUrl           string      `gorm:"column:meta_url"`
	MetaHash          string      `gorm:"column:meta_hash"`
	PoolStatus        string      `gorm:"column:pool_status"`
	RetiringEpoch     uint64      `gorm:"column:retiring_epoch"`
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
				// TODO: validate this works as expected
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

func HandleGetPoolDelegators(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HandleGetPoolHistory(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HandleGetPoolInfo(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HandleGetPoolList(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HandleGetPoolMetadata(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HandleGetPoolRelays(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HandleGetPoolUpdates(c *gin.Context) {
	c.JSON(200, gin.H{})
}
