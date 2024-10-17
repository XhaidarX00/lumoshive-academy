package model

type Account struct {
	id    int
	Name  string
	Email string
	Phone string
}

func NewAccount(id int, name string, email string, phone string) Account {
	return Account{id, name, email, phone}
}

func AddAccount(account ...Account) []Account {
	slice_account := []Account{}

	slice_account = append(slice_account, account...)
	return slice_account
}
