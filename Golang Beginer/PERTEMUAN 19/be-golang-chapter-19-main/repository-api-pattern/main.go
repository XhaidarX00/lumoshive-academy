package main

import (
	"be-golang-chapter-19/repository-api-pattern/database"
	"be-golang-chapter-19/repository-api-pattern/handler"
	"fmt"
	"log"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var endpoint string
	fmt.Print("masukkan enpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "login":
		handler.Login(db)
	}
}
