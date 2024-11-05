package service

import (
	"database/sql"
	"main/repository"
	"net/http"
)

type Service struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{Repo: repo}
}

func (s *Service) TokenCheckService(db *sql.DB, token string, w http.ResponseWriter) bool {
	if s.Repo.TokenCheckRepo(db, token, w) {
		return true
	}
	return s.cleanExpiredTokensService(db, w)
}

func (s *Service) cleanExpiredTokensService(db *sql.DB, w http.ResponseWriter) bool {
	return s.Repo.CleanExpiredTokensRepo(db, w)
}

func (s *Service) RoleCheckAccService(role string, w http.ResponseWriter, r *http.Request) bool {
	return s.Repo.RoleCheckAccRepo(role, w, r)
}
