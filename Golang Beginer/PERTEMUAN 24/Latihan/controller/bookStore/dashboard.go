package bookstore

import (
	"latihan/controller"
	"latihan/model"
	"latihan/service"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	var data model.GetDhasboardData
	err := service.ServiceF.GetDhasboardDataService(&data)
	if err != nil {
		controller.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Total_Books":    data.Total_Books,
		"Total_Sales":    data.Total_Sales,
		"Highest_Rating": data.Highest_Rating,
	}

	controller.RenderTemplate(w, "dashboard.html", result)

}
