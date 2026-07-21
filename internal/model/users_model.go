package model

import "time"


type User struct {
	ID int64
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUser struct{
	Name string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginUser struct {
	Email string `form:"email"`
	Password string `form:"password"`
}

type UpdateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

