package main

import (
	"fmt"
)

func main_() {

	// perbedaan array dan slice
	// slice := []int{1,2,3}
	// array := [3]int{1,2,3}

	// ada deklarasi didalam kurung untuk array

	// if expression -> untuk mengeksekusi blok kode tertentu jika kondisi yang diberikan bernilai benar
	// age := 18
	// if age >= 19 {
	// 	fmt.Println("Kamu Sudah cukup usia")
	// } else {
	// 	fmt.Println("Kamu belum cukup usia")
	// }

	// switch expression
	// day := "monday"
	// switch day {
	// case "saturday", "sunday":
	// 	fmt.Println("it's a weekend")
	// case "monday", "thuesday", "wednesday", "thursday", "friday":
	// 	fmt.Println("it's a weekday")
	// default:
	// 	fmt.Println("unknown day")
	// }

	// switch whitout expression
	// age := 20

	// switch {
	// case age >= 18:
	// 	fmt.Println("Kamu sudah cukup usia")
	// case age <= 18:
	// 	fmt.Println("Kamu belum cukup usia")
	// default:
	// 	fmt.Println("Usia tidak diketahui")
	// }

	// For Expresion -> mengulang blok kode selama kondisi masih terpenuhi
	// for tunggal
	// count := 0
	// for count < 5 {
	// 	fmt.Println("Count : ", count)
	// 	count++
	// }

	// for loop klasik
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Count : ", i)
	// }

	// for tanpa clausa atau tanpa henti kecuali ada break
	// count := 0
	// for {
	// 	fmt.Println("Count : ", count)
	// 	count++
	// 	if count == 5 {
	// 		break
	// 	}
	// }

	//  for dengan range
	// nums := []int{1, 2, 3, 4, 5}
	// for index, value := range nums {
	// 	fmt.Printf("Index ke - %d memiliki Value : %d\n", index, value)
	// }

	// break and continue
	// for i := 0; i < 100; i++ {
	// 	if i%10 == 0 {
	// 		continue
	// 	}

	// 	if i%2 == 0 {
	// 		// jika genap maka print genap
	// 		fmt.Println("GENAP", i)
	// 	} else {
	// 		// jika genap maka print ganjil
	// 		fmt.Println("GANJIL", i)
	// 	}

	// 	if i == 15 {
	// 		break
	// 	}

	// }

	// func -> blok kode yang dapat dipanggil dan melakukan tugas tertentu
	// memiliki 3 dasar nama fungsi, parameters dan return type

	// sayHello()
	// sayHelloName("Iskandar")
	// fmt.Println("Hasil Add :", add(1, 7))
	q, r := divide(10, 5)
	fmt.Println("Quotient :", q)
	fmt.Println("Reminder :", r)

	// func variadik -> fungsi yang menerima jumlah argumen yang berubah-ubah
	// fmt.Println("Sum : ", sum_variadik(1, 2, 3, 4, 5))

	// func anonymus -> fungsi tanpa nama
	// sum := func(a int, b int) int {
	// 	return a + b
	// }

	// fmt.Println("Sum : ", sum(1, 6))

	// func closure -> fungsi tanpa nama yang menangkap variabel lingkup disekitarnya
	// counter := 0
	// increment := func() int {
	// 	counter++
	// 	return counter
	// }

	// fmt.Println(increment())
	// fmt.Println(increment())

	// func sebagai parameter -> memasukan func lain kedalam parameter
	// result := applyOperation(6, 1, add)
	// fmt.Println("Result : ", result)

}

// func sum_variadik(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func applyOperation(a int, b int, operation func(int, int) int) int {
// 	return operation(a, b)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// func tanpa parameter
// func sayHello() {
// 	fmt.Println("Hello world")
// }

// func tanpa parameter
// func sayHelloName(name string) {
// 	fmt.Println("Hello", name)
// }

// func dengan return
// func add(a int, b int) int {
// 	return a + b
// }

// func dengan banyak return
func divide(a int, b int) (int, int) {
	quotient := a / b
	reminder := a % b

	return quotient, reminder
}
