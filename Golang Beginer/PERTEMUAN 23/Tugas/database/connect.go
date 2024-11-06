package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=postgres dbname=webgolang sslmode=disable password=@Dardar777 host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	log.Println("Database connection established successfully")
	DB = db
}
