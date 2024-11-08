package customers

type Customer struct {
	ID           int
	Username     string
	Password     string
	Name         string
	Phone_Number string
	Token        string
}

type Address struct {
	ID     int
	street string
	city   string
	county string
	postal string
}
