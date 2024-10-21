package main

import (
	"fmt"
)

func sendData(ch chan int, value int) {
	fmt.Println("Mengirim data ke channel : ", value)
	ch <- value
}

// func receiveData(ch <-chan int, done chan<- bool) {
// 	value := <-ch
// 	fmt.Println("Data diterima di function : ", value)
// 	done <- true
// }

func printWord(word string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(word)
	}
}

func main_materi() {
	fmt.Println("Materi Golang Pertemuan 6")
	// memungkinkan goroutines lain untuk berkomunikasi dengan fungsi
	// ch := make(chan int)

	// // // menjalankan gorutine untuk mengirim data
	// go sendData(ch, 42)

	// result := <-ch
	// fmt.Println("Data diterima : ", result)

	// go receiveData(ch)

	// ch <- 55
	// fmt.Println("Data telah dikirim")
	// Direction Channel
	// arah untuk menunjukan apakah mereka digunakan untuk mengirim atau menerima data

	// dataCh := make(chan int)
	// doneCh := make(chan bool)

	// go sendData(dataCh, 42)

	// go receiveData(dataCh, doneCh)

	// if <-doneCh {
	// 	fmt.Println("semua pekerjaan selesai....")
	// }

	// menambahkan dependency
	// inisialisasi modul
	// go mod init myproject

	// menambahkan dependency menggunakan perintah go get
	// go get github.com/google/uuid
	// upgrade dependency
	// go get github.com/google/uuid@v1.2.0

	// newUUID := uuid.New()
	// fmt.Println("New UUID : ", newUUID.String())

	// Concurrency
	// konsep menjalankan beberapa tugas secara bersamaan dalam satu program.
	// namun, ini tidak berarti bahwa semua tugas benar-benar berjalan pada saat yang sama

	// parallel programing
	// adalah eksekusi dari beberapa proses atau thread secara benar-benar simultan,
	// biasanya pada beberapa prosesor atau core,
	// meningkatkan kinerja aplikasi dengan membagi tugas besar menjadi beberapa tugas kecil yang dapat dijalankan secara paralel

	// concurency vs paralel programing
	// concurency cocok digunakan ketika aplikasi perlu menangani banyak tugas,
	// paralel cocok digunakan jika memerlukan banyak pemerosesan data

	// goroutine adalah unit dasar concurency di go,
	// mereka adalah fungsi atau methode yang berjalan bersamaan dengan fungsi atau methode lainnya
	// hanya butuh 2kb memory
	// cara kerja pemanggilan menggunakan kata depan go
	// bisa dijalankan di satu thread atau banyak thread
	// go printWord("Hello", 100)

	// printWord("World", 100)

	// goroutin anosymouse
	// go func() {
	// 	fmt.Println("Hello from anonymuse goruntine")
	// }()
	// time.Sleep(1 * time.Second) // menunggu secenak agar goroutine bisa selesai

	// Sistem basis data
	// dbms -> basis data -> model data
	// basis data kumpulan data yang terstruktur sehingga mudah dikelola
	// biasanya terbentuk tabel yang terdiri dari baris dan kolom
	// Basis data relasional -> menggunakan tabel untuk menyimpan data dan hubungan antar data
	// dbms -> untuk membuat dan mengelola basis data

	// model data adalah cara untuk mendefinisikan dan meyusun data dalam sistem basis data
	// model data menentukan bagaimana data disimpan, diakses, dan dihubungkan satu sama lain
	// model relasional : data diorganisir dalam tabel yang terdiri dari baris dan kolom setiap tabel memiliki
	// kunci utama atau primary key unik, dan hubungan antar tabel diatur mellaui primary key

	// model hierarkis: data diorganisir dalam struktur pohon, dim mana setiap entitas memiliki
	// satu induk dan dapat memiliki banyak anak

	// model nonsql dihunakan dalam data yang tidak terstruktur atau semi terstruktur
	// ada beberapa jenis model nonSql, termasuk basis data dokumen, basis data graf dan basis data key-value

	// mode data ada bebrapa komponen utama
	// Entitas objek nyata atau konsep informasi perlu disimpan dalam basis data.
	// contoh atribut untuk entitas "pelanggan", "produk", dan "pesanan".
	// Atribut : karakteristik atau propertu dari entitas yang perlu disimpan.
	// contoh untuk entitas pelanggan : "nama", "alamat", dan "nomor telphone"
	// Hubungan (Relationships): hubungan menentukan bagaimana tabel-tabel dalam basis data saling berhubungan
}
