package config

import (
	log "github.com/inconshreveable/log15"

	"github.com/spf13/viper"
)

var logger = log.New("module", "config")

type envConfigs struct {
	RpcEndpoint     string   `mapstructure:"RPC_ENDPOINT"`
}

var ENV *envConfigs

func InitEnv() {
	ENV = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	// Tell viper the path/location of your env file. If it is root just add "."

	// Tell viper the name of your file
	viper.SetConfigName("app")

	// Tell viper the type of your file
	viper.SetConfigType("env")
	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		logger.Error("Error un marshal", err)
	}
	return
}