package main

import (
	// "errors"
	"fmt"
	// "Project/model"
	// "reflect"
	// "Project/testing"
)

func main_() {
	fmt.Println("Main...\n ")

	// Reflect -> untuk memeriksa type dan nilai dari variabel saat runtime
	// ini berguna untuk operasi yang membutuhkan intropeksi tipe seperti serialisasi dan deserialisasi data
	// var x float32 = 3.7
	// fmt.Println("Tipe : ", reflect.TypeOf(x))
	// fmt.Println("Tipe : ", reflect.ValueOf(x))

	// v := reflect.ValueOf(x)
	// fmt.Println("Apakah tipe float 32 ? ", v.Kind() == reflect.Float32)

	// defer -> menunda eksekusi fungsi sampai fungsi yang mangandung
	// defer selesai dieksekusi
	// fmt.Println("Mulai")
	// defer fmt.Println("Selesai") // perintah yang dijalankan diakhir program
	// fmt.Println("Tengah")

	// Panic -> memberhentikan eksekusi program saat ditemukan kesalahan yang tidak dapat ditangani
	// fmt.Println("Mulai")
	// panic("Terjadi Kesalahan") // perintah yang dijalankan diakhir program
	// fmt.Println("Selesai")

	// Error Handling -> di go dilakukan dengan mengembalikan nilai kesalahan dengan return string
	// result, err := divide(4, 0)
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// 	return
	// } else {
	// 	fmt.Println("Hasil : ", result)
	// }

	// Recover -> untuk menangkap nilai yang dikirimkan ke panic dan menghentikan unwinding stack,
	// hanya efektif ketika dipanggil dalam fungsi yang tertunda oleh defer
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Pulih dari panic: ", r)
	// 	}
	// }()
	// fmt.Println("mulai")
	// panic("Terjadi Kesalahan")
	// fmt.Println("selesai") // Baris ini tidak akan dieksekusi karena panic.

	// layout format
	// %T untuk type
	// %% untuk karakter %
	// %t untuk Boolean
	// %f desimal tanpa exsponen
	// %g berpindah ke %e dan %f yang lebih ringkas
	// %q string yang di quote
	// %p pointer dalam hexadesimal
	// bisa menggunakan padding untuk kanan %4d dan kiri %-4d

	// Properti public dimulai dengan huruf Besar
	// properti private dimulai dengan huruf kecil untuk penamaan

	// import package harus menggunakan penamaan properti publik
	// testing.PrintNama("Haidar")

	// package inisialisasi dengan menjalankan langsung init ketika package lain diimport atau diakses
	// result := testing.PrintNama("Haidar")
	// fmt.Println(result)

	// 	result := PrintConsole()
	// 	fmt.Println(result)
	// }

	// func divide(a, b float64) (float64, error) {
	// 	if b == 0 {
	// 		return 0, errors.New("Pembagian dengan nol")
	// 	}

	// 	return a / b, nil
}
