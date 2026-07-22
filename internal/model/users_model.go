package model

import "time"


type User struct {
	ID int64 `form:"ID" json:"id"`
	Name string `form:"name" json:name"` 
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"-"`
	CreatedAt time.Time `form:"createdAt" json:"createdAt"` 
	UpdatedAt time.Time `form:"updatedAt" json:"updatedAt"`
}

type CreateUser struct{
	Name string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"-"`
}

type LoginUser struct {
	Email string `form:"email"`
	Password string `form:"password"`
}

type UpdateUser struct {
	Name string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"-"`
}

