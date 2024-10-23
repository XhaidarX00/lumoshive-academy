package product

import (
	"context"
	"fmt"
	"main/utils"
	"strconv"
	"strings"
)

type Product struct {
	ID      string
	NAME    string
	PRICE   float64
	ADDRESS string
}

type ChosenProduct struct {
	IDCUSTOMER string
	ID         string
	NAME       string
	PRICE      float64
	ADDRESS    string
}

type BankProducts struct {
	Products          []Product
	tempChosenProduct []ChosenProduct
}

func BankProductManager() *BankProducts {
	bankProducts := &BankProducts{
		Products:          []Product{},
		tempChosenProduct: []ChosenProduct{},
	}
	bankProducts.inisializationProduct()
	return bankProducts
}

func (b *BankProducts) inisializationProduct() {
	b.Products = append(b.Products, Product{ID: "P001", NAME: "Lemari", PRICE: 1500000.00, ADDRESS: "Bogor"})
	b.Products = append(b.Products, Product{ID: "P002", NAME: "Meja", PRICE: 800000.00, ADDRESS: "Jakarta"})
	b.Products = append(b.Products, Product{ID: "P003", NAME: "Kursi", PRICE: 500000.00, ADDRESS: "Bandung"})
	b.Products = append(b.Products, Product{ID: "P004", NAME: "Sofa", PRICE: 2500000.00, ADDRESS: "Surabaya"})
	b.Products = append(b.Products, Product{ID: "P005", NAME: "Kasur", PRICE: 3000000.00, ADDRESS: "Depok"})
	b.Products = append(b.Products, Product{ID: "P006", NAME: "Lemari Pakaian", PRICE: 1750000.00, ADDRESS: "Bogor"})
	b.Products = append(b.Products, Product{ID: "P007", NAME: "Meja Tamu", PRICE: 1200000.00, ADDRESS: "Jakarta"})
	b.Products = append(b.Products, Product{ID: "P008", NAME: "Kursi Kantor", PRICE: 950000.00, ADDRESS: "Bandung"})
	b.Products = append(b.Products, Product{ID: "P009", NAME: "Rak Buku", PRICE: 700000.00, ADDRESS: "Surabaya"})
	b.Products = append(b.Products, Product{ID: "P010", NAME: "Tempat Tidur", PRICE: 4000000.00, ADDRESS: "Depok"})
}

func (b *BankProducts) DisplayProduct(ctx context.Context, key string) {
	if len(b.Products) == 0 {
		b.inisializationProduct()
	}

	var tempchoise []string

	for {
		fmt.Println(utils.ColorMessage("yellow", "\n============== 🛍️   Daftar Produk  🛍️ ============== "))
		fmt.Println(strings.Repeat("-", 50))
		for i, product := range b.Products {
			msg := fmt.Sprintf("%d. | %s | %s | Rp%.2f | %s", i+1, product.ID, product.NAME, product.PRICE, product.ADDRESS)
			fmt.Println(utils.ColorMessage("blue", msg))
			if (i+1)%4 == 0 {
				fmt.Println(strings.Repeat("-", 50))
			}
		}
		fmt.Println(strings.Repeat("-", 50))

		var input string
		fmt.Print(utils.ColorMessage("yellow", "Masukkan Nomor Product (0 untuk selesai): "))
		fmt.Scan(&input)
		utils.ClearScreen()

		intInput, err := strconv.Atoi(input)
		if err != nil {
			utils.ErrorMessage("Input harus berupa angka")
			continue
		}
		if intInput == 0 {
			msg := fmt.Sprintln("Terima kasih telah memilih produk.")
			fmt.Println(utils.ColorMessage("blue", msg))
			break
		}

		if intInput < 0 || intInput > len(b.Products) {
			msg := fmt.Sprintf("Input jangan kurang dari 0 atau lebih dari %d\n", len(b.Products))
			utils.ErrorMessage(msg)
			continue
		}

		idCustomer, okID := ctx.Value(key).(string)
		if !okID {
			utils.ErrorMessage("User ID Customer Invalid")
			return
		}

		product := b.Products[intInput-1]
		b.tempChosenProduct = append(b.tempChosenProduct, ChosenProduct{
			IDCUSTOMER: idCustomer,
			ID:         product.ID,
			NAME:       product.NAME,
			PRICE:      product.PRICE,
			ADDRESS:    product.ADDRESS,
		})

		msg := fmt.Sprintf(" %s - %s, Harga: Rp%.2f", product.ID, product.NAME, product.PRICE)
		tempchoise = append(tempchoise, "Anda memilih:"+utils.ColorMessage("green", msg))
		displayTempchosen(tempchoise)
	}
}

func displayTempchosen(tempchoise []string) {
	for _, msg := range tempchoise {
		fmt.Println(msg)
	}
}

func CountEstimate(address string, addressCustomer string) int {
	if address == addressCustomer {
		return 1
	}

	citiesNearby := map[string][]string{
		"Bogor":    {"Jakarta", "Depok"},
		"Jakarta":  {"Bogor", "Depok"},
		"Bandung":  {"Bogor"},
		"Surabaya": {"Depok"},
		"Depok":    {"Bogor", "Jakarta"},
	}

	if nearbyCities, ok := citiesNearby[address]; ok {
		for _, nearbyCity := range nearbyCities {
			if nearbyCity == addressCustomer {
				return 2
			}
		}
	}

	return 3
}

func (b *BankProducts) DisplayChosenProducts(ctx context.Context, uidKey string, addressKey string) (int, int) {
	idCustomer, okID := ctx.Value(uidKey).(string)
	if !okID {
		utils.ErrorMessage("User ID Customer Invalid")
		return 0, 0
	}

	address, okAdddress := ctx.Value(addressKey).(string)
	if !okAdddress {
		utils.ErrorMessage("User ID Customer Invalid")
		return 0, 0
	}

	if len(b.tempChosenProduct) == 0 {
		utils.ErrorMessage("Keranjang Kosong!!")
		return 0, 0
	}

	fmt.Println(utils.ColorMessage("yellow", "\n\n============== 🛒   Produk Didalam Keranjang   🛒 ============== \n"))
	fmt.Println(strings.Repeat("-", 50))

	estimateCount := 0
	var totalHarga float64
	for i, chosenProduct := range b.tempChosenProduct {
		if idCustomer == chosenProduct.IDCUSTOMER {
			fmt.Printf("%d. %s | %s | Rp%.2f | %s \n", i+1, chosenProduct.ID, chosenProduct.NAME, chosenProduct.PRICE, chosenProduct.ADDRESS)
			totalHarga += chosenProduct.PRICE

			estimateCount_ := CountEstimate(chosenProduct.ADDRESS, utils.Capitalize(address))
			if estimateCount < estimateCount_ {
				estimateCount = estimateCount_
			}
		}
	}

	fmt.Println(strings.Repeat("-", 50))
	dotBlue := utils.ColorMessage("blue", "❖")
	fmt.Printf("%s 🏷️   Total Harga : Rp%.2f\n", dotBlue, totalHarga)
	fmt.Printf("%s 🚚  Estimasi Barang Sampai %d hari\n\n", dotBlue, estimateCount)

	return int(totalHarga), estimateCount
}

func (b *BankProducts) CheckoutProduct(ctx context.Context, uidKey string, addressKey string) {
	totalHarga, estimasi := b.DisplayChosenProducts(ctx, uidKey, addressKey)
	if totalHarga == 0 || estimasi == 0 {
		return
	}

	var input string
	for {
		msg := fmt.Sprintf("%d : ", totalHarga)
		fmt.Print("Masukan Nominal Berikut " + utils.ColorMessage("green", msg))
		fmt.Scan(&input)
		utils.ClearScreen()

		int_input, err := strconv.Atoi(input)
		if err != nil {
			utils.ErrorMessage("Masukan input berupa angka!!")
			continue
		} else if int_input != totalHarga {
			utils.ErrorMessage("Masukan nilai yang sama persis!!")
		} else {
			utils.SuccesMessage("Pesanan Berhasil Dipesan dan Akan Segera Dikirim")
			utils.SuccesMessage(fmt.Sprintf(utils.ColorMessage("yellow", fmt.Sprintf("Estimasi Pengiriman %d Hari\n\n", estimasi))))
			b.tempChosenProduct = []ChosenProduct{}
			return
		}
	}
}

func (b *BankProducts) RemoveProductFromCarts(ctx context.Context, uidKey string, addressKey string) {
	var input string
	var done string

	for {
		totalHarga, estimasi := b.DisplayChosenProducts(ctx, uidKey, addressKey)
		if totalHarga == 0 || estimasi == 0 {
			return
		}

		fmt.Print(utils.ColorMessage("yellow", "Masukkan nomor produk yang ingin dihapus: "))
		fmt.Scan(&input)
		utils.ClearScreen()

		int_input, err := strconv.Atoi(input)
		if err != nil {
			utils.ErrorMessage("Masukan input berupa angka!!")
			continue
		}

		if int_input < 0 || int_input > len(b.tempChosenProduct) {
			utils.ErrorMessage(fmt.Sprintf("Input jangan kurang dari 0 atau lebih dari %d ", len(b.tempChosenProduct)))
			continue
		}

		utils.SuccesMessage("Produk berhasil dihapus")
		int_input -= 1
		b.tempChosenProduct = append(b.tempChosenProduct[:int_input], b.tempChosenProduct[int_input+1:]...)

		fmt.Print(utils.ColorMessage("yellow", "Apakah cukup? (y/t): "))
		fmt.Scan(&done)
		utils.ClearScreen()

		done = strings.ToLower(done)
		if len(done) != 1 || (done != "y" && done != "t") {
			utils.ErrorMessage("Input harus 'y' atau 't' dan tidak boleh lebih dari satu karakter!")
			continue
		}

		if done == "y" {
			return
		}
	}
}
