package main

import (
	"inventory/database"
	"inventory/handler"
	"inventory/repository"
	"inventory/service"
	"inventory/utils"
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
