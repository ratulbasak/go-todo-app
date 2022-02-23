package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	dbConn string
	environment string
}

func InitConfig(key string) string {
	// viper.SetConfigName(".env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		panic(fmt.Errorf("Invalid type assertion \n"))
	}
	return value
}