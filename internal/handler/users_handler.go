package handler

import (
	"fmt"
	"koda-b8-backend1/internal/lib"
	"koda-b8-backend1/internal/model"
	"koda-b8-backend1/internal/svc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *svc.UserService
}

func NewUserHandler(service *svc.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Register godoc
//
//	@Summary		Register a new User
//	@Description	Create a new user account
//	@Tags			Auth
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			name		formData	string	true	"User name"
//	@Param			email		formData	string	true	"User Email"	format(email)
//	@Param			password	formData	string	true	"User Password"	format(password)
//	@Success		201			{object}	lib.Response
//	@Failure		400			{object}	lib.Response
//	@Router			/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req model.CreateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	if err := h.service.Register(&req); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, lib.Response{
		Success: true,
		Message: "Register Success",
	})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login using email and password
//	@Tags			Auth
//	@Accept			application/x-www-form-urlencoded
//	@Produce		json
//	@Param			email		formData	string	true	"User Email"	format(email)
//	@Param			password	formData	string	true	"User Password"	format(password)
//	@Success		200			{object}	lib.Response
//	@Failure		400			{object}	lib.Response
//	@Failure		500			{object}	lib.Response
//	@Router			/login [post]
func (h *UserHandler) Login(c *gin.Context) {

	var form model.LoginUser

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := h.service.Login(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	token, err := lib.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Login Success",
		Result: gin.H{
			"token": token,
		},
	})
}

// GetUser godoc
//
//	@Summary		Get all User
//	@Description	Retrieve all User
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			page			query		int		false	"Page Number"		default(1)
//	@Param			limit			query		int		false	"Items per page"	default(5)
//	@Param			search[name]	query		string	false	"Search by name"
//	@Param			search[email]	query		string	false	"Search by email"
//	@Success		200				{object}	lib.Response
//	@Failure		500				{object}	lib.Response
//	@Router			/users [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 5
	}

	searchName := c.Query("search[name]")
	searchEmail := c.Query("search[email]")

	users, err := h.service.GetUser(page, limit, searchName, searchEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success",
		Result:  users,
	})
}

// GetUserByID godoc
//
//	@Summary		Get User by Id
//	@Description	Retrieve user Id
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	lib.Response
//	@Failure		400	{object}	lib.Response
//	@Failure		404	{object}	lib.Response
//	@Param			id	path		int	true	"User ID"
//	@Router			/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid user id",
		})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success",
		Result:  user,
	})
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update User
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id			path		int		true	"User ID"
//	@Param			name		formData	string	false	"User nama"
//	@Param			email		formData	string	false	"User email"	format(email)
//	@Param			password	formData	string	false	"User password"	format(password)
//	@Param			picture		formData	string	false	"profile picture"
//	@Success		200			{object}	lib.Response
//	@Failure		400			{object}	lib.Response
//	@Failure		500			{object}	lib.Response
//	@Router			/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var req model.UpdateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("picture")
	if err == nil {

		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

		err = c.SaveUploadedFile(file, "./uploads/"+fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		req.Picture = fileName
	}

	err = h.service.UpdateUser(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User updated successfully",
	})
}

// DeleteUser godoc
//
//	@Summary		Delete User
//	@Description	Delete user by ID
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"user id"
//	@Success		200	{object}	lib.Response
//	@Failure		400	{object}	lib.Response
//	@Router			/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User deleted successfully",
	})

}

// CreateUser godoc
//
//	@Summary		Create New User
//	@Description	Create New User with profile picture
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		BearerAuth
//	@Param			name		formData	string	true	"User Name"
//	@Param			email		formData	string	true	"User email"	format(email)
//	@Param			password	formData	string	true	"User password"	format(password)
//	@Param			picture		formData	file	false	"Profile Picture"
//	@Success		201			{object}	lib.Response
//	@Failure		400			{object}	lib.Response
//	@Failure		500			{object}	lib.Response
//	@Router			/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {

	var req model.CreateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("picture")
	if err == nil {
		const maxSize = 2 << 20
		if file.Size > maxSize {
			c.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: "maximum file size 2 mb",
			})
			return
		}

		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

		err = c.SaveUploadedFile(file, "./uploads/"+fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		req.Picture = fileName
	}

	if err := h.service.CreateUser(&req); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, lib.Response{
		Success: true,
		Message: "User created successfully",
	})
}

// UploadFile godoc
//
//	@Summary		Upload File
//	@Description	Upload file Image
//	@Tags			Upload
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		BearerAuth
//	@Param			file	formData	file	true	"Image File"
//	@Success		200		{object}	lib.Response
//	@Failure		400		{object}	lib.Response
//	@Failure		500		{object}	lib.Response
//	@Router			/upload [post]
func (h *UserHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println("File Size:", file.Size)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "file is required",
		})
		return
	}

	const maxSize = 2 << 20
	if file.Size > maxSize {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "maximum file size 2 mb",
		})
		return
	}

	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

	err = c.SaveUploadedFile(file, "./uploads/"+fileName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Upload Success",
		Result: gin.H{
			"file": fileName,
		},
	})
}
