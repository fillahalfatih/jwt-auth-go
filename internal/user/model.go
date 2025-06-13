package user

import (
	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	gorm.Model
}