package handler

import (
	"koda-b8-backend1/internal/model"
	"koda-b8-backend1/internal/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)



type UserHandler struct{
	service *svc.UserService
}

func NewUserHandler(service *svc.UserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req model.CreateUser

	if err := c.ShouldBind(&req)
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message" : err.Error(),
		})
		return
	}
	if err := h.service.Register(&req) 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success" : true,
		"message" :"Register Success",

	})

}
func (h *UserHandler) Login(c *gin.Context) {

	var req model.LoginUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := h.service.Login(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login Success",
	})
}
func (h *UserHandler) GetUser(c *gin.Context) {

	users := h.service.GetUser()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": users,
	})
}



