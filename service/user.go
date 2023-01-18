package service

import (
	"todolist/model"
	"todolist/repository"
	"todolist/util"
)

type UserService interface {
	GetUser(username string, password string) (model.User, error)
	RegisterUser(user model.User, password string) (model.User, string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(usernameOrEmail, password string) (model.User, error) {
	user, err := s.repo.GetUser(usernameOrEmail, password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) RegisterUser(user model.User, password string) (model.User, string, error) {
	pass, err := util.HashPassword(password)
	if err != nil {
		return model.User{}, "", err
	}
	user, err = s.repo.RegisterUser(user, pass)
	if err != nil {
		return model.User{}, "", err
	}

	token, err := util.CreateJWT(user.ID)
	if err != nil {
		return model.User{}, "", err
	}

	return user, token, nil
}
