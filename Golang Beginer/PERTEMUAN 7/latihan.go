package main

import (
	"fmt"
	"main/utils"
	"sync"
)

func main_latihan() {
	utils.ClearScreen()
	fmt.Println("=------------------=")
	fmt.Println("Latihan Pertemuan 7")
	fmt.Println("=------------------=\n ")

	// buatkan func untuk countdown dijalankan didalam go routine
	// hasil kalkulasi pengulangannya ditampilkan di main function

	// ch := make(chan int)
	// go func() {
	// 	for i := 100; i > 0; i-- {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// for {
	// 	select {
	// 	case data := <-ch:
	// 		fmt.Println("Hasl Kalkulasi ", data)
	// 	}
	// }

	// for value := range ch {
	// 	fmt.Println("Data diterima : ", value)
	// }

	// buatkan 3 function ada parameter channel dan msg
	// dijalankan didalam goroutin yang berbeda
	// buatkan 3 channel untuk setiap func
	// tampilkan pesan dari 3 func yang paling cepat response

	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// ch3 := make(chan string)

	// go SendMsg1(ch1, "Channel 1")
	// go SendMsg2(ch2, "Channel 2")
	// go SendMsg3(ch3, "Channel 3")

	// select {
	// case data := <-ch1:
	// 	fmt.Println("Data diterima dari ", data)
	// case data := <-ch2:
	// 	fmt.Println("Data diterima dari ", data)
	// case data := <-ch3:
	// 	fmt.Println("Data diterima dari ", data)
	// }

	// buatkan satu func dimana func memfilter nilai ganjil untuk ditampilkan
	// kemudian jalankan func tersebut sebanyak 10 kali pada go routine berbeda
	// func memiliki parameter sync.Waitgroup dan range angka yang akan difilter
	// jalankan menggunakan 4 core cpu

	// set cpu menggunakan 4 core
	// runtime.GOMAXPROCS(2)

	// input := 5
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	// go worker(input, &wg)
	// 	go worker2(input, &wg, ch)
	// }

	// wg.Wait()
	// fmt.Println("Semua pekerjaan selesai")

	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for value := range ch {
	// 	utils.SuccesMessage(fmt.Sprintf("%d adalah angka ganjil", value))
	// }

	// buatkan 1 object struct properti qty, name,
	// buatkan func untuk mengurangi qyt
	// func memiliki parameter jumlah qyt yang akan dikurangi
	// jalankan func didalam go routine  sebanyak 10 kali
	// cetak qyt terakhirnya dari struct tersebut

	product := product{name: "Mobil", qty: 100}
	amount := 1

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go product.lessProductQyt(amount, &wg)
	}

	wg.Wait()
	utils.SuccesMessage(fmt.Sprintf("Qty terakhir untuk %s : %d\n", product.name, product.qty))

}

type product struct {
	name string
	qty  int
	mu   sync.Mutex
}

func (p *product) lessProductQyt(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	p.qty -= amount
	fmt.Printf("Mengurangi Qyt sejumah %d \n", amount)
	p.mu.Unlock()
}

// waitgroup
// func worker(count int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < count; i++ {
// 		if i%2 != 0 {
// 			utils.SuccesMessage(fmt.Sprintf("%d adalah angka ganjil", i))
// 		}
// 	}
// }

// func worker2(count int, wg *sync.WaitGroup, ch chan int) {
// 	defer wg.Done()
// 	for i := 0; i < count; i++ {
// 		if i%2 != 0 {
// 			utils.SuccesMessage(fmt.Sprintf("%d adalah angka ganjil", i))
// 		}
// 	}
// }

// func SendMsg1(data chan string, msg string) {
// 	time.Sleep(10 * time.Millisecond)
// 	data <- msg
// }

// func SendMsg2(data chan string, msg string) {
// 	time.Sleep(20 * time.Millisecond)
// 	data <- msg
// }

// func SendMsg3(data chan string, msg string) {
// 	time.Sleep(5 * time.Millisecond)
// 	data <- msg
// }
