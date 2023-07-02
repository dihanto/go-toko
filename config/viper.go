package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitLoadConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	workDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(workDirectory)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
