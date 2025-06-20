// main.go
package main

import (
	"jwt-auth-go/config"
	"jwt-auth-go/middleware"
	"jwt-auth-go/internal/user"
	"jwt-auth-go/pkg/jwt"
	"jwt-auth-go/routes"
	"log"
)

func main() {
	// 1. Panggil fungsi ConnectDB untuk mendapatkan koneksi database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// 2. Jalankan migrasi
	err = config.Migrate(db, &user.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// 3. Buat semua instance (Dependency Injection)
	jwtService := jwt.NewService()
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService, jwtService)

	authMiddleware := middleware.NewAuthMiddleware(userService, jwtService)

	allHandlers := &routes.Handlers{
		UserHandler:    userHandler,
		AuthMiddleware: authMiddleware,
	}

	// 4. Setup routes dengan memberikan handler
	router := routes.SetupRoutes(allHandlers)

	log.Println("Starting server on port :8080")
	log.Fatal(router.Run(":8080"))
}