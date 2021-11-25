package v1

import (
	"github.com/cloudstruct/blockchain-query-api/internal/api/v1/block"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	ConfigureRoutesEpoch(g)
	block.ConfigureRoutes(g)
}
