package service

import (
	"latihan/model"
	"latihan/repository"
)

type Service struct {
	Repo *repository.Travel
}

func NewService(repo *repository.Travel) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetPageDataService(data *[]model.ResponseDataPage, search string, low_to_high string, page int) error {
	return s.Repo.GetDataPageRepo(data, search, low_to_high, page)
}
