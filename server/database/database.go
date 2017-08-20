package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=admin dbname=heartrate sslmode=disable password=admin")
	defer db.Close()
	if err != nil {
		log.Fatalf("database error: %s", err)
	}
	return db
}
