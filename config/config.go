package config

import (
	"log"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("simplepki")
	viper.AddConfigPath("/etc/simplepki.yaml")
	viper.AddConfigPath("$HOME/.simplepki")
	viper.AddConfigPath(".")
	
	viper.SetEnvPrefix("simplepki")
	viper.BindEnv("endpoint")
	viper.BindEnv("account")
	viper.BindEnv("chain")
	viper.BindEnv("id")	
	viper.BindEnv("token")
}

func Load() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// this is ok; jut use env
		} else {
			log.Printf("Error reading config: %s\n", err.Error())
		}
	}
}