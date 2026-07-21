package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return  func(c *gin.Context) {
		if c.Request.URL.Path =="/login" || c.Request.URL.Path == "/register" {
			c.Next()
			return
		}
		token := c.GetHeader("Authorization")

		if token != "hello" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}