package service

import (
	"latihan/model"
	"latihan/model/response"
	"latihan/repository"
)

type Service struct {
	Repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetPageDataService(searchDate string, low_to_high string, page int) (response.PaginationResponse, error) {
	return s.Repo.GetDataPageRepo(searchDate, low_to_high, page)
}

func (s *Service) PlaceDetailService(data *model.ResponsePlaceDetail, id int) error {
	return s.Repo.PlaceDetailRepo(data, id)
}

func (s *Service) LocationByIdService(data *model.Locations) error {
	return s.Repo.LocationByIdRepo(data)
}

func (s *Service) TourPLanByIdService(data *[]model.TourPLan, tour_id int) error {
	return s.Repo.TourPLanByIdRepo(data, tour_id)
}
