package service

import (
	"latihan/repository"

	"go.uber.org/zap"
)

var ServiceF *Service

type Service struct {
	Repo   *repository.Repository
	Logger *zap.Logger
}

func NewService(repo *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: logger,
	}
}

// func NewService(repo repository.RepositoryI) *Service {
// repo := repository.NewRepository()
// ServiceF = &Service{Repo: repo}
// return &Service{Repo: repo}
// }

// func (s *Service) GetUsersDataService(users *[]customers.Customer) {
// 	s.Repo.GetUsersRepo(users)
// }

// func (s *Service) GetUsersDetailService(users *customers.Customer) {
// 	s.Repo.GetUserDetailRepo(users)
// }

// func (s *Service) DeleteUserService(id int) error {
// 	return s.Repo.DeleteUserRepo(id)
// }

// func (s *Service) DeleteTodoService(id int) error {
// 	return s.Repo.DeleteTodoRepo(id)
// }
