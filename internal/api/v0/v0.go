package v0

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0/account"
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0/block"
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0/epoch"
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0/meta"
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0/tip"
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0/totals"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	account.ConfigureRoutes(g)
	block.ConfigureRoutes(g)
	epoch.ConfigureRoutes(g)
	meta.ConfigureRoutes(g)
	tip.ConfigureRoutes(g)
	totals.ConfigureRoutes(g)
}
