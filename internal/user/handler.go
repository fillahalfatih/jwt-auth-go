package user

import (
	"jwt-auth-go/config"
	"jwt-auth-go/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	// get the request body
	// and bind it to the struct
	var registerRequest struct {
		Email    string
		Password string
	}

	err := c.ShouldBindJSON(&registerRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user in the database
	user := model.User{
		Email:    registerRequest.Email,
		Password: string(hashedPassword),
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}