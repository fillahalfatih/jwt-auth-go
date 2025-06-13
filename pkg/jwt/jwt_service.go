package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTService adalah interface untuk service JWT
type Service interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// jwtService adalah implementasi dari JWTService
type jwtService struct {
	secretKey string
}

// NewService membuat instance baru dari jwtService
func NewService() Service {
	// Ambil secret key dari environment variable
	return &jwtService{secretKey: os.Getenv("SECRET_KEY")}
}

func (s *jwtService) GenerateToken(userID uint) (string, error) {
	// Buat claims dengan subject (sub) berisi userID
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	// Buat token baru dengan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}