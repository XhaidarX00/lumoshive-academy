package main

import (
	"elearning/database"
	"elearning/handler"
	"elearning/repository"
	"elearning/service"
	"elearning/utils"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		utils.ErrorMessage(err.Error())
		utils.ExitProgram()
	}

	defer db.Close()

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	handler.MainMenu(db, userService)
}
