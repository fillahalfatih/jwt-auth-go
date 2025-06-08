// internal/user/handler.go
package user

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Handler sekarang butuh service
type Handler struct {
    service Service
}

// Constructor untuk handler
func NewHandler(service Service) *Handler {
    return &Handler{service: service}
}

// Jadikan RegisterHandler sebagai method dari struct Handler
func (h *Handler) RegisterHandler(c *gin.Context) {
    var registerRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // 1. Bind request
    if err := c.ShouldBindJSON(&registerRequest); err != nil {
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
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}