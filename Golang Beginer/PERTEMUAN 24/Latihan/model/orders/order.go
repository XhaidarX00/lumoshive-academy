package orders

type Order struct {
	ID              string
	Customer_id     int
	Name_customer   string
	Payment_methode string
	Total_amount    int
	Discount        float64
	Final_amount    float64
	Status          string
	OrderDate       string
}

type Order_Item struct {
	ID       string
	Order_id string
	Book_id  string
	Quantity int
	Subtotal int
}
