package main

import (
	"fmt"

	"main.go/model"
)

func main() {

	// add account
	account_1 := model.NewAccount(1, "Dar1", "hdar@gamil.com", "6282")
	account_2 := model.NewAccount(2, "Dar2", "hdar1@gamil.com", "62822")

	saldo1 := model.NewSaldo(account_1)
	saldo2 := model.NewSaldo(account_2)

	fmt.Println("+++==================================+++")
	sliceSaldo := []model.Saldo{}
	sliceSaldo = append(sliceSaldo, saldo1, saldo2)
	model.PrintDataAdd(sliceSaldo)
	fmt.Println("----------------------------------------")

	// menambahkan saldo
	err := saldo1.DebitSaldo(60)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	} else {
		model.PrintDebit(sliceSaldo, saldo1)
		fmt.Println("----------------------------------------")
	}

	// mengurangi saldo
	errC := saldo1.CreditSaldo(10)
	if errC != nil {
		fmt.Println("Error : ", errC)
		return
	} else {
		model.PrintCredit(sliceSaldo, saldo1)
		fmt.Println("+++==================================+++")
	}
}
