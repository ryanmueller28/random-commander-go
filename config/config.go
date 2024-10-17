package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port   string
	Secret string
}

func SetupConfiguration() error {
	var configuration *Config
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return err
	}
	return nil
}

func ServerConfig() string {
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8000")
	appServer := fmt.Sprintf("%v:%v", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
	return appServer
}
