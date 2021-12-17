package pool

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	g.GET("/pool", HandleGetPools)
	// groupPool := g.Group("/pool")
	// groupPool.GET("/:number", HandleGetPool)
}

func HandleGetPools(c *gin.Context) {
	db := cardano_db_sync.GetHandle()
	var pool_update models.PoolUpdate

	// Retrieve latest registration ID per pool
	max_txn_ids := db.Model(&PoolUpdate{}).Select("max(registered_tx_id)").Group("hash_id")
	result := db.Where("registered_tx_id in (?)", max_txn_ids).Select("view").Joins("inner join pool_hash on pool_update.hash_id = pool_hash.id").Find(&pool_update)
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
	r := NewPoolsResponse(&pool)
	c.JSON(200, r)
}
