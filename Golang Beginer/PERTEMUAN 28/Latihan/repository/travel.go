package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"latihan/library"
	"latihan/model"
	"latihan/model/response"
)

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (t *Repo) GetDataPageRepo(searchDate string, sort string, page int) (response.PaginationResponse, error) {
	limit := 6
	offset := (page - 1) * limit

	query := `
	SELECT 
		p.id, 
		e.id AS event_id,
		p.name, 
		p.description,
		p.price,
		e.date_event,
		SUM(CASE WHEN t.status_order = TRUE THEN 1 ELSE 0 END) AS people,
		COALESCE(ROUND(AVG(r.rating), 1), 0) AS rating
	FROM event e
	JOIN place p ON e.place_id = p.id
	JOIN transaction t ON t.event_id = e.id
	JOIN reviews r ON r.transaction_id = t.id`

	if searchDate != "" {
		query = fmt.Sprintf("%s WHERE DATE(e.date_event) = $1 GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event ORDER BY e.id", query)
	} else if sort == "low_to_high" {
		query = fmt.Sprintf("%s GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event ORDER BY p.price ASC LIMIT $1 OFFSET $2", query)
	} else if sort == "high_to_low" {
		query = fmt.Sprintf("%s GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event ORDER BY p.price DESC LIMIT $1 OFFSET $2", query)
	} else {
		query = fmt.Sprintf("%s GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event ORDER BY p.id LIMIT $1 OFFSET $2", query)
	}

	var rows *sql.Rows
	var err error

	if searchDate != "" {
		rows, err = t.DB.Query(query, searchDate)
	} else {
		rows, err = t.DB.Query(query, limit, offset)
	}
	if err != nil {
		return response.PaginationResponse{}, err
	}

	defer rows.Close()

	var data []model.ResponseDataPage
	for rows.Next() {
		var response_ model.ResponseDataPage
		if err := rows.Scan(&response_.ID, &response_.Event_id, &response_.Name, &response_.Description, &response_.Price, &response_.Date, &response_.People, &response_.Rating); err != nil {
			return response.PaginationResponse{}, err
		}

		data = append(data, response_)
	}

	if len(data) == 0 {
		return response.PaginationResponse{}, errors.New("data tidak ditemukan")
	}

	var totalItems int
	err = t.DB.QueryRow("SELECT COUNT(*) FROM place").Scan(&totalItems)
	if err != nil {
		return response.PaginationResponse{}, err
	}

	totalPage := (totalItems + limit - 1) / limit

	result := library.PageResponse(
		"Berhasil Mendapatkan Data",
		limit,
		page,
		totalItems,
		totalPage,
		data,
	)

	return result, nil

}

func (t *Repo) PlaceDetailRepo(data *model.ResponsePlaceDetail, id int) error {
	query := `
	SELECT 
		p.id, 
		p.name, 
		p.description,
		p.price,
		e.date_event,
		SUM(CASE WHEN t.status_order = TRUE THEN 1 ELSE 0 END) AS people,
		COALESCE(ROUND(AVG(r.rating), 1), 0) AS rating,
		COUNT(r.rating) AS rating_count
	FROM event e
	JOIN place p ON e.place_id = p.id
	JOIN transaction t ON t.event_id = e.id
	LEFT JOIN reviews r ON r.transaction_id = t.id
	WHERE p.id = $1
	GROUP BY p.id, p.name, p.description, p.price, e.date_event
	ORDER BY p.id
	`

	err := t.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Name,
		&data.Description,
		&data.Price,
		&data.Date,
		&data.People,
		&data.Rating,
		&data.RatingCount,
	)

	if err != nil {
		return err
	}

	query2 := `
	SELECT 
		g.id, 
		g.photo_url,
		g.description
	FROM gallery g
	JOIN place p ON g.place_id = p.id
	WHERE p.id = $1
	GROUP BY g.id, g.photo_url, g.description
	ORDER BY g.id
	`

	rows, err := t.DB.Query(query2, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var dataPhotoUrl []model.PhotoDetailPlace
	for rows.Next() {
		var photo_url model.PhotoDetailPlace
		if err := rows.Scan(&photo_url.ID, &photo_url.Photo_url, &photo_url.Description); err != nil {
			return err
		}

		dataPhotoUrl = append(dataPhotoUrl, photo_url)
	}

	data.Photo_url = dataPhotoUrl

	return nil
}

func (t *Repo) LocationByIdRepo(data *model.Locations) error {
	query := `
	SELECT 
		id, 
		loglat, 
		description
	FROM locations
	WHERE place_id = $1`

	if err := t.DB.QueryRow(query, data.ID).Scan(&data.ID, &data.Longlat, &data.Description); err != nil {
		return err
	}

	return nil
}

func (t *Repo) TourPLanByIdRepo(datas *[]model.TourPLan, tour_id int) error {
	query := `
	SELECT 
		tp.plan_id, 
		tp.title, 
		tp.day_number, 
		tp.description, 
		tp.meals, 
		tp.accommodation
	FROM tour_plan tp
	JOIN tours t ON t.tour_id = tp.tour_id
	JOIN event e ON e.id = t.event_id
	WHERE tp.tour_id = $1`

	rows, err := t.DB.Query(query, tour_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var data model.TourPLan
		if err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Day_number,
			&data.Description,
			&data.Meals,
			&data.Accommodation,
		); err != nil {
			return err
		}

		*datas = append(*datas, data)
	}

	return nil
}
