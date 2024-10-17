package router

import (
	"github.com/gin-gonic/gin"
	"random-commander/controller"
)

func DeckRoutes(route *gin.Engine, ctrl *controller.DeckController) {
	v1 := route.Group("/v1")
	v1.GET("/deck/random", ctrl.GetRandomDeck)
}
