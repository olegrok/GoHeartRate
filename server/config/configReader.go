package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Database struct {
		DBname   string `json:"dbname"`
		Host     string `json:"host"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	Address string `json:"address"`
	Options struct {
		ReadTimeout               time.Duration `json:"readTimeout"`
		WriteTimeout              time.Duration `json:"writeTimeout"`
		RequestWaitInQueueTimeout time.Duration `json:"requestWaitInQueueTimeout"`
		Concurrency               int           `json:"concurrency"`
	}
}

var Cfg Config

func LoadConfig(path string) Config {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalf("error in configuration file %s: %s", path, err)
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&Cfg)
	if err != nil {
		log.Fatalf("error in configuration file %s: %s", path, err)
	}

	return Cfg
}
