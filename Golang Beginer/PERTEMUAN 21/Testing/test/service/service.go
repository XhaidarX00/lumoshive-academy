package service

import (
	"errors"
	"main/model"
	"main/repository"
)

type AuthService interface {
	Login(req *model.LoginRequest) (bool, error)
	GetUserById(id string) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	if userRepo == nil {
		panic("user repository is required")
	}

	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) GetUserById(id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("id tidak boleh kosong!")
	}

	user, err := s.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *authService) Login(req *model.LoginRequest) (bool, error) {
	if req == nil {
		return false, errors.New("request cannot be nil")
	}

	if req.Username == "" || req.Password == "" {
		return false, errors.New("username and password are required")
	}

	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, errors.New("pengguna tidak terdaftar")
	}

	if err := valid(user, req); err != nil {
		return false, err
	}

	return true, nil
}

func valid(user *model.User, req *model.LoginRequest) error {
	if user.Username == req.Username {
		if user.Password == req.Password {
			return nil
		} else {
			return errors.New("Pasword Salah")
		}
	} else {
		return nil
	}
}
