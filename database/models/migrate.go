package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	// Add Models
	db.AutoMigrate(&Realm{})

	// Add Foreign Keys

	fmt.Printf("Auto-migration completed.\n")
}
