package auth

import (
	"context"
	"fmt"
	"main/utils"

	"github.com/google/uuid"
)

type Account struct {
	IDCOSTUMER string
	USERNAME   string
	PASSWORD   string
}

type BankAccounts struct {
	Accounts []Account
}

func BankAccountManager() *BankAccounts {
	bankAccounts := &BankAccounts{
		Accounts: []Account{},
	}
	bankAccounts.initializationAccounts()
	return bankAccounts
}

func (b *BankAccounts) initializationAccounts() {
	b.Accounts = append(b.Accounts, Account{IDCOSTUMER: uuid.New().String(), USERNAME: "haidar", PASSWORD: "admin123"})
	b.Accounts = append(b.Accounts, Account{IDCOSTUMER: uuid.New().String(), USERNAME: "client", PASSWORD: "client123"})
}

func (b *BankAccounts) DisplayAccounts() {
	for _, account := range b.Accounts {
		fmt.Printf("ID: %s, Username: %s, Password: %s\n", account.IDCOSTUMER, account.USERNAME, account.PASSWORD)
	}
}

func (b *BankAccounts) AuthLogin(ctxv context.Context, keyUn string, keyPw string) string {
	if len(b.Accounts) == 0 {
		b.initializationAccounts()
	}

	un, okUn := ctxv.Value(keyUn).(string)
	if !okUn {
		utils.ErrorMessage("Username tidak valid")
		return ""
	}

	pw, okPw := ctxv.Value(keyPw).(string)
	if !okPw {
		utils.ErrorMessage("Password tidak valid")
		return ""
	}

	fmt.Printf("%s %s\n", un, pw)

	for _, account := range b.Accounts {
		if account.USERNAME == un {
			if account.PASSWORD == pw {
				utils.SuccesMessage("Login sukses. Selamat Berbelanja !!")
				fmt.Println(account.IDCOSTUMER)
				return account.IDCOSTUMER
			} else {
				utils.ErrorMessage("Password salah")
				return ""
			}
		}
	}

	utils.ErrorMessage("Username tidak ditemukan")
	return ""
}

func (b *BankAccounts) CheckAccount(un string, pw string) bool {
	for _, account := range b.Accounts {
		if account.USERNAME == un {
			if account.PASSWORD == pw {
				return true
			} else {
				return false
			}
		}
	}

	return false
}
