// internal/user/service.go
package user

import (
    "errors"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

// Interface untuk service
type Service interface {
    RegisterUser(email, password string) (*User, error)
    LoginUser(email, password string) (*User, error)
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

// Logika bisnis untuk login
func (s *service) LoginUser(email, password string) (*User, error) {
    
    // 1. Panggil repository untuk mencari user
    user, err := s.repository.FindByEmail(email)
    if err != nil {
        // Jika error-nya adalah record tidak ditemukan, beri pesan yang jelas
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        // Untuk error database lainnya
        return nil, err
    }

    // 2. Verifikasi password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        // Error ini biasanya berarti password salah
        return nil, errors.New("invalid password")
    }

    // 3. Jika semua cocok, kembalikan user
    return user, nil
}