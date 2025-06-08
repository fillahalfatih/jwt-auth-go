package config

import (
	"log"
	"jwt-auth-go/internal/model"
)

func Migrate() {
	err := DB.AutoMigrate(
		&model.User{},
	)
	
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migration completed successfully!")
}