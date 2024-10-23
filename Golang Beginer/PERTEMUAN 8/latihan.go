package main

import (
	"context"
	"fmt"
	"main/utils"
	"time"
)

// buatkan 2 func yang memiliki parameter context
// func 1 berfungsi untuk menampilkan nilai dari context yang dikirim
// kemudian func 2 berfungsi untuk mengolah nilai dari context

// func printValue(ctx context.Context, key string) {
// 	value := ctx.Value(key)
// 	if value != nil {
// 		utils.SuccesMessage(fmt.Sprintf("Hasil value dari context : %d", value))
// 	} else {
// 		utils.ErrorMessage("Key tidak ditemukan dalam context.")
// 	}
// }

// func addValue(ctx context.Context, key string, amount int) int {
// 	value := ctx.Value(key)
// 	if value != nil {
// 		if v, ok := value.(int); ok {
// 			v += amount
// 			return v
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	key := "angka"
// 	ctx := context.Background()
// 	ctxWithValue := context.WithValue(ctx, key, 10)

// 	printValue(ctxWithValue, key)

// 	amount := 20
// 	newValue := addValue(ctxWithValue, key, amount)

// 	if newValue != 0 {
// 		utils.SuccesMessage(fmt.Sprintf("Nilai baru setelah penambahan: %d", newValue))
// 	} else {
// 		utils.ErrorMessage("Tidak dapat menambahkan nilai.")
// 	}
// }

// buatkan 3 func
// func 1 mencetak text setiap 2 detik
// func 2 mencetak text setiap 1 detik
// func 3 mencetak text setiap 3 detik
// buatkan context untuk membatalkan semua func yang berjalan
// di detik ke 5

// func printData(ctx context.Context, name string) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Fungsi %s dibatalkan\n", name)
// 			return
// 		default:
// 			switch name {
// 			case "hewan":
// 				fmt.Printf("%s sedang mencari makan\n", name)
// 				time.Sleep(2 * time.Second)
// 			case "haidar":
// 				fmt.Printf("%s sedang coding golang\n", name)
// 				time.Sleep(1 * time.Second)
// 			default:
// 				fmt.Printf("%s ngga tau lagi ngapain\n", name)
// 				time.Sleep(3 * time.Second)
// 			}
// 		}
// 	}
// }

// func main() {
// 	ctx := context.Background()
// 	ctxwithcancel, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	go printData(ctxwithcancel, "hewan")
// 	go printData(ctxwithcancel, "haidar")
// 	go printData(ctxwithcancel, "buah")

// 	time.Sleep(5 * time.Second)
// 	cancel()
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Program Dicancel setelah 5 Detik")
// }

// buat func tambah key value ke context
// buat func tampil value setiap 2 detik
// buat func ubah nilai value

// func initialization(ctx *context.Context, key string, value int) {
// 	*ctx = context.WithValue(*ctx, key, value)
// }

// func display(ctx context.Context, key string) {
// 	fmt.Println("Value Context :", ctx.Value(key))
// }

// func changeValue(ctx *context.Context, key string) {
// 	newValue := (*ctx).Value(key).(int) + 1
// 	*ctx = context.WithValue(*ctx, key, newValue)
// }

// func main() {
// 	key := "user"
// 	value := 12345
// 	ctx := context.Background()

// 	initialization(&ctx, key, value)
// 	for {
// 		time.Sleep(2 * time.Second)
// 		display(ctx, key)

// 		changeValue(&ctx, key)
// 	}

// }

// buatkan satu variabel kemudia inisialisasikan dengan nilai tertentu
// buatkan satu func untuk mengurangi nilai variabel itu
// jalankan func menggunakan goroutin
// buatkan kondisi didetik ke 4 operasi pengulangannya dihentikan

// func worker(ctxc context.Context, ctxv *context.Context, key string, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	select {
// 	case <-ctxc.Done():
// 		fmt.Println("Progam Dihentikan")
// 		return
// 	default:
// 		mutex.Lock()
// 		newValue := (*ctxv).Value(key).(int) + 1
// 		*ctxv = context.WithValue(*ctxv, key, newValue)
// 		mutex.Unlock()
// 		wg.Done()
// 	}
// }

// func display(ctx context.Context, key string) {
// 	fmt.Println("Value Context :", ctx.Value(key))
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex
// 	ctx := context.Background()
// 	ctxwithcancel, cancel := context.WithCancel(ctx)
// 	ctxwithvalue := context.WithValue(ctx, "user", 1)
// 	defer cancel()

// 	for {
// 		time.Sleep(2000 * time.Millisecond)
// 		wg.Add(1)
// 		go worker(ctxwithcancel, &ctxwithvalue, "user", &wg, &mutex)
// 		display(ctxwithvalue, "user")
// 	}
// }

// buatkan 1 struct dengan 2 properti tipe data bebas
// buat func untuk memasukan struct ke slice
// buatkan context menggunakan deadline untuk 10 detik ke depan
// agar func tidak lagi bisa memasukan data struct kedalam slice

// type Person struct {
// 	ID   int
// 	NAME string
// }

// func (p Person) addPerson(ctx context.Context, data *[]Person, id int, name string) {

// 	select {
// 	case <-ctx.Done():
// 		utils.SuccesMessage("Program mencapai deadline")
// 		return
// 	default:
// 		time.Sleep(2 * time.Second)
// 		*data = append(*data, Person{ID: id, NAME: name})
// 		fmt.Printf("Data %v\n", *data)
// 	}
// }

// func main() {
// 	ctx := context.Background()
// 	deadline := time.Now().Add(10 * time.Second)
// 	ctxwithdeadline, cancel := context.WithDeadline(ctx, deadline)
// 	defer cancel()

// 	var data []Person

// 	person := Person{ID: 1, NAME: "Haidar"}

// 	for i := 0; i < 10; i++ {
// 		person.addPerson(ctxwithdeadline, &data, i+1, fmt.Sprintf("User", i))
// 	}

// }

// buat slice dengan 2 data struct
// struct role dan nama
// role memiliki akses menu yang berbeda
// buat 1 func berufungsi untuk cek 2 data struct dalam slice
// kondisi :
// - if ada masukan ke context value
//
// buat func untuk menampilkan beberapa menu sesuai dengan data yang dipilih
// pengecekan data tida bisa diakses dalam waktu tertertentu with deadline

type Data struct {
	role string
	nama string
}

func display(ctx context.Context, key string) {
	fmt.Println("Value Context :", ctx.Value(key))
}

func (d Data) addData(SliceData []Data, role string, name string, ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Akses ditolak: waktu habis!")
		return
	default:
		for _, data := range SliceData {
			if data.role == role && data.nama == name {
				fmt.Printf("Selamat datang %s (%s)\n", data.nama, data.role)
				ctxWithValue := context.WithValue(ctx, data.role, data.nama)
				display(ctxWithValue, data.role)
				showMenu(data.role)
				return
			}
		}
		fmt.Println("Data tidak valid!")
	}
}

func showMenu(role string) {
	switch role {
	case "admin":
		fmt.Println("==== Menu Admin ====")
		adminMenu := []string{"1. Dashboard", "2. User Management", "3. Reports"}

		for _, item := range adminMenu {
			fmt.Println(utils.ColorMessage("green", item))
		}
		fmt.Println("\n ")

	case "client":
		fmt.Println("==== Menu Client ====")
		clientMenu := []string{"1. Profile", "2. Orders", "3. Support"}
		for _, item := range clientMenu {
			fmt.Println(utils.ColorMessage("green", item))
		}
		fmt.Println("\n ")

	default:
		fmt.Println("Menu tidak ditemukan!")
	}
}

func main_latihan() {
	var SliceData []Data
	data := Data{}
	SliceData = append(SliceData, Data{role: "admin", nama: "haidar"})
	SliceData = append(SliceData, Data{role: "client", nama: "haidar2"})

	fmt.Printf("data : %v\n\n", SliceData)
	var role string
	var name string

	ctx := context.Background()
	deadline := time.Now().Add(10 * time.Second)
	ctxWithDeadline, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	for {
		fmt.Print("Masukan Role : ")
		fmt.Scan(&role)
		fmt.Print("Masukan Nama : ")
		fmt.Scan(&name)
		utils.ClearScreen()

		data.addData(SliceData, role, name, ctxWithDeadline)
		var choice string
		fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
		fmt.Scan(&choice)
		if choice != "y" {
			break
		}
	}
}
