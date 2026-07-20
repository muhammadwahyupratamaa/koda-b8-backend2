package svc

import (
	"errors"
	"koda-b8-backend1/internal/model"
	"koda-b8-backend1/internal/repo"
)


type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(r *repo.UserRepo) *UserService{
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Register(req *model.CreateUser) error{
	if req.Email == "" {
		return errors.New("Email is required")
	}
	if req.Password == "" {
		return errors.New("Password is required")
	}
	user := s.repo.FindByEmail(req.Email)

	if user != nil {
		return  errors.New("Email already exist")
	}
	s.repo.Create(req)
	return nil
}

func (s *UserService) Login(req *model.LoginUser) error{
	user := s.repo.FindByEmail(req.Email)

	if user == nil {
		return errors.New("User not found")
	}
	if user.Password != req.Password {
		return  errors.New("Invalid password")
	}
	return  nil
}

func (s *UserService) GetUser() []model.User{
	return s.repo.FindAll()
}
