package service

import (
	"be-golang-chapter-19/repository-api-pattern/model"
	"be-golang-chapter-19/repository-api-pattern/repository"
	"encoding/json"
	"errors"
	"fmt"
)

type CustomerService struct {
	RepoCustomer repository.CustomerRepositoryDB
}

func NewCustomerService(repo repository.CustomerRepositoryDB) *CustomerService {
	return &CustomerService{RepoCustomer: repo}
}

func (cs *CustomerService) InputDataCustomer(username string, password string, email string) error {
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	customer := model.Customer{
		Username: username,
		Password: password,
		Email:    email,
	}

	err := cs.RepoCustomer.Create(&customer)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("berhasil input data customer dengan id ", customer.ID)

	return nil
}

func (cs *CustomerService) GetAllDataUser() error {

	customers, err := cs.RepoCustomer.GetAll()

	if err != nil {
		return err
	}

	if customers == nil || len(*customers) == 0 {
		fmt.Println("No customers found")
		return err
	}

	jsonData, err := json.MarshalIndent(*customers, " ", "")

	if err != nil {
		return err
	}

	fmt.Println("Data user :", string(jsonData))

	return nil

}

func (cs *CustomerService) LoginService(user model.Customer) (*model.Response, error) {

	customers, err := cs.RepoCustomer.GetCustumerLogin(user)

	if err != nil {
		return nil, err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       customers,
	}
	return &response, nil
}
