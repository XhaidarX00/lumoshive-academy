// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name string
// 	Age  int
// }

// // Definisikan struktur data untuk Produk
// type Product struct {
// 	Nama  string `json:"nama"`
// 	Merk  string `json:"merk"`
// 	Harga int    `json:"harga"`
// }

// // Definisikan struktur data untuk Objek JSON utama
// type Catalog struct {
// 	ID        int       `json:"id"`
// 	Nama      string    `json:"nama"`
// 	Deskripsi string    `json:"deskripsi"`
// 	Produk    []Product `json:"produk"`
// }

// func main() {
// 	// encoding basic
// 	p := Person{Name: "Febry", Age: 30}
// 	jsonData1, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData1))

// 	// Data untuk Objek JSON utama (Catalog)
// 	catalog := Catalog{
// 		ID:        1,
// 		Nama:      "Produk Elektronik",
// 		Deskripsi: "Daftar produk elektronik terbaru",
// 		Produk: []Product{
// 			{Nama: "Smartphone", Merk: "Samsung", Harga: 3000000},
// 			{Nama: "Laptop", Merk: "Apple", Harga: 15000000},
// 			{Nama: "Smartwatch", Merk: "Fitbit", Harga: 2000000},
// 		},
// 	}

// 	// Encoding (marshalling) objek JSON utama ke dalam format JSON
// string to json
// 	jsonData2, err := json.MarshalIndent(catalog, "", " ")
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}

// 	// Tampilkan JSON yang dihasilkan
// 	fmt.Println(string(jsonData2))

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // definisikan struktur data untuk person
// type Person struct {
// 	Name string
// 	Age  int
// }

// // Definisikan struktur data untuk Produk
// type Product struct {
// 	Nama  string
// 	Merk  string
// 	Harga int
// }

// // Definisikan struktur data untuk Objek JSON utama
// type Catalog struct {
// 	ID        int
// 	Nama      string
// 	Deskripsi string
// 	Produk    []Product
// }

// func main() {
// 	jsonData1 := []byte(`{"name":"Jerry","age":25}`) // konversi ke dalam byte
// 	var p Person
// 	err1 := json.Unmarshal(jsonData1, &p)
// 	if err1 != nil {
// 		fmt.Println("Error unmarshalling JSON:", err1)
// 		return
// 	}
// 	fmt.Println("Name:", p.Name)
// 	fmt.Println("Age:", p.Age)

// 	// Data JSON yang akan di-decode
// 	jsonData2 := []byte(`{
// 			"id": 1,
// 			"nama": "Produk Elektronik",
// 			"deskripsi": "Daftar produk elektronik terbaru",
// 			"produk": [
//             {"nama": "Smartphone", "merk": "Samsung", "harga": 3000000},
//             {"nama": "Laptop", "merk": "Apple", "harga": 15000000},
//             {"nama": "Smartwatch", "merk": "Fitbit", "harga": 2000000}
//         ]
//     }`)
// 	// Variabel untuk menampung hasil decoding
// 	var catalog Catalog

// 	// Decoding (unmarshalling) JSON ke dalam objek JSON utama (Catalog)
// 	// json to string
// 	err2 := json.Unmarshal(jsonData2, &catalog)
// 	if err2 != nil {
// 		fmt.Println("Error unmarshalling JSON:", err2)
// 		return
// 	}

// 	// Tampilkan data yang telah di-decode
// 	fmt.Println("ID:", catalog.ID)
// 	fmt.Println("Nama:", catalog.Nama)
// 	fmt.Println("Deskripsi:", catalog.Deskripsi)
// 	fmt.Println("Produk:")

// 	for _, product := range catalog.Produk {
// 		fmt.Printf("  - Nama: %s, Merk: %s, Harga: %d\n", product.Nama, product.Merk, product.Harga)
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// // Definisikan struktur data untuk sebuah buku
// type Book struct {
// 	Title  string `json:"title"`
// 	Author string `json:"author,omitempty"` // omitempty akan menghilangkan field jika kosong
// 	ISBN   string `json:"isbn,omitempty"`   // omitempty akan menghilangkan field jika kosong
// 	Pages  int    `json:"-"`                // '-' akan mengabaikan field ini dalam proses encoding dan decoding
// }

// func main() {
// 	// call function encodeJson
// 	encodeJson()
// 	// call function decodeJson
// 	decodeJson()
// 	// call function ignoreTagJson
// 	ignoreTagJson()
// }

// // encoding json with json tag
// func encodeJson() {
// 	p := Person{Name: "John", Age: 30}
// 	jsonData, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))

// }

// // decoding json with json tag
// func decodeJson() {
// 	jsonStr := `{"name":"Jane Smith","age":25}`
// 	var p Person
// 	err := json.Unmarshal([]byte(jsonStr), &p)
// 	if err != nil {
// 		fmt.Println("Error unmarshalling JSON:", err)
// 		return
// 	}
// 	fmt.Println("Name:", p.Name)
// 	fmt.Println("Age:", p.Age)
// }

// // implementation ingore tag json
// func ignoreTagJson() {
// 	// Contoh data buku dengan beberapa field kosong
// 	book := Book{
// 		Title: "Golang Programming",
// 		// Author dan ISBN kosong
// 		Pages: 256,
// 	}

// 	// Encoding (marshalling) struct Book ke JSON
// 	jsonData, err := json.Marshal(book)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}
// 	fmt.Println("Encoded JSON:", string(jsonData))
// 	// Decoding (unmarshalling) JSON ke struct Book
// 	jsonStr := `{"title":"Golang Programming","isbn":"978-3-16-148410-0"}`
// 	var decodedBook Book
// 	err = json.Unmarshal([]byte(jsonStr), &decodedBook)
// 	if err != nil {
// 		fmt.Println("Error unmarshalling JSON:", err)
// 		return
// 	}
// 	fmt.Printf("Decoded Book:\nTitle: %s\nAuthor: %s\nISBN: %s\nPages: %d\n",
// 		decodedBook.Title, decodedBook.Author, decodedBook.ISBN, decodedBook.Pages)

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main_materi() {

// 	// call function encodingJsonMap
// 	encodingJsonMap()
// 	// call function decodingJsonMap
// 	decodingJsonMap()
// }

// func encodingJsonMap() {
// 	// Definisikan map dengan data JSON
// 	data := map[string]interface{}{
// 		"name": "John Doe",
// 		"age":  30,
// 		"address": map[string]interface{}{
// 			"street": "123 Main St",
// 			"city":   "Anytown",
// 		},
// 		"hobbies": []string{"reading", "traveling", "swimming"},
// 	}

// 	// Encoding (marshalling) map ke JSON
// 	jsonData, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}

// 	// Tampilkan JSON yang dihasilkan
// 	fmt.Println(string(jsonData))
// }

// func decodingJsonMap() {
// 	// JSON data yang akan di-decode
// 	jsonStr := `{
// 			"name": "Gigi",
// 			"age": 25,
// 			"address": {
// 				"street": "Kedoya",
// 				"city": "Jakbar"
// 			},
// 			"hobbies": ["Renang", "Nonton"]
// 		}`
// 	// Variabel untuk menampung hasil decoding
// 	var data map[string]interface{}

// 	// Decoding (unmarshalling) JSON ke map
// 	err := json.Unmarshal([]byte(jsonStr), &data)
// 	if err != nil {
// 		fmt.Println("Error unmarshalling JSON:", err)
// 		return
// 	}

// 	// Tampilkan data yang telah di-decode
// 	fmt.Println("Name:", data["name"])
// 	fmt.Println("Age:", data["age"])
// 	fmt.Println("Address:", data["address"].(map[string]interface{}))
// 	fmt.Println("Hobbies:", data["hobbies"].([]interface{}))
// }

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Productt struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func mainn() {

	// call function streamEncodingJson
	// streamEncodingJson()
	// call function streamDecodeJson
	streamDecodeJson()
}

func streamEncodingJsonn() {
	// Buka file untuk menulis
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Buat encoder baru
	encoder := json.NewEncoder(file)
	// Data yang akan diencode
	products := []Product{
		{Name: "Laptop", Price: 1500},
		{Name: "Smartphone", Price: 800},
		{Name: "Tablet", Price: 400},
	}

	// Encode dan tulis setiap produk ke file
	for _, product := range products {
		if err := encoder.Encode(&product); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}

}

func streamDecodeJson() {
	// Buka file JSON
	file, err := os.Open("./product.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Buat decoder baru
	decoder := json.NewDecoder(file)
	// Decode array JSON
	var products []Product
	if err := decoder.Decode(&products); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Tampilkan produk yang telah di-decode
	for _, product := range products {
		fmt.Printf("Product: %+v\n", product)
	}
}

// Normalisasi dan denormalisasi
/*
- kebutuhan data yang diperlukan saat ditampilkan dalam aplikasi frontend
- performa biasaya untuk normalisasi dilakukan dipertama, karena datanya bisa terpisah-pisah
- saat aplikasi dibangun biasanya dinormalisasi
- jika data sudah besar dan banyak fitur dan lainnya, sehingga dapat melakukan query yang sangat komplek
- itu akan mempengaruhi performa jadi penting adanya denormalisasi seperti product disatukan dengan tabel order

*/
