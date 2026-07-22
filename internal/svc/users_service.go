package svc

import (
	"errors"
	"koda-b8-backend1/internal/model"
	"koda-b8-backend1/internal/repo"
	"net/mail"
	"strings"
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
	
	if req.Name == "" {
		return errors.New("Name is required!")
	}
	if req.Email == "" {
		return errors.New("Email is required!")
	}
	if _,err := mail.ParseAddress(req.Email)
	err != nil{
		return errors.New("Email not valid!")
	}
	if req.Password == "" {
		return errors.New("Password is required!")
	}
	if len(req.Password) < 8 {
		return errors.New("Password must be 8 character!")
	}
	user := s.repo.FindByEmail(req.Email)

	if user != nil {
		return  errors.New("Email already exist1")
	}
	err := s.repo.Create(req)
	if err != nil {
	return err
}

return nil
}

func (s *UserService) Login(req *model.LoginUser) error{
	user := s.repo.FindByEmail(req.Email)

	if strings.TrimSpace(req.Email) == "" {
    return errors.New("Email address is required")
	}

	if strings.TrimSpace(req.Password) == "" {
    return errors.New("Password is required")
	}

	if user == nil {
		return errors.New("User not found")
	}
	if user.Password != req.Password {
		return  errors.New("Invalid password")
	}
	return  nil
}

func (s *UserService) GetUser() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id int64) (*model.User, error) {

	user := s.repo.FindByID(id)

	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (s *UserService) UpdateUser(id int64, req *model.UpdateUser) error{
	user := s.repo.FindByID(id)

	if user == nil {
		return errors.New("User not found!")
	}
	if req.Name == "" {
		return errors.New("Name is required!")
	} 
	if req.Email == "" {
		return errors.New("Email is required!")
	}
	if _,err := mail.ParseAddress(req.Email)
	err != nil{
		return errors.New("Email not valid!")
	}
	if req.Password == "" {
		return errors.New("Password is required!")
	}
	if len(req.Password) < 8 {
		return errors.New("Password must be 8 character")
	}
	return s.repo.Update(id,req)
}

func (s *UserService) DeleteUser(id int64) error{
	user := s.repo.FindByID(id)
	if user == nil {
		return errors.New("User Not Found")
	}
	return  s.repo.Delete(id)
}
func (s *UserService) CreateUser(req *model.CreateUser) error {

	if req.Name == "" {
		return errors.New("Name is required!")
	}

	if req.Email == "" {
		return errors.New("Email is required!")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("Email not valid!")
	}

	if req.Password == "" {
		return errors.New("Password is required!")
	}

	if len(req.Password) < 8 {
		return errors.New("Password must be 8 character!")
	}

	user := s.repo.FindByEmail(req.Email)

	if user != nil {
		return errors.New("Email already exist!")
	}

	return s.repo.Create(req)
}