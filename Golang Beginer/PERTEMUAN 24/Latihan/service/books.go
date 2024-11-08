package service

import (
	"latihan/model"
	"latihan/model/books"
)

func (s *Service) GetDhasboardDataService(data *model.GetDhasboardData) error {
	return s.Repo.GetDhasboardDataRepo(data)
}

func (s *Service) AddBookDataService(data books.Book) error {
	return s.Repo.AddBookDataRepo(data)
}

func (s *Service) GetBookDataService(data *[]books.Book) error {
	return s.Repo.GetBookDataRepo(data)
}

func (s *Service) EditBookDataService(data books.Book) error {
	return s.Repo.EditBookDataRepo(data)
}

func (s *Service) DeleteBookService(id string) error {
	return s.Repo.DeleteBookRepo(id)
}
