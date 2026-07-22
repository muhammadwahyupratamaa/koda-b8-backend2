package main

import (
	"log"

	"koda-b8-backend1/internal/di"
	"koda-b8-backend1/internal/lib"
	"koda-b8-backend1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	container := di.NewContainer(db)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.AuthMiddleware())


	r.POST("/register", container.UserHandler().Register)
	r.POST("/login", container.UserHandler().Login)
	r.POST("/users", container.UserHandler().CreateUser)
	r.GET("/users", container.UserHandler().GetUser)
	r.POST("/upload", container.UserHandler().UploadFile)
	r.GET("/users/:id", container.UserHandler().GetUserByID)
	r.PUT("/users/:id", container.UserHandler().UpdateUser)
	r.DELETE("/users/:id", container.UserHandler().DeleteUser)
	r.Run(":8080")
}