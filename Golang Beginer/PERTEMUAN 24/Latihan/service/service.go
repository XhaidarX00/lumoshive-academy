package service

import (
	"latihan/repository"
)

var ServiceF *Service

type Service struct {
	Repo repository.Repository
}

func NewService() {
	repo := repository.NewRepository()
	ServiceF = &Service{Repo: repo}
}

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
