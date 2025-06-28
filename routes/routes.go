// routes/routes.go
package routes

import (
	"jwt-auth-go/internal/product"
	"jwt-auth-go/internal/user"

	"github.com/gin-gonic/gin"
)

// Buat struct ini untuk menampung semua handler
type Handlers struct {
    AuthMiddleware gin.HandlerFunc
    UserHandler    *user.Handler
    ProductHandler *product.ProductHandler // Ganti dengan handler yang sesuai jika ada
    // Jika nanti ada ProductHandler, tambahkan di sini
    // ProductHandler *product.Handler
}

// Ubah signature fungsi ini untuk menerima struct Handlers
func SetupRoutes(handlers *Handlers) *gin.Engine {
    r := gin.Default()

    v1 := r.Group("/v1")

    // --- User Routes ---
    userRoutes := v1.Group("/users")
    {
        userRoutes.POST("/register", handlers.UserHandler.RegisterHandler)
        userRoutes.POST("/login", handlers.UserHandler.LoginHandler)
        userRoutes.GET("/validate", handlers.AuthMiddleware, handlers.UserHandler.ValidateHandler)
    }

    // --- Product Routes ---
    productRoutes := v1.Group("/products")
    {
        productRoutes.GET("/", handlers.ProductHandler.GetProducts)
        productRoutes.GET("/:id", handlers.ProductHandler.GetProductByID)
        productRoutes.POST("/create", handlers.AuthMiddleware, handlers.ProductHandler.PostProduct)
    }

    return r
}