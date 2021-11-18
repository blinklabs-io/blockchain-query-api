package api

import (
	"fmt"
	"github.com/cloudstruct/blockchain-query-api/internal/api/v1"
	"github.com/cloudstruct/blockchain-query-api/internal/config"
	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	// Disable gin debug output
	gin.SetMode(gin.ReleaseMode)

	// Configure router
	router := gin.Default()

	// Configure PING route
	router.GET("/ping", HandlePing)

	// Configure v1 routes
	routerGroupV1 := router.Group("/v1")
	v1.ConfigureRoutes(routerGroupV1)

	// Start listener
	router.Run(fmt.Sprintf("%s:%d", cfg.Api.Address, cfg.Api.Port))
}

func HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
