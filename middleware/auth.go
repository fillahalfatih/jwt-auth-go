// middleware/auth.go
package middleware

import (
	"jwt-auth-go/internal/user"
	"net/http"
	"jwt-auth-go/pkg/jwt"

	"github.com/gin-gonic/gin"
	gojwt "github.com/golang-jwt/jwt/v4"
)

// NewAuthMiddleware membuat instance middleware baru yang bergantung pada service
func NewAuthMiddleware(userService user.Service, jwtService jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil token dari Authorization header (lebih umum daripada cookie)
		tokenString, err := c.Cookie("Authorization") // <-- Baca dari cookie
        if err != nil {
            // Jika cookie tidak ada, kirim error Unauthorized
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
            return
        }

		// 2. Validasi token
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
			// 3. Dapatkan ID user dari claim
			userID := uint(claims["sub"].(float64))

			// 4. Cari user di database (logika ini tetap di sini karena ini tugas middleware)
			foundUser, err := userService.FindUserByID(userID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				return
			}

			c.Set("currentUser", foundUser)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}