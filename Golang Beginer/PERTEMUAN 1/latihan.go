package main

import (
	"fmt"
	"strconv"
)

func main_test() {
	// buat 4 variabel, int, float32, int, float32
	// int 30, float32 24.5, int -45, float32 0.67
	// hitung matematika buat var baru
	// a + b * c / d
	// hasilnya dikonversi ke string

	a := 30
	b := float32(24.5)
	c := -45
	d := float32(0.67)

	// perhitungan
	hasil := float32(a) + b*float32(c)/d

	// var str4 string = strconv.FormatFloat(f2, 'f', 6, 64)
	hasil_cvt := strconv.FormatFloat(float64(hasil), 'f', 6, 64)

	fmt.Println("Hasil: ", hasil_cvt)
}
