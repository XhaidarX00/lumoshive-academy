package main

import "fmt"

func main() {
	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	diskon := 15.0
	jumlahTerjual := 100

	harga_diskon := hargaJual - (hargaJual * diskon / 100)

	total_pendapatan := harga_diskon * float64(jumlahTerjual)

	total_biaya := (hargaBeli * float64(jumlahTerjual)) + (biayaOperasional * float64(jumlahTerjual))

	total_keuntungan := total_pendapatan - total_biaya

	fmt.Printf("Harga Jual Setelah Diskon: Rp %.0f\n", harga_diskon)
	fmt.Printf("Total Pendapatan: Rp %.0f\n", total_pendapatan)
	fmt.Printf("Total Biaya: Rp %.0f\n", total_biaya)
	fmt.Printf("Total Keuntungan: Rp %.0f\n", total_keuntungan)
}
