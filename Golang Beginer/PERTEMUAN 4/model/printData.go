package model

import "fmt"

func PrintDataAdd(saldos []Saldo) {
	fmt.Printf("Akun Berhasil Ditambahkan %+v\n", saldos)
}

func printDataDebit(saldos []Saldo) {
	fmt.Printf("Saldo Berhasil Ditambahkan %+v\n", saldos)
}

func PrintDebit(sliceSaldo []Saldo, saldo ...Saldo) {
	for _, s := range saldo {
		for i, saldo := range sliceSaldo {
			if saldo.Account.id == s.Account.id {
				sliceSaldo[i] = s
			}
		}
	}
	printDataDebit(sliceSaldo)

}

func printDataCredit(saldos []Saldo) {
	fmt.Printf("Saldo Berhasil Dikurangi %+v\n", saldos)
}

func PrintCredit(sliceSaldo []Saldo, saldo ...Saldo) {
	for _, s := range saldo {
		for i, saldo := range sliceSaldo {
			if saldo.Account.id == s.Account.id {
				sliceSaldo[i] = s
			}
		}
	}
	printDataCredit(sliceSaldo)

}
