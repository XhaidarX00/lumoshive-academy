package main

import (
	"context"
	"fmt"
	"main/auth"
	"main/product"
	"main/utils"
	"os"
	"time"
)

func main() {
	utils.ClearScreen()
	ctx := context.Background()
	deadline := time.Now().Add(10 * time.Second)
	ctxWithCancel, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	product := product.BankProductManager()
	ctxWithValue := LoginHandler(ctx)
	// fmt.Println("Id custom : ", idCostumer)
	// ctxWithValue := context.WithValue(ctx, "IDCOSTUMER", idCostumer)
	// ctxWithValue = context.WithValue(ctxWithValue, "ADDRESS", address)

	menu(ctxWithCancel, product, ctxWithValue)

}

func LoginHandler(ctx context.Context) context.Context {
	var username string
	var password string
	var address string
	account := auth.BankAccountManager()

	for {
		account.DisplayAccounts()
		fmt.Println("=---------------= 🔑 Login 🔑 =---------------=")
		fmt.Print("Masukan Username : ")
		fmt.Scan(&username)
		fmt.Print("Masukan Paswword : ")
		fmt.Scan(&password)
		fmt.Print("Masukan Alamat : ")
		fmt.Scan(&address)
		utils.ClearScreen()

		UN, okUN := ctx.Value(fmt.Sprintf("Un%s", username)).(string)
		PW, okPW := ctx.Value(fmt.Sprintf("Pw%s", password)).(string)

		if !okUN || !okPW {
			if !account.CheckAccount(UN, PW) {
				ctxWithValue := context.WithValue(ctx, fmt.Sprintf("Un%s", username), username)
				ctxWithValue = context.WithValue(ctxWithValue, fmt.Sprintf("Pw%s", password), password)

				idCostumer := account.AuthLogin(ctxWithValue, fmt.Sprintf("Un%s", username), fmt.Sprintf("Pw%s", password))
				fmt.Println("Id custom : ", idCostumer)

				if idCostumer != "" {
					ctxWithValue = context.WithValue(ctx, "IDCOSTUMER", idCostumer)
					ctxWithValue = context.WithValue(ctxWithValue, "ADDRESS", address)
					return ctxWithValue
				}

			} else {
				return nil
			}

		}

		return nil
	}

}

// Menu
func menu(ctx context.Context, product *product.BankProducts, ctxWithValue context.Context) {

	for {
		select {
		case <-ctx.Done():
			utils.ClearScreen()
			utils.ErrorMessage("Akses ditolak: waktu sesi habis, silahkan login kembali!")

			ctxWithCancel, cancel := resetSessionTimeout()
			defer cancel()

			ctx = ctxWithCancel

			LoginHandler(ctxWithValue)
			fmt.Printf("Data Context : %v\n", ctxWithValue)
			// idCostumer, address := LoginHandler(ctx)
			// ctxWithValue = context.WithValue(ctx, "IDCOSTUMER", idCostumer)
			// ctxWithValue = context.WithValue(ctxWithValue, "ADDRESS", address)

			continue
		default:
			var input int
			fmt.Println("+++ ====== Menu ====== +++")
			fmt.Println("1. " + utils.ColorMessage("green", "Daftar Produk"))
			fmt.Println("2. " + utils.ColorMessage("green", "Keranjang"))
			fmt.Println("3. " + utils.ColorMessage("green", "Checkout"))
			fmt.Println("4. " + utils.ColorMessage("green", "Hapus Produk Keranjang"))
			fmt.Println("0. " + utils.ColorMessage("red", "Keluar"))
			fmt.Print("Masukkan nomor menu: ")
			fmt.Scan(&input)
			utils.ClearScreen()

			switch input {
			case 1:
				utils.ClearScreen()
				product.DisplayProduct(ctxWithValue, "IDCOSTUMER")
			case 2:
				utils.ClearScreen()
				product.DisplayChosenProducts(ctxWithValue, "IDCOSTUMER", "ADDRESS")
			case 3:
				utils.ClearScreen()
				product.CheckoutProduct(ctxWithValue, "IDCOSTUMER", "ADDRESS")
			case 4:
				utils.ClearScreen()
				product.RemoveProductFromCarts(ctxWithValue, "IDCOSTUMER", "ADDRESS")
			case 0:
				exitMainmenu()
			default:
				utils.ErrorMessage("Pilihan tidak valid")
			}
		}
	}
}

func exitMainmenu() {
	defer os.Exit(0)
	utils.ClearScreen()
	utils.SuccesMessage("Keluar dari Program\n")
}

func resetSessionTimeout() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	deadline := time.Now().Add(20 * time.Second)
	return context.WithDeadline(ctx, deadline)
}
