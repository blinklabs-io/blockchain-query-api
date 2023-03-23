package api

import (
	"fmt"
	"github.com/blinklabs-io/blockchain-query-api/internal/api/v0"
	"github.com/blinklabs-io/blockchain-query-api/internal/config"
	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) error {
	// Disable gin debug output
	gin.SetMode(gin.ReleaseMode)

	// Configure router
	router := gin.Default()

	// Configure PING route
	router.GET("/ping", HandlePing)

	// Configure api/v0 routes
	apiGroup := router.Group("api")
	routerGroupV0 := apiGroup.Group("/v0")
	v0.ConfigureRoutes(routerGroupV0)

	// Start listener
	err := router.Run(fmt.Sprintf("%s:%d", cfg.Api.Address, cfg.Api.Port))
	return err
}

func HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
