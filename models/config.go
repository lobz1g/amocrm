package models

import (
	"encoding/json"
	"os"
)

const PATH_TO_CONFIG = "./config/amocrm.json"

type config struct {
	Domain string `json:"domain"`
	Login  string `json:"login"`
	Key    string `json:"key"`
}

func getConfig() (*config, error) {
	file, err := os.Open(PATH_TO_CONFIG)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var cfg *config
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
