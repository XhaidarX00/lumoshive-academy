package controller

import (
	"encoding/json"
	"fmt"
	"latihan/library"
	"latihan/model"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func (t *Travel) AddTransactionController(w http.ResponseWriter, r *http.Request) {
	var data model.AddTransaction
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		library.ResponseToJson(w, "Invalid request body", nil)
		return
	}

	validate := validator.New()

	err := validate.VarWithValue(data.ConfirmEmail, data.Email, "eqfield")
	if err != nil {
		t.logger.Error("Error PlaceDetailController", zap.Error(err))
		library.ResponseToJson(w, "Email tidak sama", nil)
		return
	}

	err = validate.Struct(&data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: Field %s failed validation with tag %s\n", err.Field(), err.Tag())
		}

		library.ResponseToJson(w, "Kesalahan Input Data", nil)
		return
	}

	err = t.Service.AddTransactionService(&data)
	if err != nil {
		t.logger.Error("Error PlaceDetailController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil menambahkan transaction", data)
}
