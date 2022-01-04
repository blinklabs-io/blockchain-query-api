package pool

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/pool", HandleGetPoolList)
	// groupPool := g.Group("/pool")
	// groupPool.GET("/:address", HandleGetPool)
}

// func HandleGetPool(c *gin.Context) {
// 	var uriParams GetPoolUriParams
// 	if err := c.ShouldBindUri(&uriParams); err != nil {
// 		// TODO: provide a more useful error message
// 		c.JSON(400, gin.H{"msg": err.Error()})
// 		return
// 	}
// 	// Retrieve pool from DB
// 	var pool models.PoolUpdate
// 	db := cardano_db_sync.GetHandle()
// 	// Retrieve latest registration ID per pool
// 	max_txn_ids := db.Model(&PoolUpdate{}).Select("max(registered_tx_id)").Group("hash_id")
// 	poolIdQuery := db.Select("id").Where("view = ?", uriParams.Pool).Table("pool_hash")
// 	result := db.Model(&models.PoolUpdate{}).Where("registered_tx_id in (?) AND id in (?)", max_txn_ids, poolIdQuery).First(&pool)
// 	if result.Error != nil {
// 		// Not found
// 		if cardano_db_sync.IsRecordNotFoundError(result.Error) {
// 			c.JSON(404, gin.H{"msg": "epoch not found"})
// 			return
// 		}
// 		// Some other database error
// 		// TODO: log this failure
// 		c.JSON(500, gin.H{"msg": "unexpected error"})
// 		return
// 	}
// 	// Create response from returned item
// 	r := NewPoolResponse(&pool)
// 	c.JSON(200, r)
// }

func HandleGetPoolList(c *gin.Context) {
	db := cardano_db_sync.GetHandle()
	var pools []*models.PoolUpdate

	// Retrieve latest registration ID per pool
	max_txn_ids := db.Model(&models.PoolUpdate{}).Select("max(registered_tx_id)").Group("hash_id")
	result := db.Model(&models.PoolUpdate{}).Where("registered_tx_id in (?)", max_txn_ids).Limit(100).Find(&pools)
	//.Select("view AS addresses").Joins("inner join pool_hash on pool_update.hash_id = pool_hash.id").Find(&pool_update)
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
	var r []*PoolResponse
	for _, v := range pools {
		r = append(r, NewPoolResponse(v))
	}
	c.JSON(200, r)
}
