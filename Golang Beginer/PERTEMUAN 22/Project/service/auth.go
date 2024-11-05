package service

import (
	"main/database"
	"main/model"
	"main/repository"
	"time"
)

func (s *Service) LoginService(user *model.User) error {
	err := s.Repo.LoginRepo(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RegisterService(user *model.User) error {
	err := s.Repo.RegisterRepo(user)
	if err != nil {
		return err
	}
	return nil
}

func CheckToken() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	db, err := database.InitDB2()
	if err != nil {
		return
	}

	repo := repository.NewRepository(db)
	services := NewService(repo)

	for {
		select {
		case <-ticker.C:
			if !services.Repo.CleanExpiredTokensRepo2(db) {
				return
			}
		}
	}
}

func (s *Service) GetRoleService(token string) (string, error) {
	return s.Repo.GetRoleRepo(token)
}
