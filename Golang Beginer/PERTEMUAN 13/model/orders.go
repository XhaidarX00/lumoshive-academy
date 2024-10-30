package model

import "time"

type Orders struct {
	Order_id         uint16
	Customer_id      uint16
	Driver_id        uint16
	Pickup_location  uint16
	Dropoff_location uint16
	Total_fare       float64
	Order_date       time.Time
}
