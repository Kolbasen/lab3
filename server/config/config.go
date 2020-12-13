package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	DbName     string `json:"db_name"`
	User       string `json:"user"`
	Host       string `json:"host"`
	DisableSSL bool   `json:"disable_ssl"`
}

type Application struct {
	TaxPercent      float64 `json:"no_tax_percent"`
	RecommendedTips float64 `json:"recommended_tips"`
}

type Config struct {
	Db  Database    `json:"db"`
	App Application `json:"application"`
}

// LoadConfiguration from json file
func LoadConfiguration(file string) (*Config, error) {
	var config *Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
