package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Storage string `mapstructure:"storage"`
	Port    string `mapstructure:"port"`
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("fatal error config file: %s", err))
	}
	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("fatal error decode file: %s", err))
	}

	return &config
}

// func SavePort() *Config {

// }
