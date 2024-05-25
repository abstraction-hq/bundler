package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PrivateKey        string
	RpcEndpoint       string
	DatabaseDirectory string
	Port              int
}

func NewConfig(network string) *Config {
	viper.SetDefault("port", 4337)
	viper.SetDefault("data_directory", "/tmp/bundler")
	viper.SetDefault("supported_entry_points", "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")

	// Read in from .env file if available
	viper.SetConfigName(network)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			// Can ignore
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	return &Config{
		PrivateKey:        viper.GetString("private_key"),
		RpcEndpoint:       viper.GetString("rpc_endpoint"),
		DatabaseDirectory: viper.GetString("data_directory"),
		Port:              viper.GetInt("port"),
	}
}
