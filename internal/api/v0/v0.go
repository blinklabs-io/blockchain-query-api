package v0

import (
	"github.com/cloudstruct/blockchain-query-api/internal/api/v0/block"
	"github.com/cloudstruct/blockchain-query-api/internal/api/v0/epoch"
	"github.com/cloudstruct/blockchain-query-api/internal/api/v0/meta"
	"github.com/cloudstruct/blockchain-query-api/internal/api/v0/tip"
	"github.com/cloudstruct/blockchain-query-api/internal/api/v0/totals"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	block.ConfigureRoutes(g)
	epoch.ConfigureRoutes(g)
	meta.ConfigureRoutes(g)
	tip.ConfigureRoutes(g)
	totals.ConfigureRoutes(g)
}
