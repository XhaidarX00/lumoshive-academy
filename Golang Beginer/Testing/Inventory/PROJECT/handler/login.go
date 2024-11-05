package handler

import (
	"database/sql"
	"fmt"
	"inventory/model"
	"inventory/service"
	"inventory/utils"
	"inventory/view"
	"strconv"
)

func Login(db *sql.DB, userService *service.UserService) bool {
	var input string
	for {
		rangeMenu := view.LoginMenu()
		fmt.Scan(&input)
		utils.ClearScreen()

		IntInput, err := strconv.Atoi(input)
		if err != nil || IntInput < 1 || IntInput > rangeMenu {
			utils.ErrorMessage("input tidak valid!")
			continue
		}

		switch IntInput {
		case 1:
			var user model.Users
			utils.ConvertJsonToStruct(&user, "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/dataLoginAdmin.json")

			result, username, err := userService.LoginService(user)
			if err != nil {
				utils.ErrorMessage("data login invalid!")
				return false
			}

			utils.DisplayDataJson(result)

			if username != "" {
				utils.SuccesMessage("Sukses Login Sebagai " + utils.Capitalize(username))
			}

			return true

		case 2:
			utils.ExitProgram()
		default:
			utils.ErrorMessage("Pilihan Tidak Valid!!")
		}
	}
}
