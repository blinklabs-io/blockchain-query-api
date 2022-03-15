package epoch

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/epoch", HandleGetEpochList)
	groupEpoch := g.Group("/epoch")
	// Operates on epoch
	groupEpoch.GET("/current", HandleGetEpochLatest)
	groupEpoch.GET("/latest", HandleGetEpochLatest)
	// Specific epoch
	groupEpoch.GET("/:number", HandleGetEpoch)
	groupEpochNum := groupEpoch.Group("/:number")
	// Stake
	groupEpochNum.GET("/stake", HandleGetEpochStake)
	groupEpochStake := groupEpochNum.Group("/stake")
	groupEpochStake.GET("/account/:account", HandleGetEpochStakeByAccount)
	groupEpochStake.GET("/pool/:pool", HandleGetEpochStakeByPool)
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

func HandleGetEpochLatest(c *gin.Context) {
	// Retrieve epoch from DB
	var epoch models.Epoch
	db := cardano_db_sync.GetHandle()
	latestEpochQuery := db.Select("max(no)").Table("epoch")
	result := db.Where("no = (?)", latestEpochQuery).Find(&epoch)
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

type StakeAmount struct {
	Amount uint64
}

func HandleGetEpochStake(c *gin.Context) {
	var uriParams GetEpochUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		// TODO: provide a more useful error message
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Retrieve epoch from DB
	var epochStake models.EpochStake
	var amount StakeAmount
	db := cardano_db_sync.GetHandle()
	result := db.Model(&epochStake).Where(&models.EpochStake{
		EpochNumber: uriParams.Number}).Select("SUM(amount) AS amount").Find(&amount)
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
	r := NewEpochStakeResponse(&amount)
	c.JSON(200, r)
}

func HandleGetEpochStakeByAccount(c *gin.Context) {
	var uriParams GetEpochStakeUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		// TODO: provide a more useful error message
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Retrieve epoch from DB
	var epochStake models.EpochStake
	var amount StakeAmount
	db := cardano_db_sync.GetHandle()
	addrIdQuery := db.Select("id").Where("view = ?", uriParams.Account).Table("stake_address")
	result := db.Model(&epochStake).Where("epoch_no = ? AND addr_id = (?)", uriParams.Number, addrIdQuery).Select("SUM(amount) AS amount").Find(&amount)
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
	r := NewEpochStakeResponse(&amount)
	c.JSON(200, r)
}

func HandleGetEpochStakeByPool(c *gin.Context) {
	var uriParams GetEpochStakeUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		// TODO: provide a more useful error message
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Retrieve epoch from DB
	var epochStake models.EpochStake
	var amount StakeAmount
	db := cardano_db_sync.GetHandle()
	poolIdQuery := db.Select("id").Where("view = ?", uriParams.Pool).Table("pool_hash")
	result := db.Model(&epochStake).Where("epoch_no = ? AND pool_id = (?)", uriParams.Number, poolIdQuery).Select("SUM(amount) AS amount").Find(&amount)
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
	r := NewEpochStakeResponse(&amount)
	c.JSON(200, r)
}

func HandleGetEpochList(c *gin.Context) {
	// Retrieve epochs from DB
	db := cardano_db_sync.GetHandle()
	var epochs []*models.Epoch
	// TODO: make this paginate, don't limit specifically
	result := db.Model(&models.Epoch{}).Limit(100).Find(&epochs)
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
	var r []*EpochResponse
        for _, v := range epochs {
		r = append(r, NewEpochResponse(v))
	}
	c.JSON(200, r)
}
