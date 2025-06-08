// middleware/auth.go
package middleware

import (
	"fmt"
	"jwt-auth-go/internal/user"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// NewAuthMiddleware membuat instance middleware baru yang bergantung pada service
func NewAuthMiddleware(userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil token dari Authorization header (lebih umum daripada cookie)
		tokenString, err := c.Cookie("Authorization") // <-- Baca dari cookie
        if err != nil {
            // Jika cookie tidak ada, kirim error Unauthorized
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
            return
        }

		// 2. Validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Gunakan secret key dari environment variable, sama seperti saat membuat token
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 3. Cek apakah token sudah kedaluwarsa
			exp := int64(claims["exp"].(float64))
			if time.Now().Unix() > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				return
			}

			// 4. Dapatkan ID user dari claim
			// Ingat, di LoginHandler kita menyimpannya dengan key "sub"
			userID := uint(claims["sub"].(float64))

			// 5. Cari user di database melalui service
			foundUser, err := userService.FindUserByID(userID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				return
			}

			// 6. Simpan user di context untuk digunakan di handler selanjutnya
			c.Set("currentUser", foundUser)

			// Lanjutkan ke handler berikutnya
			c.Next()

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}