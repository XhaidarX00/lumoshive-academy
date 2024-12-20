package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=travel sslmode=disable password=@Dardar777 host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	log.Println("Database connection established successfully")
	DB = db

	return db, nil
}
