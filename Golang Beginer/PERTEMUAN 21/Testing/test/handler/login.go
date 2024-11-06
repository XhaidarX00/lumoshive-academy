package handler

import (
	"encoding/json"
	"main/model"
	"main/service"
	"main/utils"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	if authService == nil {
		panic("auth service is required")
	}
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		utils.WriteErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi request
	if loginReq.Username == "" || loginReq.Password == "" {
		utils.WriteErrorResponse(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	success, err := h.authService.Login(&loginReq)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "pengguna tidak terdaftar" || err.Error() == "invalid credentials" {
			statusCode = http.StatusUnauthorized
		}
		response := model.LoginResponse{
			Success: success,
			Message: err.Error(),
		}

		utils.WriteJSONResponse(w, response, statusCode)
		return
	}

	response := model.LoginResponse{
		Success: success,
		Message: "Login successful",
	}

	utils.WriteJSONResponse(w, response, http.StatusOK)
}
