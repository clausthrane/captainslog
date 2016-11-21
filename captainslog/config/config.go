package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("mydata")
	viper.AddConfigPath("$HOME/.captainslog")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("No configuration file loaded - exiting")
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}
