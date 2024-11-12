package bookstore

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/model"
	"net/http"

	"go.uber.org/zap"
)

func (b *Books) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	var data model.GetDhasboardData
	err := b.Service.GetDhasboardDataService(&data)
	if err != nil {
		b.logger.Error("Error Dashboardhandler", zap.Error(err))
		pagehandler.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Total_Books":    data.Total_Books,
		"Total_Sales":    data.Total_Sales,
		"Highest_Rating": data.Highest_Rating,
	}

	pagehandler.RenderTemplate(w, "dashboard.html", result)

}
