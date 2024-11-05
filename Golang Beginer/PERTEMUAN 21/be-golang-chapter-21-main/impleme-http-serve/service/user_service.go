package service

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"be-golang-chapter-21/impleme-http-serve/repository"
)

type UserService struct {
	RepoUser repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{RepoUser: repo}
}

func (usr *UserService) LoginService(user *model.User) error {
	err := usr.RepoUser.Login(user)
	if err != nil {
		return err
	}

	return nil
}

func (usr *UserService) UserByID(id int) (*model.User, error) {

	User, err := usr.RepoUser.UserByID(id)
	if err != nil {
		return nil, err
	}
	return User, nil
}
