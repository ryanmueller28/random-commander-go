package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"random-commander/services"
)

type DeckController struct {
	DeckService *services.DeckService
}

func NewDeckController(deckService *services.DeckService) *DeckController {
	return &DeckController{
		DeckService: deckService,
	}
}

func (c *DeckController) GetRandomDeck(ctx *gin.Context) {
	deck := c.DeckService.NewDeck()
	ctx.JSON(http.StatusOK, deck)
}
