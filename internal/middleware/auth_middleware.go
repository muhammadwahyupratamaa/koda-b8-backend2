package middleware

import (
	"net/http"
	"strings"

	"koda-b8-backend1/internal/lib"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		if path == "/login" ||
			path == "/register" ||
			strings.HasPrefix(path, "/uploads") {
			c.Next()
			return
		}

		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token required",
			})
			c.Abort()
			return
		}

		split := strings.Split(auth, " ")

		if len(split) != 2 || split[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token format",
			})
			c.Abort()
			return
		}

		_, err := lib.VerifyToken(split[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
