package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realbucksavage/miniauth/database/models"
)

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=miniauth dbname=miniauth password=miniauth sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	models.AutoMigrate(db)
	fmt.Printf("Database connection ready.\n")
	return db, err
}
