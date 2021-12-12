package v1

import (
	"github.com/cloudstruct/blockchain-query-api/internal/api/v1/block"
	"github.com/cloudstruct/blockchain-query-api/internal/api/v1/epoch"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	block.ConfigureRoutes(g)
	epoch.ConfigureRoutes(g)
}
