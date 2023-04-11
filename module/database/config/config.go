package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbUsername         string `mapstructure:"db_username"`
	DbPassword         string `mapstructure:"db_password"`
	DbHost             string `mapstructure:"db_host"`
	DbPort             string `mapstructure:"db_port"`
	DbName             string `mapstructure:"db_name"`
	DbMain             string `mapstructure:"db_main"`
	DbPostgresUsername string `mapstructure:"dbPostgres_username"`
	DbPostgresPassword string `mapstructure:"dbPostgres_password"`
	DbPostgresHost     string `mapstructure:"dbPostgres_host"`
	DbPostgresPort     string `mapstructure:"dbPostgres_port"`
	DbPostgresName     string `mapstructure:"dbPostgres_name"`
	DbPostgresMain     string `mapstructure:"dbPostgres_main"`
	DbStart            string `mapstructure:"db_start"`
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
