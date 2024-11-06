package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=webgolang sslmode=disable password=@Dardar777 host=localhost"
	db, err := sql.Open("postgres", connStr)

	return db, err
}
