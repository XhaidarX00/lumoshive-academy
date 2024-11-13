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
		library.ResponseToJson(w, err.Error(), nil)
		return
	}

	search := r.URL.Query().Get("date")

	sort := r.URL.Query().Get("sort")
	if sort == "" {
		t.logger.Error("Error TravelController", zap.String("Error get sort", sort))
	}

	var data []model.ResponseDataPage
	err = t.Service.GetPageDataService(&data, search, sort, page)
	if err != nil {
		t.logger.Error("Error TravelController", zap.Error(err))
		return
	}

	library.ResponseToJson(w, "Berhasil Get Data", data)
}
