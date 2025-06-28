// internal/user/handler.go
package user

import (
	"jwt-auth-go/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler sekarang butuh service
type Handler struct {
	service   Service
	jwtService jwt.Service
}

// Constructor untuk handler
func NewHandler(service Service, jwtService jwt.Service) *Handler {
	return &Handler{service: service, jwtService: jwtService}
}

// Jadikan RegisterHandler sebagai method dari struct Handler
func (h *Handler) RegisterHandler(c *gin.Context) {
    var registerRequest RegisterRequest

    // 1. Bind request
    if c.ShouldBindJSON(&registerRequest) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // 2. Panggil service (logika bisnis dipindahkan ke service)
    _, err := h.service.RegisterUser(registerRequest.Email, registerRequest.Password)
    if err != nil {
        // Di sini bisa diperiksa jenis errornya, misal email sudah ada, dll.
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    // 3. Beri response
    c.JSON(http.StatusCreated, RegisterResponse{Message: "User registered successfully"})
}

func (h *Handler) LoginHandler(c *gin.Context) {

    var loginRequest LoginRequest

    // 1. Bind request
    if c.ShouldBindJSON(&loginRequest) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // 2. Panggil service untuk melakukan semua logika login
    user, err := h.service.LoginUser(loginRequest.Email, loginRequest.Password)
    if err != nil {
        // Service akan mengembalikan error jika user tidak ditemukan atau password salah
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // 3. Generate JWT token melalui JWT Service
    tokenString, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

    // 4. Beri response
    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	response := LoginResponse{
		Message: "Login successful",
		Token:   tokenString,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) ValidateHandler(c *gin.Context) {
    user, _ := c.Get("currentUser")

    c.JSON(http.StatusOK, gin.H{
        "message": "You are authenticated!",
        "user":    user,
    })
}
