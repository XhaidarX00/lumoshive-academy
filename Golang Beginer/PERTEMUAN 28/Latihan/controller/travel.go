package controller

import (
	"latihan/library"
	"latihan/model"
	"latihan/service"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type Travel struct {
	Service *service.Service
	logger  *zap.Logger
}

func NewTravelHandelr(serv *service.Service, log *zap.Logger) *Travel {
	return &Travel{
		Service: serv,
		logger:  log,
	}
}

func (t *Travel) TravelController(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		t.logger.Error("Error TravelController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	searchDate := r.URL.Query().Get("date")
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		t.logger.Error("Error TravelController", zap.String("Error get sort", sort))
	}

	result, err := t.Service.GetPageDataService(searchDate, sort, page)
	if err != nil {
		t.logger.Error("Error TravelController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	library.JsonResponse(w, result)
}

func (t *Travel) PlaceDetailController(w http.ResponseWriter, r *http.Request) {
	event_id, err := strconv.Atoi(r.URL.Query().Get("event_id"))
	if err != nil {
		t.logger.Error("Error PlaceDetailController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	var data model.ResponsePlaceDetail
	err = t.Service.PlaceDetailService(&data, event_id)
	if err != nil {
		t.logger.Error("Error TravelController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Mendapatkan Data", data)
}

func (t *Travel) GetLocationByIdController(w http.ResponseWriter, r *http.Request) {
	place_id, err := strconv.Atoi(r.URL.Query().Get("place_id"))
	if err != nil {
		t.logger.Error("Error PlaceDetailController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	var data model.Locations
	data.ID = place_id
	err = t.Service.LocationByIdService(&data)
	if err != nil {
		t.logger.Error("Error GetLocationByIdController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Mendapatkan Location", data)
}

func (t *Travel) GetTourPLanByIdController(w http.ResponseWriter, r *http.Request) {
	tour_id, err := strconv.Atoi(r.URL.Query().Get("tour_id"))
	if err != nil {
		t.logger.Error("Error PlaceDetailController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	var data []model.TourPLan
	err = t.Service.TourPLanByIdService(&data, tour_id)
	if err != nil {
		t.logger.Error("Error GetTourPLanByIdController", zap.Error(err))
		response := library.NotFoundRequest("Data Tidak Ditemukan")
		library.JsonResponse(w, response)
		return
	}

	library.ResponseToJson(w, "Berhasil Mendapatkan Tour Plan", data)
}
