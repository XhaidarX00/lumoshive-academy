package handler

import (
	"context"
	"database/sql"
	"inventory/service"
	"inventory/utils"
	"inventory/view"
	"time"
)

func MainMenu(db *sql.DB, userService *service.UserService) {
	role := Login(db, userService)
	if role {
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
				excMenu(db)
			}
		}
	}
}

func excMenu(db *sql.DB) {
	funcDisplayName := view.ChoiceMenu("admin")
	if funcDisplayName == nil {
		utils.ErrorMessage("Menu tidak ditemukan!!")
		utils.ExitProgram()
	}

	index := utils.ChoseMenu(funcDisplayName)
	if index == 0 {
		utils.ExitProgram()
	}

	switch index {
	case 1:
		if !service.ManageProduct(db) {
			return
		}
	case 2:
		if !service.ManageCategory(db) {
			return
		}
	case 3:
		if !service.ManageLocations(db) {
			return
		}
	case 4:
		if !service.SearchBy(db) {
			return
		}
	case 5:
		if !service.ManageTrxHistory(db) {
			return
		}

	default:
		utils.ErrorMessage("pilihan invalid!")
		return
	}
}
