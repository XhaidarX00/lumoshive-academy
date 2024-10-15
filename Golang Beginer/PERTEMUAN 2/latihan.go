package main

import (
	"fmt"
)

func main_latihan() {
	// buatkan func jika beli 1 barang tidak dapat diskon
	// 2 beli 10 persen
	// 4 beli 50 persen
	// lebih dari 4 75 persen

	jumlahBeli := 7
	total_pembelian := 50000

	diskon := []int{10, 50, 75}
	result := harga_per_pembelian(jumlahBeli, total_pembelian, diskon)
	fmt.Println("Harga Total : ", result)

	// hitunglah gaji karyawan jika satu jam dibayar 50000
	// lembur 1 jam 60000
	// total bekerja 40 jam dengan lembur 5 jam
	// buatkan function untuk menghitung total gaji

	// jamKerja := 40
	// jamLembur := 5
	// gaji := hitungGaji(jamKerja, jamLembur)

	// fmt.Println("Total Gaji : Rp.", gaji)

}

func diskon_harga(diskon int, hargaJual int) int {
	harga_diskon := hargaJual - (hargaJual * diskon / 100)
	return harga_diskon
}

func harga_per_pembelian(jumlahBeli int, harga int, diskon []int) int {
	switch {
	case jumlahBeli == 1:
		temp_harga := harga
		return temp_harga
	case jumlahBeli == 2:
		temp_harga := diskon_harga(diskon[0], harga)
		return temp_harga
	case jumlahBeli < 4:
		temp_harga := diskon_harga(diskon[0], harga)
		return temp_harga
	case jumlahBeli == 4:
		temp_harga := diskon_harga(diskon[1], harga)
		return temp_harga
	case jumlahBeli > 4:
		temp_harga := diskon_harga(diskon[2], harga)
		return temp_harga
	default:
		return 0
	}

}

func hitungGaji(jamKerja int, jamLembur int) int {
	harga_lembur_perjam := 60000
	harga_kerja_perjam := 50000

	total_bayaran := (harga_kerja_perjam * jamKerja) + (harga_lembur_perjam * jamLembur)
	return total_bayaran
}
