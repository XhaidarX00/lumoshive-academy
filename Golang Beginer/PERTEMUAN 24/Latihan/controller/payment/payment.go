// handler/payment_handler.go
package payments

import (
	"database/sql"
	"encoding/json"
	"io"
	"latihan/library"
	"latihan/model/payment"
	"latihan/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

var Svc *service.Service

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	// Tentukan skema (http atau https)
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	domain := scheme + r.Host

	// Ambil file dan data dari form
	file, data, err := r.FormFile("photo")
	if err != nil {
		library.ResponseToJson(w, err.Error(), nil)
		return
	}
	defer file.Close()

	// Simpan foto di folder asset
	dst, err := os.Create(filepath.Join("asset/", data.Filename))
	if err != nil {
		library.ResponseToJson(w, err.Error(), nil)
		return
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	// Ambil nilai-nilai lainnya dari form
	name := r.FormValue("name")
	is_active := r.FormValue("is_active")
	is_activeBool, err := strconv.ParseBool(is_active)
	if err != nil {
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	// Waktu saat ini untuk CreatedAt dan UpdatedAt
	currentTime := time.Now()

	// Foto URL
	photo_url := strings.Join([]string{domain, "/asset/", data.Filename}, "")

	// Membuat objek Payment dengan data yang diterima
	payment := payment.Payment{
		Name:      name,
		Photo:     sql.NullString{String: photo_url, Valid: photo_url != ""},
		IsActive:  is_activeBool,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		DeletedAt: nil,
	}

	// Simpan data pembayaran di database
	if err := service.ServiceF.Repo.Create(&payment); err != nil {
		response := library.InternalServerError("Gagal Menambahkan Payment")
		library.JsonResponse(w, response)
		return
	}

	// Kirimkan respons sukses
	library.ResponseToJson(w, "Berhasil Menambahkan Payment", payment)
}

func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := Svc.Repo.GetAll()
	if err != nil {
		response := library.InternalServerError(err.Error())
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Get Payment", payments)
}

func GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payment, err := service.ServiceF.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var payment payment.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payment.ID = id
	if err := service.ServiceF.Repo.Update(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := service.ServiceF.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
