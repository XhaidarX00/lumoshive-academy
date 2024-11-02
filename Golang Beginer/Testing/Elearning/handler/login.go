package handler

import (
	"database/sql"
	"elearning/model"
	"elearning/service"
	"elearning/utils"
	"elearning/view"
	"fmt"
	"strconv"
)

func Login(db *sql.DB, userService *service.UserService) string {
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
			utils.ConvertJsonToStruct(&user, "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/dataLoginStudent.json")

			result, role, err := userService.LoginService(user)
			if err != nil {
				// utils.ErrorMessage(err.Error())
				utils.ErrorMessage("data login invalid!")
				return ""
			}

			utils.DisplayDataJson(result)

			if role != "" {
				utils.SuccesMessage("Sukses Login Sebagai " + utils.Capitalize(role))
			}
			return role

		case 2:
			utils.ExitProgram()
		default:
			utils.ErrorMessage("Pilihan Tidak Valid!!")
		}
	}
}
