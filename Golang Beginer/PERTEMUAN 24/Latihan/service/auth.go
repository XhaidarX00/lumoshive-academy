package service

import (
	"latihan/library"
	"latihan/model/customers"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (s *Service) RegisterService(user *customers.Customer) error {
	err := s.Repo.RegisterRepo(user)
	if err != nil {
		s.Logger.Error("Error Service :", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) LoginService(user *customers.Customer) error {
	err := s.Repo.LoginRepo(user)
	if err != nil {
		s.Logger.Error("Error Service :", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) TokenCheck(token string) string {
	err := s.Repo.TokenCheckRepo(token)
	if err != "" {
		s.Logger.Error("Error :", zap.String("Service", "Token Tidak Ditemukan"))
		return err
	}

	return ""
}

// Fungsi untuk membersihkan token yang sudah kadaluarsa
func (s *Service) CleanExpiredTokens(w http.ResponseWriter) bool {
	err := s.Repo.CleanExpiredTokensRepo()
	if err != "" {
		s.Logger.Error("Error :", zap.String("Service", "Token Tidak Ditemukan"))
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
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := s.Repo.CleanExpiredTokensRepo(); err != "" {
				s.Logger.Error("Error :", zap.String("Service", "Token Tidak Ditemukan"))
				return
			}
		}
	}
}

func (s *Service) GetRoleService(token string) (string, error) {
	return s.Repo.GetRoleRepo(token)
}

func (s *Service) GetCustomerByIDService(id int) (string, error) {
	return s.Repo.GetCustomerByIDRepo(id)
}
