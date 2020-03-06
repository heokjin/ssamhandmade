package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func InitDB() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Got error when connect database")
	}
	log.Println("DB:", db)
	// for task table
	db.LogMode(true)
}

func InitSchema() {
	db.AutoMigrate(
		&User{},
	)
}