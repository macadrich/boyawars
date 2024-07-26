package config

import (
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	NetworkURL      string
	ChainID         uint
	ContractAddress string
}

func Config() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config.yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return v
}

func LoadDBConfig() *DBConfig {
	v := Config()
	return &DBConfig{
		NetworkURL:      v.GetString("Blockchain.NetworkURL"),
		ChainID:         v.GetUint("Blockchain.ChainID"),
		ContractAddress: v.GetString("Blockchain.ContractAddress"),
	}
}
