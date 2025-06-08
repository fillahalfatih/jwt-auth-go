// internal/user/service.go
package user

import "golang.org/x/crypto/bcrypt"

// Interface untuk service
type Service interface {
    RegisterUser(email, password string) (*User, error)
}

type service struct {
    repository Repository // Bergantung pada interface repository
}

// Constructor untuk service
func NewService(repository Repository) Service {
    return &service{repository: repository}
}

// Logika bisnis untuk registrasi
func (s *service) RegisterUser(email, password string) (*User, error) {
    // 1. Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // 2. Buat objek user
    newUser := &User{
        Email:    email,
        Password: string(hashedPassword),
    }

    // 3. Panggil repository untuk menyimpan ke DB
    err = s.repository.CreateUser(newUser)
    if err != nil {
        return nil, err
    }

    return newUser, nil
}