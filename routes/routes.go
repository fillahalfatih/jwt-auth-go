// routes/routes.go
package routes

import (
    "jwt-auth-go/internal/user" // Impor user untuk mendapatkan struct Handler
    "github.com/gin-gonic/gin"
)

// Ubah signature fungsi ini untuk menerima user.Handler
func SetupRoutes(userHandler *user.Handler) *gin.Engine {
    r := gin.Default()

    v1 := r.Group("/v1")

    v1.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })
    
    // Panggil handler sebagai method dari instance yang kita terima
    v1.POST("/register", userHandler.RegisterHandler)

    return r
}