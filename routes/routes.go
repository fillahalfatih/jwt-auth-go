package routes

import (
	"jwt-auth-go/internal/user"
	"github.com/gin-gonic/gin"
)


func SetupRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	
	v1.POST("/register", user.RegisterHandler)

	return r
}