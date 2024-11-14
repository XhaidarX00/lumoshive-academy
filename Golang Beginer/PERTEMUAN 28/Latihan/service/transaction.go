package service

import "latihan/model"

func (s *Service) AddTransactionService(data *model.AddTransaction) error {
	return s.Repo.AddTransactionRepo(data)
}
