package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

//go:embed hello.txt
var helloFile string

//go:embed lumoshive.png
var photo []byte

func main() {

	// cetak text dalam file
	fmt.Println(helloFile)

	// Tampilkan panjang data foto
	fmt.Printf("Panjang data foto: %d bytes\n", len(photo))

	// Misalnya, Anda ingin menulis foto ke file baru
	err := os.WriteFile("output_photo_lumoshive.jpg", photo, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Foto telah ditulis ke output_photo_lumoshive.jpg")
}
