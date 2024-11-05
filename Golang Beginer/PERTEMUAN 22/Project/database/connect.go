package database

import (
	"database/sql"
	"fmt"
	"main/model"
	"main/utils"
	"net/http"

	_ "github.com/lib/pq"
)

func InitDB(w http.ResponseWriter) (*sql.DB, error) {
	connStr := "user=postgres dbname=webgolang sslmode=disable password=@Dardar777 host=localhost"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		response := model.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("failed to open database: %v", err),
		}

		utils.WriteJSONResponseStruct(w, response)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		response := model.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("failed to ping database: %v", err),
		}
		utils.WriteJSONResponseStruct(w, response)
		return nil, err
	}

	// log.Println("Database connection established successfully")
	// response := model.Response{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Database connection established successfully",
	// 	Data:       nil,
	// }

	// utils.WriteJSONResponseStruct(w, response)
	return db, nil
}

func InitDB2() (*sql.DB, error) {
	connStr := "user=postgres dbname=webgolang sslmode=disable password=@Dardar777 host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// log.Println("Database connection established successfully")
	return db, nil
}
