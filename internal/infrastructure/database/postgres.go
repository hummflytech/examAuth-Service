package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBconnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load variables from .env file, relying on system envs")
	}

	dns := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect to the database")
	}

	return db
}
