package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "❌ Authorization header is required"})
			c.Abort()
			return
		}
	}
}
