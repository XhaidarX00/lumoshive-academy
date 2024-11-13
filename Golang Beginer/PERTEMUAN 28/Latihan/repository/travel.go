package repository

import (
	"database/sql"
	"fmt"
	"latihan/model"
)

type Travel struct {
	DB *sql.DB
}

func NewTravel(db *sql.DB) *Travel {
	return &Travel{
		DB: db,
	}
}

// queryPlus = fmt.Sprintf("%s WHERE id = $1 GROUP BY p.id, p.name, p.description, p.price, e.date_event", query)
// queryPlus = fmt.Sprintf("%s GROUP BY p.id, p.name, p.description, p.price, e.date_event ORDER BY p.price ASC", query)

func (t *Travel) GetDataPageRepo(data *[]model.ResponseDataPage, search string, sort string, page int) error {
	limit := 6
	offset := (page - 1) * limit

	query := `
	SELECT 
		p.id, 
		p.name, 
		p.description,
		p.price,
		e.date_event,
		SUM(CASE WHEN t.status_order = TRUE THEN 1 ELSE 0 END) AS people,
		COALESCE(ROUND(AVG(r.rating), 1), 0) AS rating
	FROM event e
	JOIN place p ON e.place_id = p.id
	JOIN transaction t ON t.event_id = e.id
	LEFT JOIN reviews r ON r.transaction_id = t.id`

	if search != "" {
		query = fmt.Sprintf("%s WHERE DATE(date_event) = $1 GROUP BY p.id, p.name, p.description, p.price, e.date_event ORDER BY p.price ASC LIMIT $2 OFFSET $3", query)
	} else if sort == "low_to_high" {
		query = fmt.Sprintf("%s GROUP BY p.id, p.name, p.description, p.price, e.date_event ORDER BY p.price ASC LIMIT $1 OFFSET $2", query)
	} else if sort == "high_to_low" {
		query = fmt.Sprintf("%s GROUP BY p.id, p.name, p.description, p.price, e.date_event ORDER BY p.price DESC LIMIT $1 OFFSET $2", query)
	} else {
		query = fmt.Sprintf("%s GROUP BY p.id, p.name, p.description, p.price, e.date_event LIMIT $1 OFFSET $2", query)
	}
	var rows *sql.Rows
	var err error

	if search != "" {
		rows, err = t.DB.Query(query, search, limit, offset)
	} else {
		rows, err = t.DB.Query(query, limit, offset)
	}
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var response model.ResponseDataPage
		if err := rows.Scan(&response.ID, &response.Name, &response.Description, &response.Price, &response.Date, &response.People, &response.Rating); err != nil {
			return err
		}

		*data = append(*data, response)
	}

	return nil
}
