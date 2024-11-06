package service

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"be-golang-chapter-21/impleme-http-serve/repository"
)

type CustomerService struct {
	RepoCustomer repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return CustomerService{RepoCustomer: repo}
}

func (cs *CustomerService) LoginService(customer model.Customer) error {

	err := cs.RepoCustomer.Login(&customer)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CustomerService) CustomerByID(id int) (*model.Customer, error) {

	customer, err := cs.RepoCustomer.CustomerByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
