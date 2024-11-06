package repository

import (
	"database/sql"
	"errors"
	"main/model"
)

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
	GetById(id string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	// Periksa koneksi database
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	user := &model.User{}
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("pengguna tidak terdaftar")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetById(id string) (*model.User, error) {
	// Periksa koneksi database
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	user := &model.User{}
	query := "SELECT id, username, password FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("pengguna tidak terdaftar")
		}
		return nil, err
	}

	return user, nil
}
