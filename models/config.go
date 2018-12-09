package models

import (
	"../logHandler"
	"encoding/json"
	"os"
)

type config struct {
	Domain string `json:"domain"`
	Login  string `json:"login"`
	Key    string `json:"key"`
}

func getConfig() config {
	file, err := os.Open("./config/amocrm.json")
	defer file.Close()
	if err != nil {
		defer logHandler.WriteLogFile(err, "config", "getConfig()")
	}

	var cfg config
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&cfg)
	if err != nil {
		defer logHandler.WriteLogFile(err, "config", "jsonParser.Decode()")
	}
	return cfg
}
