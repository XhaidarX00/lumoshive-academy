package service

import (
	"main/model"
	"main/repository"
)

var ServiceF *Service

type Service struct {
	Repo repository.Repository
}

func NewService() {
	repo := repository.NewRepository()
	ServiceF = &Service{Repo: repo}
}

func (s *Service) GetUsersDataService(users *[]model.Users) {
	s.Repo.GetUsersRepo(users)
}

func (s *Service) GetUsersDetailService(users *model.Users) {
	s.Repo.GetUserDetailRepo(users)
}

func (s *Service) DeleteUserService(id int) error {
	return s.Repo.DeleteUserRepo(id)
}

func (s *Service) DeleteTodoService(id int) error {
	return s.Repo.DeleteTodoRepo(id)
}
