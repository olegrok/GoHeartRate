package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Config is structure that stores information about server and database settings
var Config struct {
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

//LoadConfig loads config from json file "path"
func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error in configuration file %s: %s", path, err)
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&Config)
	if err != nil {
		log.Fatalf("error in configuration file %s: %s", path, err)
	}
}
