package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/olegrok/GoHeartRate/server/config"
	"log"
)

// DB is a pointer to database object
var DB *gorm.DB

// Connect creates new connection with database
func Connect() {
	cfg := config.Config.Database
	var err error
	DB, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.Host, cfg.Login, cfg.DBname, cfg.Password))
	if err != nil {
		log.Fatalf("database error: %s", err)
	}
	//DB.LogMode(true)
}
