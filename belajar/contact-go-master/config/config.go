package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapStructure:"port"`
	Storage string `mapStructure:"storage"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	// search defined path of file
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			return nil, errors.New(".env tidak ditemukan")
		}
		return nil, fmt.Errorf("fatal error config file %s", err)
	}

	// unmarshal parameter file .env to struct
	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("fatal error decode : %s", err)
	}

	return &config, nil
}
