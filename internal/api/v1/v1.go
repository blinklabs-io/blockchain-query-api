package v1

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	ConfigureRoutesEpoch(g)
	ConfigureRoutesBlock(g)
}
