package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server
	Database
}

type Server struct {
	GRpcHost string
	GRpcPort uint16
	ResHost  string
	RestPort uint16
}

type Database struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     uint16
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read configuration file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to map configuration file to struct: %w", err)
	}

	return config, nil
}
