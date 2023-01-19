package service

import (
	"errors"
	"fmt"
	"todolist/model"
	"todolist/repository"
	"todolist/util"
)

type UserService interface {
	// GetUserWithPassword(username string, password string) (model.User, error)
	RegisterUser(user model.User, password string) (model.User, string, error)
	LoginUser(usernameOrEmail string, password string) (model.User, string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// func (s *userService) GetUserWithPassword(usernameOrEmail, password string) (model.User, error) {
// 	user, err := s.repo.GetUserWithPassword(usernameOrEmail, password)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

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

func (s *userService) LoginUser(usernameOrEmail string, password string) (model.User, string, error) {
	userDB, err := s.repo.GetUser(usernameOrEmail)
	if err != nil {
		return model.User{}, "", err
	}

	fmt.Println("hasil : ", userDB.Password)

	isTrue := util.CheckPasswordHash(userDB.Password, password)
	if isTrue {
		user, err := s.repo.GetUserWithPassword(usernameOrEmail, password)
		if err != nil {
			return model.User{}, "", err
		}
		token, err := util.CreateJWT(user.ID)
		if err != nil {
			return model.User{}, "", err
		}
		return user, token, nil
	} else {
		return model.User{}, "", errors.New("Invalid username or password")
	}

}
