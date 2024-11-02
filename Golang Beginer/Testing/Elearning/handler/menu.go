package handler

import (
	"context"
	"database/sql"
	"elearning/service"
	"elearning/utils"
	"elearning/view"
	"time"
)

func MainMenu(db *sql.DB, userService *service.UserService) {
	role := Login(db, userService)
	if role != "" {
		ctx := context.Background()
		deadline := time.Now().Add(300 * time.Second)
		ctxwithdeadline, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		for {
			select {
			case <-ctxwithdeadline.Done():
				utils.ErrorMessage("Akses Ditolak, Sesi Login Expired!")
				ctxWithDeadline, cancel := utils.ResetSessionTimeout(300)
				defer cancel()
				ctxwithdeadline = ctxWithDeadline
				role = Login(db, userService)
				continue

			default:
				excMenu(role, db, userService)
			}
		}
	}
}

func excMenu(role string, db *sql.DB, userService *service.UserService) {
	funcDisplayName := view.ChoiceMenu(role)
	if funcDisplayName == nil {
		utils.ErrorMessage("Menu tidak ditemukan!!")
		utils.ExitProgram()
	}

	index := utils.ChoseMenu(funcDisplayName)
	if index == 0 {
		utils.ExitProgram()
	}

	switch role {
	case "admin":
		if !service.RoleExcAdmin(db, userService, index) {
			return
		}
	case "mentor":
		if !service.RoleExcMentor(db, userService, index) {
			return
		}
	case "student":
		if !service.RoleExcStudent(db, userService, index) {
			return
		}
	}
}
