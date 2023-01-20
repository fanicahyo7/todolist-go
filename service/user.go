package service

import (
	"errors"
	"todolist/model"
	"todolist/repository"
	"todolist/util"
)

type UserService interface {
	// GetUserWithPassword(username string, password string) (model.User, error)
	RegisterUser(user model.User, password string) (model.User, string, error)
	LoginUser(usernameOrEmail string, password string) (model.UserResponse, string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user model.User, password string) (model.User, string, error) {
	pass, err := util.HashPassword(password)
	if err != nil {
		return model.User{}, "", err
	}
	user, err = s.repo.SaveUser(user, pass)
	if err != nil {
		return model.User{}, "", err
	}

	token, err := util.CreateJWT(user.ID)
	if err != nil {
		return model.User{}, "", err
	}

	return user, token, nil
}

func (s *userService) LoginUser(usernameOrEmail string, password string) (model.UserResponse, string, error) {
	userDB, err := s.repo.GetUser(usernameOrEmail)
	if err != nil {
		return model.UserResponse{}, "", err
	}

	passwordFromDB := userDB.Password

	isTrue := util.CheckPasswordHash(password, passwordFromDB)
	if isTrue {
		token, err := util.CreateJWT(userDB.ID)
		if err != nil {
			return model.UserResponse{}, "", err
		}

		userResponse := model.UserResponse{
			ID:       userDB.ID,
			Username: userDB.Username,
			Email:    userDB.Email,
			Created:  userDB.Created,
			Updated:  userDB.Updated,
		}

		return userResponse, token, nil
	} else {
		return model.UserResponse{}, "", errors.New("invalid username or password")
	}

}
