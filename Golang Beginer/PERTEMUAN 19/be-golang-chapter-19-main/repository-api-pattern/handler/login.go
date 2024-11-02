package handler

import (
	"be-golang-chapter-19/repository-api-pattern/model"
	"be-golang-chapter-19/repository-api-pattern/repository"
	"be-golang-chapter-19/repository-api-pattern/service"
	"database/sql"
	"encoding/json"
	"fmt"

	"io"
	"os"
)

func Login(db *sql.DB) {

	user := model.Customer{}
	file, err := os.Open("body.json")

	if err != nil {
		fmt.Println("Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}

	repo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(repo)

	result, err := customerService.LoginService(user)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	jsonData, err := json.MarshalIndent(result, " ", "")

	if err != nil {
		fmt.Println("err :", err)
	}

	fmt.Println(string(jsonData))
}
