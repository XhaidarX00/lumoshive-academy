package service

// func CreateCustomer() (*model.Customer, error) {
// 	db, _ := database.ConnectDB()
// 	userRepo := repository.NewSqlRepo[model.Customer](db)
// 	NewCustomer := model.Customer{
// 		Name:        "Iskandar",
// 		Username:    "CST1",
// 		Password:    "pass123",
// 		PhoneNumber: "customer1@gmail.com",
// 	}

// 	insertQuery := "INSERT INTO customers (name, username, password, phone_number) VALUES ($1, $2, $3, $4)"
// 	customer, err := userRepo.Create(NewCustomer, insertQuery, NewCustomer.Name, NewCustomer.Username, NewCustomer.Password, NewCustomer.PhoneNumber)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &customer, nil
// }
