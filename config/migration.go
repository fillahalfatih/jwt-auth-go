package config

import (
	"log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, models ...interface{}) error {
	log.Println("Database migration completed successfully!")
	return db.AutoMigrate(models...)
}