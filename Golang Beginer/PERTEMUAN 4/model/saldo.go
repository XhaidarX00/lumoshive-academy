package model

import "errors"

type Saldo struct {
	Saldo   int
	Account Account
}

func NewSaldo(account Account) Saldo {
	return Saldo{0, account}
}

func (saldo *Saldo) DebitSaldo(nominal int) error {
	if nominal == 0 {
		return errors.New("Nominal Tidak Boleh 0")
	}
	saldo.Saldo += nominal
	return nil
}

func (saldo *Saldo) CreditSaldo(nominal int) error {
	if nominal == 0 {
		return errors.New("Nominal Tidak Boleh 0")
	}
	saldo.Saldo -= nominal
	return nil
}
