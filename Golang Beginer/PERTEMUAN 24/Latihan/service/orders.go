package service

import "latihan/model/orders"

func (s *Service) GetOrderDataService(data *[]orders.Order) error {
	return s.Repo.GetOrderDataRepo(data)
}
