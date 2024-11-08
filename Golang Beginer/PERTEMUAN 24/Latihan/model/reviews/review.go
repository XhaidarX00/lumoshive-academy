package reviews

import "time"

type Review struct {
	ID          string
	Order_id    string
	Book_id     string
	Customer_id int
	Rating      float32
	Review_text string
	Review_date time.Time
}
