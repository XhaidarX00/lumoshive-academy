package model

import "time"

type ResponseDataPage struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Photo_url   string  `json:"photo_url"`
	Price       int     `json:"price"`
	Date        string  `json:"date_event"`
	People      int     `json:"people"`
	Rating      float32 `json:"rating"`
}

type PLace struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Photo_url   string    `json:"photo_url"`
	Price       int       `json:"price"`
	Created_at  time.Time `json:"create_at"`
}

type Review struct {
	ID         int       `json:"id"`
	Event_id   int       `json:"event_id"`
	Rating     float32   `json:"rating"`
	ReviewDate time.Time `json:"review_date"`
}

type Event struct {
	ID        int       `json:"id"`
	Place_id  int       `json:"place_id"`
	DateEvent time.Time `json:"date_event"`
}

type Transaction struct {
	ID          int       `json:"id"`
	Event_id    string    `json:"evemt_id"`
	StatusOrder bool      `json:"status_order"`
	Created_at  time.Time `json:"create_at"`
}
