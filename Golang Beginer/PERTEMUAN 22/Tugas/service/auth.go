package service

import (
	"main/library"
	"main/model"
	"net/http"
	"time"
)

func (s *Service) RegisterService(user *model.Users) error {
	err := s.Repo.RegisterRepo(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) TokenCheck(token string, w http.ResponseWriter) bool {
	err := s.Repo.TokenCheckRepo(token)
	if err != "" {
		response := library.UnauthorizedRequest(err)
		library.JsonResponse(w, response)
		return false
	}

	return true
}

// Fungsi untuk membersihkan token yang sudah kadaluarsa
func (s *Service) CleanExpiredTokens(w http.ResponseWriter) bool {
	err := s.Repo.CleanExpiredTokensRepo()
	if err != "" {
		response := library.UnauthorizedRequest(err)
		library.JsonResponse(w, response)
		return false
	}

	return true
}

func (s *Service) RoleCheckAcc(role string, w http.ResponseWriter, r *http.Request) bool {
	// Check role permissions based on HTTP method
	switch role {
	case "dev":
		if r.Method != http.MethodPut {
			library.ResponseToJson(w, "Forbidden: Only 'PUT' method is allowed for dev role", nil)
			return false
		}

		return true
	case "admin":
		return true
	default:
		library.ResponseToJson(w, "Forbidden: Unrecognized role", nil)
		return false
	}
}

func (s *Service) CheckToken() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := s.Repo.CleanExpiredTokensRepo(); err != "" {
				return
			}
		}
	}
}

func (s *Service) GetRoleService(token string) (string, error) {
	return s.Repo.GetRoleRepo(token)
}
