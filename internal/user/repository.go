// internal/user/repository.go
package user

import (
    "gorm.io/gorm"
)

// Definisikan interface agar mudah di-mock saat testing
type Repository interface {
    CreateUser(user *User) error
    FindByEmail(email string) (*User, error)
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

// Implementasi method FindByEmail
func (r *repository) FindByEmail(email string) (*User, error) {
    var user User
    err := r.db.First(&user, "email = ?", email).Error
    if err != nil {
        return nil, err // Akan mengembalikan error, misal gorm.ErrRecordNotFound jika tidak ada
    }
    return &user, nil
}