package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"random-commander/controller"
)

func InitRoutes(ctrl *controller.DeckController) *gin.Engine {
	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	err := router.SetTrustedProxies([]string{})
	if err != nil {
		return nil
	}

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	RegisterRoutes(router, ctrl)
	return router
}

func RegisterRoutes(route *gin.Engine, ctrl *controller.DeckController) {
	route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
	})

	route.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	DeckRoutes(route, ctrl)
}
