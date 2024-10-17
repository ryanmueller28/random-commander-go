package main

import (
	"github.com/spf13/viper"
	"log"
	"random-commander/config"
	"random-commander/controller"
	"random-commander/router"
	"random-commander/services"
	"time"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "America/Chicago")
	loc, _ := time.LoadLocation("America/Chicago")

	time.Local = loc
	if err := config.SetupConfiguration(); err != nil {
		log.Fatal(err)
	}

	serv := &services.DeckService{}
	ctrl := controller.NewDeckController(serv)

	rtr := router.InitRoutes(ctrl)
	err := rtr.Run(":8080")
	if err != nil {
		return
	}
}
