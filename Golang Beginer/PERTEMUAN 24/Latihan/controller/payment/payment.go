// handler/payment_handler.go
package payments

import (
	"database/sql"
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
	"go.uber.org/zap"
)

type Payment struct {
	Service *service.Service
	logger  *zap.Logger
}

func NewPaymenHandelr(serv *service.Service, log *zap.Logger) *Payment {
	return &Payment{
		Service: serv,
		logger:  log,
	}
}

func (p *Payment) CreatePayment(w http.ResponseWriter, r *http.Request) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	domain := scheme + r.Host

	file, data, err := r.FormFile("photo")
	if err != nil {
		p.logger.Error("Error createpaymenthandler", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("asset/", data.Filename))
	if err != nil {
		p.logger.Error("Error createpaymenthandler", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		p.logger.Error("Error createpaymenthandler", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	name := r.FormValue("name")
	is_active := r.FormValue("is_active")
	is_activeBool, err := strconv.ParseBool(is_active)
	if err != nil {
		p.logger.Error("Error createpaymenthandler", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	currentTime := time.Now()

	photo_url := strings.Join([]string{domain, "/asset/", data.Filename}, "")

	payment_ := payment.Payment{
		Name:      name,
		Photo:     sql.NullString{String: photo_url, Valid: photo_url != ""},
		IsActive:  is_activeBool,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		DeletedAt: nil,
	}

	if err := p.Service.Repo.Create(&payment_); err != nil {
		p.logger.Error("Error createpaymenthandler", zap.Error(err))
		response := library.InternalServerError("Gagal Menambahkan Payment")
		library.JsonResponse(w, response)
		return
	}

	result := payment.Payment{
		ID:       payment_.ID,
		Name:     payment_.Name,
		Photo:    payment_.Photo,
		IsActive: payment_.IsActive,
	}
	library.ResponseToJson(w, "Berhasil Menambahkan Payment", result)
}

func (p *Payment) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := p.Service.Repo.GetAll()
	if err != nil {
		p.logger.Error("Error GetAllPayments", zap.Error(err))
		response := library.InternalServerError(err.Error())
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Get Payment", payments)
}

func (p *Payment) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		p.logger.Error("Error GetPaymentByID", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		// p.logger.Error("Gagal mendapatkan Id User",
		// 	zap.String("Name Handler", "Controller.Payment"),
		// 	zap.String("Name Func", "GetPaymentByID"),
		// 	zap.Error(err),
		// )
		return
	}

	payment, err := p.Service.Repo.GetByID(id)
	if err != nil {
		p.logger.Error("Error GetPaymentByID", zap.Error(err))
		// p.logger.Error("Terjadi Kesalahan database",
		// 	zap.String("Name Handler", "Controller.Payment"),
		// 	zap.String("Name Func", "GetPaymentByID"),
		// 	zap.Error(err),
		// )
		response := library.NotFoundRequest(err.Error())
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Get Payment", payment)
}

func (p *Payment) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	domain := scheme + r.Host

	file, data, err := r.FormFile("photo")
	if err != nil {
		p.logger.Error("Error UpdatePayment", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("asset/", data.Filename))
	if err != nil {
		p.logger.Error("Error UpdatePayment", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		p.logger.Error("Error UpdatePayment", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	name := r.FormValue("name")
	is_active := r.FormValue("is_active")
	is_activeBool, err := strconv.ParseBool(is_active)
	if err != nil {
		p.logger.Error("Error UpdatePayment", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	photo_url := strings.Join([]string{domain, "/asset/", data.Filename}, "")

	payment_ := payment.Payment{
		Name:     name,
		IsActive: is_activeBool,
	}

	if photo_url != "" {
		payment_.Photo = sql.NullString{String: photo_url, Valid: photo_url != ""}
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		p.logger.Error("Error UpdatePayment", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	payment_.ID = id
	if err := p.Service.Repo.Update(&payment_); err != nil {
		p.logger.Error("Error UpdatePayment", zap.Error(err))
		response := library.InternalServerError("Gagal Update Payment")
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Update Payment", payment_)
}

func (p *Payment) DeletePayment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		p.logger.Error("Error DeletePayment", zap.Error(err))
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	if err := p.Service.Repo.Delete(id); err != nil {
		p.logger.Error("Error DeletePayment", zap.Error(err))
		response := library.InternalServerError(err.Error())
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Hapus Payment", id)
}
