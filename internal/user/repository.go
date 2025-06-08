// internal/user/repository.go
package user

import (
    "gorm.io/gorm"
)

// Definisikan interface agar mudah di-mock saat testing
type Repository interface {
    CreateUser(user *User) error
}

type repository struct {
    db *gorm.DB
}

// Constructor untuk membuat instance repository baru
func NewRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

// Implementasi method CreateUser
func (r *repository) CreateUser(user *User) error {
    return r.db.Create(user).Error
}