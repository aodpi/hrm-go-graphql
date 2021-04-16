package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error

	log.Print("This is the environment: ", env)
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)

	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config.AddConfigPath("config/")
	config.AddConfigPath(".")

	err = config.ReadInConfig()

	if err != nil {
		log.Fatalf("error on parsing configuration file: %s", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
