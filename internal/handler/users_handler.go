package handler

import (
	"koda-b8-backend1/internal/lib"
	"koda-b8-backend1/internal/model"
	"koda-b8-backend1/internal/svc"
	"net/http"
	"strconv"

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

	var form model.LoginUser

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := h.service.Login(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	token, err := lib.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login Success",
		"token": token,
	})
}



func (h *UserHandler) GetUser(c *gin.Context) {

	users, err := h.service.GetUser()
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
		return
	}

c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	var req model.UpdateUser

	if err := c.ShouldBind(&req)
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.UpdateUser(id,&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		}) 
	}
	c.JSON(http.StatusOK, gin.H{
		"message" :"user updated successfully",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "User deleted successfully",
	})

}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var req model.CreateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := h.service.CreateUser(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
	})
}

