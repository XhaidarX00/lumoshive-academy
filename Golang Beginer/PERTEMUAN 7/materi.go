package main

import (
	"fmt"
)

func main_materi() {
	fmt.Println("Materi Pertemuan 7")
	// membuat buffer channel dengan kapasitas 2
	// ch := make(chan int, 2)

	// go func() {
	// 	fmt.Println("Mengirim data pertama ke channel...")
	// 	ch <- 1 // mengirim data ke channel
	// 	fmt.Println("Data pertama telah dikirim")

	// 	fmt.Println("Mengirim data kedua ke channel...")
	// 	ch <- 2 // mengirim data ke channel
	// 	fmt.Println("Data kedua telah dikirim")
	// }()

	// // menerima data dari channel
	// fmt.Println("Menunggu menerima data pertama...")
	// value1 := <-ch // menerima data dari chanel
	// fmt.Println("Data pertama diterima : ", value1)

	// fmt.Println("Menunggu menerima data kedua...")
	// value2 := <-ch // menerima data dari chanel
	// fmt.Println("Data kedua diterima : ", value2)

	// Range channel
	// Membuat buffered channel dengan kapasitas 3
	// ch := make(chan int, 3)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		ch <- i
	// 		fmt.Println("Mengirim ", i)
	// 	}
	// 	close(ch) // menutup channel setelah selesai mengirim
	// }()

	// // Menggunakan range untuk menerima data dari channel sampai channel ditutup
	// for value := range ch {
	// 	fmt.Println("Menerima : ", value)
	// }

	// fmt.Println("Semua data telah diterima")

	// select channel
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go fetchFromAPI1(ch1)
	// go fetchFromAPI2(ch2)

	// // switch case untuk channel
	// select {
	// case data := <-ch1:
	// 	fmt.Println("Menerima1 : ", data)
	// case data := <-ch2:
	// 	fmt.Println("Menerima2 : ", data)
	// }

	// default channel adalah default case yang jika tidak ada channel yang siap untuk mengirim atau menerima data
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// // goroutime untuk mengirim data ke channel 1 setelah 2 detik
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch1 <- "Data dari channel 1"
	// }()

	// // goroutime untuk mengirim data ke channel 2 setelah1 detik
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch1 <- "Data dari channel 2"
	// }()

	// // menggunakan select untuk menangani kedua channel
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case data := <-ch1:
	// 		fmt.Println("Menerima data dari ", data)
	// 	case data := <-ch2:
	// 		fmt.Println("Menerima data dari ", data)
	// 	default:
	// 		fmt.Println("Tidak ada data yang siap untuk diterima")
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }

	// Gomaxprocs variabl lingkungan yang menentukan jumlah maksimum threat yang dijalankan secara simultan untuk mengeksekusi gorouten
	// jumlah GOMAXPROCS defaultnya adalah jumlah cpu fisik yang ada di mesin tempat program go berjalan
	// mendapatkan jumlah cpu yang tersedia
	// numCpu := runtime.NumCPU()
	// fmt.Printf("Jumlah CPU : %d\n", numCpu)

	// // mengatur GOMAXPROCS menjadi 2
	// runtime.GOMAXPROCS(4)

	// // contoh penggunaan goroutine
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("Selesai menjalankan goroutine")
	// 	}()
	// }

	// // menunggu goroutine selesai
	// time.Sleep(time.Second * 2)
	// fmt.Println("Selesai menjalankan program")

	// waitGroup adalah struktur data di go yang digunakan untuk menunggu sekelompok goroutine selesai.
	// ketika memiliki beberapa goroutine yang ingin menunggu sebelum melanjutkan eksekusi program utama
	// contoh penggunaan WaitGroup
	// var wg sync.WaitGroup
	// for i := 0; i <= 5; i++ {
	// 	wg.Add(1) // menambahkan jumlah goroutine yang sedang berjalan
	// 	go worker(i, &wg)
	// }

	// wg.Wait() // menunggu sampai goroutine selesai
	// fmt.Println("Semua pekerjaan selesai")

	// Race condition adalah sebuah keaddan dimana dua atau lebih goroutine
	// berusaha untuk mengakses dan memodifikasi resource bersamaan tanpa mekanisme sinkronisasi yang tepat
	// akibatnya opreasi tersebut tidak konsisten
	// contoh racecondition
	// var counter int
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		counter++
	// 	}()
	// }

	// fmt.Println("Nilai counter: ", counter)

	// sync.Mutex adalah struktur data di go yang digunakan untuk mengontrol akses ke suatu course dari beberapa goroutine secara bersamaan
	// mutex (mutual exclusion) memungkinkan hanya satu goroutin yang dapat mengakses resource pada satu waktu tertentu
	// sehingga mencegah race condition terjadi
	// contoh sync.Mutex
	// var counter int
	// var mutex sync.Mutex
	// var wg sync.WaitGroup

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		mutex.Lock()
	// 		counter++
	// 		mutex.Unlock()
	// 		wg.Done()
	// 	}()
	// }

	// wg.Wait()
	// fmt.Println("Nilai counter: ", counter)

	// time package time menyediakan fungsionaliteas untuk mengelola waktu dan tanggal
	// fungsi time.now digunakan untuk mendapatkan waktu saaat ini sesuai dengan zona
	// waktu lokas sistem tempat program berjalan
	// now := time.Now()
	// fmt.Println("Current time: ", now)

	// time parse menconversi string yang berisi waktu ke dalam objek time
	// layout := "2006-01-02 15:04:05"
	// str := "2023-06-23 14:40:00"
	// parsedTime, err := time.Parse(layout, str)
	// if err != nil {
	// 	fmt.Println("Error parsing time : ", err)
	// 	return
	// }

	// fmt.Println("Parsed Time : ", parsedTime)

	// time format untuk mengconversi object time kedalam string sesuai dengan layout yang ditentukan
	// contoh
	// now := time.Now()
	// // format tanggal dan waktu lengkap
	// layout := "2005-08-02 15:50:05"
	// formatedTime1 := now.Format(layout)
	// fmt.Println("Fromatted time 1 : ", formatedTime1)

	// // format tanggal
	// layout2 := "2005-08-02"
	// formatedTime2 := now.Format(layout2)
	// fmt.Println("Fromatted time 2 : ", formatedTime2)

	// // format dengan nama, hari dan bulan
	// layout3 := "Monday, 02-Jan-2006 03:04:05 PM"
	// formatedTime3 := now.Format(layout3)
	// fmt.Println("Fromatted time 3 : ", formatedTime3)

	// // format dengan zona waktu
	// layout4 := "2005-08-02 15:04:05 MST"
	// formatedTime4 := now.Format(layout4)
	// fmt.Println("Fromatted time 4 : ", formatedTime4)

	// time.ticker untuk menghasilkan peristiwa secara berkala pada interval waktu tertentu
	// mengirim nilai waktu secara berulang ke channel yang bisa diterima dalam goroutin lain
	// contoh
	// ticker := time.NewTicker(1 * time.Second)
	// done := make(chan bool)
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	done <- true
	// }()

	// go func() {
	// 	for {
	// 		select {
	// 		case t := <-ticker.C:
	// 			fmt.Sprintln("Tick at ", t)
	// 		case <-done:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()

	// time.Sleep(6 * time.Second)
	// fmt.Println("Selesai")

	// time.after
	// membaut channel yang mengirimkan waktu tunggal setelah 2 detik
	// timeout := time.After(2 * time.Second)
	// // menunggu nilai dari channel time out atau done
	// select {
	// case <-timeout:
	// 	fmt.Println("Timeout occured")
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("Operation took to long")
	// }

	// fmt.Println("Selesai")

	// time.AfterFunc digunakan untuk menjadwalkan eksekusi fungsi tertentu

}

// waitgroup
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Menandakan bahwa goroutine selesai saat keluar func
// 	fmt.Printf("Worker %d sedang bekerja\n", id)
// 	// melakukan pekerjaan yang sesuai disini
// }

// func fetchFromAPI1(ch chan string) {
// 	time.Sleep(2 * time.Second) // Simulasi API lambat
// 	ch <- "Data dari API 1"
// }

// func fetchFromAPI2(ch chan string) {
// 	time.Sleep(1 * time.Second) // Simulasi API lambat
// 	ch <- "Data dari API 2"
// }
