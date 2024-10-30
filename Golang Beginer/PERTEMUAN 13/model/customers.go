package model

import "time"

type Customers struct {
	Customer_id  uint16
	Name         string
	Email        string
	Phone_number string
	Created_at   time.Time
}
