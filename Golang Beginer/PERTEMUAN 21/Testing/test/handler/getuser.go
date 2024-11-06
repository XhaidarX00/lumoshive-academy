package handler

import (
	"main/model"
	"main/utils"
	"net/http"
)

func (h *AuthHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	id := query.Get("id")

	// Memeriksa dan menggunakan nilai query parameter
	if id == "" {
		id = "Guest"
	}

	data, err := h.authService.GetUserById(id)
	if err != nil {
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get ID successful",
		Data:       data,
	}

	utils.WriteJSONResponse(w, response, http.StatusOK)

	// curl -X GET "http://localhost:8080/getid?id=usr1"
}
