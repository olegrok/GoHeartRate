package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/olegrok/GoHeartRate/server/config"
	"log"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	cfg := config.Cfg.Database
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.Host, cfg.Login, cfg.DBname, cfg.Password))
	if err != nil {
		log.Fatalf("database error: %s", err)
	}
	DB = db
	return DB
}
