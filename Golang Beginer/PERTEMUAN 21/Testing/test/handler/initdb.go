package handler

import (
	"main/database"
	"net/http"
)

func InitDBHandler(w http.ResponseWriter) {
	db, err := database.InitDB()
	if err != nil {
		return
	}

	defer db.Close()

}
