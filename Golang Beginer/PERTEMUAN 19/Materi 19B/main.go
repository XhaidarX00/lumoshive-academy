package main

import (
	"embed"
	"fmt"
	"log"
)

// //go:embed files/*
// var embeddedFiles embed.FS

// path matcher mencocokan file atau direktori dengan format tertentu
// contoh
// //go:embed *.txt artinya mencocokan semua file .txt dalam direktori saat ini
// //go:embed images/* mencocokan semua file dalam direktori images
// //go:embed static/** Mencocokan semua file dalam direktori static dan semua subdirektori

// Embed semua file di direktori static

//go:embed static/*
var staticFiles embed.FS

// Embed semua file JSON dan YAML di direktori configs
//
//go:embed configs/*.json configs/*.yaml
var configFiles embed.FS

func main() {
	// Baca semua file yang di-embed
	// dirEntries, err := embeddedFiles.ReadDir("files")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, entry := range dirEntries {
	// 	if !entry.IsDir() {
	// 		data, err := embeddedFiles.ReadFile("files/" + entry.Name())
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Printf("Konten %s: %s\n", entry.Name(), data)
	// 	}
	// }

	// Menampilkan semua file di direktori static
	fmt.Println("File di direktori static:")
	staticDirEntries, err := staticFiles.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range staticDirEntries {
		if !entry.IsDir() {
			data, err := staticFiles.ReadFile("static/" + entry.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Konten %s: %d bytes\n", entry.Name(), len(data))
		}
	}

	// Menampilkan semua file JSON dan YAML di direktori configs
	fmt.Println("\nFile JSON dan YAML di direktori configs:")
	configDirEntries, err := configFiles.ReadDir("configs")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range configDirEntries {
		if !entry.IsDir() {
			data, err := configFiles.ReadFile("configs/" + entry.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Konten %s: %s\n", entry.Name(), data)
		}
	}

}
