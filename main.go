/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/abstraction-hq/abstraction-wallet-node/cmd"
	"github.com/abstraction-hq/abstraction-wallet-node/config"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	config.InitEnv()
	cmd.Execute()
}
