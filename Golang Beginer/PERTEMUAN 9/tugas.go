package main

import (
	"encoding/json"
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	StatusOrder string `json:"statusOrder"`
	Delivered   bool   `json:"delivered"`
}

type Orders struct {
	Idorder    string    `json:"idOrder"`
	Products   []Product `json:"products"`
	TotalHarga int       `json:"totalHarga"`
}

type HistoryOrders struct {
	Time   time.Time `json:"time"`
	Orders []Orders  `json:"orders"`
}

var orders []Orders
var products []Product
var chosenProducts []Product
var historyOrders []HistoryOrders

func DoPayment(orderId string) {
	for i, order := range orders {
		if order.Idorder == orderId {
			orders[i].TotalHarga = calculateTotal(order.Products)
			time.Sleep(2 * time.Second)
			orders[i].Products = updateOrderStatus(order.Products, "paid")
			historyOrders = append(historyOrders, HistoryOrders{Time: time.Now(), Orders: orders})
			return
		}
	}
	utils.ErrorMessage("Order tidak ditemukan!")
}

func EditOrder(orderId string) {
	fmt.Println(utils.ColorMessage("blue", "\n\n=-------------- Edit Pesanan --------------="))
	DisplayProduct()
	for i, order := range orders {
		if order.Idorder == orderId {
			orders[i].Products = chosenProducts
			utils.SuccesMessage("Order berhasil diupdate!")
			if saveToFileOrder("Orders.json", orders) {
				return
			}
		}
	}
	utils.ErrorMessage("Order tidak ditemukan!")
}

func ViewOrders() {
	ordersJson, err := json.MarshalIndent(historyOrders, "", "  ")
	if err != nil {
		utils.ErrorMessage(fmt.Sprintf("Error marshalling orders: %v", err))
		return
	}

	if saveToFileHistory("HistoryOrders.json", historyOrders) {
		return
	}

	fmt.Println(string(ordersJson))
}

func updateOrderStatus(products []Product, newStatus string) []Product {
	for i := range products {
		products[i].StatusOrder = newStatus
	}

	utils.SuccesMessage("Berhasil Mengubah status menjadi paid")
	fmt.Println(utils.ColorMessage("yellow", "Pesanan Sedang Diantar"))
	return products
}

func calculateTotal(products []Product) int {
	total := 0
	for _, product := range products {
		total += product.Price
	}

	var input string
	for {
		msg := fmt.Sprintf("%d : ", total)
		fmt.Print("Masukkan Nominal Berikut " + utils.ColorMessage("green", msg))
		fmt.Scan(&input)
		utils.ClearScreen()

		intInput, err := strconv.Atoi(input)
		if err != nil {
			utils.ErrorMessage("Masukkan input berupa angka!!")
			continue
		} else if intInput != total {
			utils.ErrorMessage("Masukkan nilai yang sama persis!!")
		} else {
			utils.SuccesMessage("Pesanan Berhasil Dibuat")
			fmt.Println(utils.ColorMessage("yellow", "Pesanan Sedang Diproses dan Akan Segera Diantar"))
			return total
		}
	}
}

func OrdersProduct() string {
	idOrder := uuid.New().String()
	orders = append(orders, Orders{
		Idorder:  idOrder,
		Products: chosenProducts,
	})

	if saveToFileOrder("Orders.json", orders) {
		return idOrder
	}

	return ""
}

func saveToFileOrder(filename string, data []Orders) bool {
	file, err := os.Create(filename)
	if err != nil {
		utils.ErrorMessage(fmt.Sprintf("Error creating file: %v", err))
		return false
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		utils.ErrorMessage(fmt.Sprintf("Error encoding JSON: %v", err))
		return false
	}

	utils.SuccesMessage(fmt.Sprintf("Berhasil menyimpan data ke file %s", filename))
	return true
}

func saveToFileHistory(filename string, newData []HistoryOrders) bool {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		utils.ErrorMessage(fmt.Sprintf("Error opening file: %v", err))
		return false
	}
	defer file.Close()

	var existingData []HistoryOrders

	fileInfo, err := file.Stat()
	if err != nil {
		utils.ErrorMessage(fmt.Sprintf("Error getting file info: %v", err))
		return false
	}
	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&existingData); err != nil {
			utils.ErrorMessage(fmt.Sprintf("Error decoding JSON: %v", err))
			return false
		}
	}

	existingData = append(existingData, newData...)
	file.Truncate(0)
	file.Seek(0, 0)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(existingData); err != nil {
		utils.ErrorMessage(fmt.Sprintf("Error encoding JSON: %v", err))
		return false
	}

	utils.SuccesMessage(fmt.Sprintf("Berhasil menyimpan data ke file %s", filename))
	return true
}

func inisializationProduct() {
	products = append(products, Product{Name: "Rendang", Price: 25000, StatusOrder: "unpaid", Delivered: false})
	products = append(products, Product{Name: "Ayam Bakar", Price: 15000, StatusOrder: "unpaid", Delivered: false})
	products = append(products, Product{Name: "Es Teh Serut", Price: 10000, StatusOrder: "unpaid", Delivered: false})
	products = append(products, Product{Name: "Sayur Asem", Price: 25000, StatusOrder: "unpaid", Delivered: false})
}

func DisplayProduct() {
	if len(products) == 0 {
		inisializationProduct()
	}

	var input string
	var tempChoice []string

	for {
		fmt.Println(utils.ColorMessage("yellow", "\n============== 🛒   Produk Rumah Makan Golang 🛒 =============="))
		fmt.Println(strings.Repeat("-", 50))
		for i, product := range products {
			fmt.Printf("%d. %s | | Rp%d \n", i+1, product.Name, product.Price)
		}
		fmt.Println(strings.Repeat("-", 50))

		fmt.Print(utils.ColorMessage("yellow", "Masukkan Nomor Produk atau ketik 'done' untuk selesai : "))
		fmt.Scan(&input)
		utils.ClearScreen()

		if strings.ToLower(input) == "done" {
			msg := "Terima kasih telah memilih produk."
			fmt.Println(utils.ColorMessage("blue", msg))
			break
		}

		intInput, err := strconv.Atoi(input)
		if err != nil {
			utils.ErrorMessage("Input harus berupa angka")
			continue
		}

		if intInput < 1 || intInput > len(products) {
			msg := fmt.Sprintf("Input jangan kurang dari 1 atau lebih dari %d\n", len(products))
			utils.ErrorMessage(msg)
			continue
		}

		product := products[intInput-1]
		chosenProducts = append(chosenProducts, product)
		msg := fmt.Sprintf("Anda memilih menu: %s", utils.ColorMessage("green", product.Name))
		tempChoice = append(tempChoice, msg)

		displayTempChoice(tempChoice)
	}
}

func displayTempChoice(tempChoice []string) {
	for _, msg := range tempChoice {
		fmt.Println(msg)
	}
}

func DisplayChosenHistory() {
	time.Sleep(1 * time.Second)
	fmt.Println(utils.ColorMessage("yellow", "\n============== 🛒   History Order 🛒 =============="))
	fmt.Println(strings.Repeat("-", 50))

	total := 0
	for i, product := range chosenProducts {
		fmt.Printf("%d. %s | | Rp%d \n", i+1, product.Name, product.Price)
		total += product.Price
	}
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(utils.ColorMessage("green", fmt.Sprintf("Total Harga : %d", total)))
	fmt.Println(strings.Repeat("-", 50))

	fmt.Println(utils.ColorMessage("yellow", "Pesanan Selesai\n"))
	chosenProducts = []Product{}
}

func main() {
	var done string
	for {
		utils.ClearScreen()
		fmt.Println(utils.ColorMessage("blue", "\n\n=-------------- Memesan Menu --------------="))
		DisplayProduct()
		idOrder := OrdersProduct()
		EditOrder(idOrder)
		DoPayment(idOrder)
		DisplayChosenHistory()
		ViewOrders()
		for {
			fmt.Print(utils.ColorMessage("yellow", "\nApakah anda ingin memesan kembali? (y/t): "))
			fmt.Scan(&done)
			utils.ClearScreen()

			done = strings.ToLower(done)
			if len(done) != 1 || (done != "y" && done != "t") {
				utils.ErrorMessage("Input harus 'y' atau 't' dan tidak boleh lebih dari satu karakter!")
				continue
			}

			if done == "t" {
				return
			}

			break
		}
	}
}
