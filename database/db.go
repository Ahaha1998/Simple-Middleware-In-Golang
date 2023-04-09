package database

import (
	"challenge-12/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "w4hYu98"
	dbname   = "challenge"
	db       *gorm.DB
	err		 error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Success connecting to database.")
	db.Debug().AutoMigrate(models.UserModel{}, models.ProductModel{})
}

func GetDB() *gorm.DB {
	return db
}