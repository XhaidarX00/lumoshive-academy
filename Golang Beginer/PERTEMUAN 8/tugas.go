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
	deadline := time.Now().Add(300 * time.Second)
	ctxWithCancel, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	product := product.BankProductManager()
	idCostumer, address := LoginHandler(ctx)
	ctxWithValue := context.WithValue(ctx, "IDCOSTUMER", idCostumer)
	ctxWithValue = context.WithValue(ctxWithValue, "ADDRESS", address)

	menu(ctxWithCancel, product, ctxWithValue)

}

func LoginHandler(ctx context.Context) (string, string) {
	var username string
	var password string
	var address string
	account := auth.BankAccountManager()

	for {
		fmt.Println("=---------------= 🔑 Login 🔑 =---------------=")
		fmt.Print("Masukan Username : ")
		fmt.Scan(&username)
		fmt.Print("Masukan Paswword : ")
		fmt.Scan(&password)
		fmt.Print("Masukan Alamat : ")
		fmt.Scan(&address)
		utils.ClearScreen()

		ctxWithValue := context.WithValue(ctx, "adminUn", username)
		ctxWithValue = context.WithValue(ctxWithValue, "adminPw", password)

		idCostumer := account.AuthLogin(ctxWithValue, "adminUn", "adminPw")
		if idCostumer != "" {
			return idCostumer, address
		}
	}

}

// Menu
func menu(ctx context.Context, product *product.BankProducts, ctxWithValue context.Context) {

	for {
		select {
		case <-ctx.Done():
			utils.ClearScreen()
			utils.ErrorMessage("Akses ditolak: waktu sesi habis, silahkan login kembali!")

			idCostumer, address := LoginHandler(ctx)

			ctxWithValue = context.WithValue(ctx, "IDCOSTUMER", idCostumer)
			ctxWithValue = context.WithValue(ctxWithValue, "ADDRESS", address)

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
