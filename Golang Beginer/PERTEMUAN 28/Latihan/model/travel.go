package model

import "time"

type ResponseDataPage struct {
	ID          int     `json:"id"`
	Event_id    int     `json:"event_id"`
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
	Event_id    string    `json:"event_id"`
	StatusOrder bool      `json:"status_order"`
	Created_at  time.Time `json:"create_at"`
}

type AddTransaction struct {
	ID           int    `json:"id"`
	Name         string `validate:"required"`
	Event_id     int    `validate:"required"`
	Email        string `validate:"required,email"`
	ConfirmEmail string `validate:"required,email"`
	Phone        string `validate:"required"` // e164 adalah format internasional untuk nomor telepon
	// Phone        string `validate:"required,e164"` // e164 adalah format internasional untuk nomor telepon
	Message      string `validate:"required"`
	Status_order bool   `validate:"required"`
}

type ResponsePlaceDetail struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       int                `json:"price"`
	Date        string             `json:"date_event"`
	People      int                `json:"people"`
	Rating      float32            `json:"rating"`
	RatingCount float32            `json:"rating_count"`
	Photo_url   []PhotoDetailPlace `json:"photo_url"`
}

type PhotoDetailPlace struct {
	ID          int    `json:"id"`
	Photo_url   string `json:"photo_url"`
	Description string `json:"description"`
}
