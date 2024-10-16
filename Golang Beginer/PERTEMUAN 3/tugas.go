package main

import "fmt"

// Inisialisasi interface Kendaraan
type Kendaraan interface {
	jarak(bensin int) int
	nama() string
}

// Motor
type Motor struct {
	name      string
	kecepatan int
}

func (m Motor) jarak(bensin int) int {
	return bensin * m.kecepatan
}

func (m Motor) nama() string {
	return m.name
}

// Mobil
type Mobil struct {
	name      string
	kecepatan int
}

func (m Mobil) jarak(bensin int) int {
	return bensin * m.kecepatan
}

func (m Mobil) nama() string {
	return m.name
}

// Bajaj
type Bajaj struct {
	name      string
	kecepatan int
}

func (b Bajaj) jarak(bensin int) int {
	return bensin * b.kecepatan
}

func (b Bajaj) nama() string {
	return b.name
}

// Fungsi untuk mencari kendaraan efisien
func kendaraanEfisien(bensin int, kendaraan ...Kendaraan) (Kendaraan, int) {
	var Kendaraan_ Kendaraan
	maxJarak := 0

	for _, k := range kendaraan {
		jarak := k.jarak(bensin)
		if jarak > maxJarak {
			maxJarak = jarak
			Kendaraan_ = k
		}
	}
	return Kendaraan_, maxJarak
}

func main() {
	// inisialisasi struct
	motor := Motor{name: "Motor", kecepatan: 3}
	mobil := Mobil{name: "Mobil", kecepatan: 1}
	bajaj := Bajaj{name: "Bajaj", kecepatan: 4}

	bensin := 20

	Kendaraan, jarak := kendaraanEfisien(bensin, motor, mobil, bajaj)

	fmt.Printf("Kendaraan paling efisien: %s dengan jarak %d km\n", Kendaraan.nama(), jarak)
}
