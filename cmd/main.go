// Package main Backend CRUD API
//
//	@title						Backend CRUD API
//	@version					1.0
//	@description				REST API for User Management
//	@host						localhost:8080
//	@BasePath					/
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

package main

import (
	"log"

	"koda-b8-backend1/docs"
	"koda-b8-backend1/internal/di"
	"koda-b8-backend1/internal/lib"
	"koda-b8-backend1/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db, err := lib.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	docs.SwaggerInfo.Title = "Backend CRUD"
	docs.SwaggerInfo.Description = "This is a sample server backend CRUD."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	container := di.NewContainer(db)

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.AuthMiddleware())

	r.POST("/register", container.UserHandler().Register)
	r.POST("/login", container.UserHandler().Login)

	r.Static("/uploads", "./uploads")
	r.POST("/users", container.UserHandler().CreateUser)
	r.GET("/users", container.UserHandler().GetUser)
	r.POST("/upload", container.UserHandler().UploadFile)
	r.GET("/users/:id", container.UserHandler().GetUserByID)
	r.PUT("/users/:id", container.UserHandler().UpdateUser)
	r.DELETE("/users/:id", container.UserHandler().DeleteUser)
	r.Run(":8080")
}
