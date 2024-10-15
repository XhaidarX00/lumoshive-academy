// buatkan function untuk studi kasus berikut
// diketahui ada product sepatu adidas, puma, kappa
// harga sepatu adidas 200000
// harga sepatu puma 150000
// harga sepatu kappa 600000

// harga diskon
// jika memberli sepetu adidas dan puma potong 50000
// jika memberli sepetu puma dan kappa potong 150000
// jika memberli sepetu adidas dan kappa potong 75000
// selain itu harga normal

// tentukan total yang harus dibayar

package main

import "fmt"

func harga_barang(merek string) int {
	switch merek {
	case "adidas":
		harga := 200000
		return harga
	case "puma":
		harga := 150000
		return harga
	case "kappa":
		harga := 600000
		return harga
	default:
		return 0
	}
}

func total_harga(merek1 string, merek2 string) int {
	switch {
	case merek1 == "adidas" && merek2 == "puma":
		harga_adidas := harga_barang("adidas")
		harga_puma := harga_barang("puma")
		return (harga_adidas + harga_puma) - 50000
	case merek1 == "adidas" && merek2 == "kappa":
		harga_kappa := harga_barang("kappa")
		harga_adidas := harga_barang("adidas")
		return (harga_kappa + harga_adidas) - 75000

	case merek1 == "puma" && merek2 == "kappa":
		harga_kappa := harga_barang("kappa")
		harga_puma := harga_barang("puma")
		return (harga_kappa + harga_puma) - 150000
	case merek1 == "puma" && merek2 == "adidas":
		harga_kappa := harga_barang("kappa")
		harga_puma := harga_barang("puma")
		return (harga_kappa + harga_puma) - 150000

	case merek1 == "kappa" && merek2 == "adidas":
		harga_kappa := harga_barang("kappa")
		harga_adidas := harga_barang("adidas")
		return (harga_kappa + harga_adidas) - 75000
	case merek1 == "kappa" && merek2 == "puma":
		harga_kappa := harga_barang("kappa")
		harga_adidas := harga_barang("adidas")
		return (harga_kappa + harga_adidas) - 75000
	default:
		total_ := harga_barang(merek1) + harga_barang(merek2)
		return total_
	}
}

func main() {
	merek1 := "adidas"
	merek2 := "kappa"
	total := total_harga(merek1, merek2)

	fmt.Println("Total Harga : Rp.", total)
}
