package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App		 	string `mapStructure:"APP"`
	Port     	string `mapStructure:"PORT"`
	Database 	string `mapStructure:"DATABASE"`
	RabbitMQURL	string `mapStructure:"RABBITMQURL"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	// search defined path of file
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			return nil, errors.New(".env not found")
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