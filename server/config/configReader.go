package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	Address string `json:"address"`
	//Port    string `json:"port"`
	Options struct {
		ReadTimeout               time.Duration `json:"readTimeout"`
		WriteTimeout              time.Duration `json:"writeTimeout"`
		RequestWaitInQueueTimeout time.Duration `json:"requestWaitInQueueTimeout"`
		Concurrency               int           `json:"concurrency"`
	}
}

func LoadConfig(path string) Config {
	var config Config
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalf("error in configuration file %s: %s", path, err)
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatalf("error in configuration file %s: %s", path, err)
	}
	return config
}
