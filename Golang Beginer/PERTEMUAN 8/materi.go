// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main() {
// 	// Membuat parent context
// 	parentCtx := context.Background()

// 	// Mendefinisikan tipe unik untuk key
// 	type key string

// 	// Membuat child context pertama dengan nilai
// 	ctx1 := context.WithValue(parentCtx, key("key1"), "hidup")

// 	// Membuat child context kedua dengan nilai
// 	ctx2 := context.WithValue(ctx1, key("key2"), "harus")

// 	// Membuat child context ketiga dengan nilai
// 	ctx3 := context.WithValue(ctx2, key("key3"), "aman")

// 	// Mengakses nilai dari context ketiga
// 	printContextValue(ctx3, key("key1"))
// 	printContextValue(ctx3, key("key2"))
// 	printContextValue(ctx3, key("key3"))
// }

// // Fungsi untuk mencetak nilai dari context
// func printContextValue(ctx context.Context, k interface{}) {
// 	value := ctx.Value(k)
// 	if value != nil {
// 		fmt.Printf("Key: %v, Value: %v\n", k, value)
// 	} else {
// 		fmt.Printf("Key: %v tidak ditemukan dalam context\n", k)
// 	}
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Membuat parent context
// 	parentCtx := context.Background()

// 	// Membuat context dengan cancel
// 	ctx, cancel := context.WithCancel(parentCtx)

// 	// Fungsi defer untuk memastikan cancel dipanggil untuk membersihkan sumber daya
// 	defer cancel()

// 	// Menjalankan beberapa operasi yang menggunakan context
// 	go worker(ctx, "Worker 1")
// 	go worker(ctx, "Worker 2")
// 	go worker(ctx, "Worker 3")

// 	// Menunggu sebentar sebelum membatalkan context
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Membatalkan context...")
// 	cancel() // Membatalkan context

// 	// Menunggu sebentar untuk melihat hasil pembatalan
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Aplikasi selesai")
// }

// func worker(ctx context.Context, name string) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// Sinyal pembatalan diterima
// 			fmt.Printf("%s dibatalkan\n", name)
// 			return
// 		default:
// 			// Melakukan pekerjaan
// 			fmt.Printf("%s sedang bekerja\n", name)
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	// Membuat parent context
// 	parentCtx := context.Background()

// 	// Membuat context dengan batas waktu 3 detik
// 	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
// 	defer cancel() // Pastikan untuk memanggil cancel untuk membersihkan sumber daya

// 	// URL layanan yang akan diperiksa
// 	url := "https://jsonplaceholder.typicode.com/todos/1"

// 	// Membuat request HTTP dengan context
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	// Membuat client HTTP
// 	client := &http.Client{}

// 	// Melakukan request dengan context
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Memeriksa status response
// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Println("Layanan tersedia")
// 	} else {
// 		fmt.Println("Layanan tidak tersedia")
// 	}
// }

package main

import (
	"context"
	"fmt"
	"time"
)

func main_m() {
	// Membuat parent context
	parentCtx := context.Background()

	// Menentukan waktu deadline: 5 detik dari sekarang
	deadline := time.Now().Add(5 * time.Second)

	// Membuat context dengan deadline
	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	// Fungsi defer untuk memastikan cancel dipanggil untuk membersihkan sumber daya
	defer cancel()

	// Menjalankan operasi yang menggunakan context
	result := process(ctx)
	fmt.Println("Hasil operasi:", result)
}

func process(ctx context.Context) string {
	select {
	case <-time.After(3 * time.Second):
		return "Proses selesai tepat waktu"
	case <-ctx.Done():
		// Context dibatalkan (termasuk karena deadline)
		return "Proses dibatalkan karena melewati deadline"
	}
}
