package repository

import "database/sql"

type SqlRepo[T any] struct {
	DB *sql.DB
}

func NewSqlRepo[T any](db *sql.DB) SqlRepo[T] {
	return SqlRepo[T]{
		DB: db,
	}
}

func (r *SqlRepo[T]) Create(entity T, insertQuery string, arg ...interface{}) (T, error) {
	_, err := r.DB.Exec(insertQuery, arg...)
	return entity, err
}
