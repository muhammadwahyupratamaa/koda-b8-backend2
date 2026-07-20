package main

import (
	"koda-b8-backend1/internal/di"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	container := di.NewContainer()

	r.POST("/register", container.UserHandler().Register)
	r.POST("/login", container.UserHandler().Login)
	r.GET("/users", container.UserHandler().GetUser)

	r.Run(":8080")
}