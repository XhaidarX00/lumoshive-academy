package main

import "fmt"

func main_latihan() {
	fmt.Println("==========\nLATIHAN\n==========\n ")
	// buatkan sebuah struct dengan propert nama, score
	// buatkan fungsi untuk merubah score dari properti struct tersebut

	// player := Player{name: "Haidar", score: 10}
	// fmt.Println("Name : ", player.name)
	// fmt.Println("Score : ", player.score)

	// player.changeValuePlayer(20)
	// fmt.Println("Score : ", player.score)

	// buatkan struct lalu lakukan inisialisasi menggunakan slice
	// var person []people

	// person = append(person, people{name: "Haidar", position: "Web Dev", address: "Indonesia"})
	// person = append(person, people{name: "Darmi", position: "Mobile Dev", address: "Indonesia"})
	// person = append(person, people{name: "Darto", position: "Data Sience", address: "Indonesia"})

	// fmt.Println("----------------------------")
	// for _, people_data := range person {
	// 	fmt.Println("Name : ", people_data.name)
	// 	fmt.Println("Position : ", people_data.position)
	// 	fmt.Println("Addres : ", people_data.address)
	// 	fmt.Println("----------------------------")
	// }

	// buatkan 3 variabel yang memiliki akses terhadap satu variabel / dapat mengubah variabel 1
	// data := 100
	// x := &data
	// y := x
	// z := y

	// fmt.Println(data)
	// *x = 20
	// fmt.Println(data)
	// *y = 40
	// fmt.Println(data)
	// *z = 10
	// fmt.Println(data)

	// buatkan function untuk init suatu object struct
	// nama, alamat, no hp, lalu ditampung kedalam variabel agar bisa diprint
	// data_person := init_struct("Haidar", "Bogor", "082344")

	// fmt.Println("Name : ", data_person.name)
	// fmt.Println("Number : ", data_person.no)
	// fmt.Println("Address : ", data_person.addres)

	// buatkan satu func dengan satu parameter
	// dapat menerima tiga object
	// menghitung keliling dari masing" object tersebut
	// persegi := persegi{sisi: 5}
	// persegiPanjang := persegi_panjang{panjang: 10, lebar: 5}
	// lingkaran := lingkaran{jari_jari: 7}

	// print_hitung_keliling(persegi)
	// print_hitung_keliling(persegiPanjang)
	// print_hitung_keliling(lingkaran)

	// buat type data struct dimasukan kedalam variabel yang dapat mengubah inisialisasi struct tersebut
	// data := Person1{name: "Haidar"}
	// fmt.Println("Name Before : ", data.name)
	// dataChange := &data
	// dataChange.name = "Darmi"
	// fmt.Println("Name After : ", data.name)

	// buatkan struct kemudian lakukan init sebanyak 3 data useappand()
	// buatkan func untuk filter object struct tersebut berdasarkan tahun
	// kondisi : tampilkan 2 data jika tahun lebih dari 2005

	var person []Person_data

	person = append(person, Person_data{name: "Haidar1", year: 2010})
	person = append(person, Person_data{name: "Haidar2", year: 2003})
	person = append(person, Person_data{name: "Haidar3", year: 2007})

	print_data(person)
}

type Person_data struct {
	name string
	year int
}

func print_data(person []Person_data) {
	fmt.Println("------------------------")
	for _, person_ := range person {
		if person_.year > 2005 {
			fmt.Println("Name : ", person_.name)
			fmt.Println("Tahun : ", person_.year)
			fmt.Println("------------------------")
		}
	}
}

// type Person1 struct {
// 	name string
// }

// type Hitung_keliling interface {
// 	keliling() int
// }

// type persegi struct {
// 	sisi int
// }

// func (p persegi) keliling() int {
// 	return 4 * p.sisi
// }

// type persegi_panjang struct {
// 	panjang, lebar int
// }

// func (pp persegi_panjang) keliling() int {
// 	return 2 * (pp.panjang + pp.lebar)
// }

// type lingkaran struct {
// 	jari_jari int
// }

// func (l lingkaran) keliling() int {
// 	return 2 * l.jari_jari * 22 / 7
// }

// func print_hitung_keliling(h Hitung_keliling) {
// 	result := h.keliling()
// 	fmt.Println("Keliling : ", result)
// }

// type person struct {
// 	name, addres, no string
// }

// func init_struct(nama string, alamat string, no_hp string) person {
// 	data := person{
// 		name:   nama,
// 		addres: alamat,
// 		no:     no_hp,
// 	}

// 	return data
// }

// type Player struct {
// 	name  string
// 	score int
// }

// func (p *Player) changeValuePlayer(val int) {
// 	p.score = val
// }

// type people struct {
// 	name, position, address string
// }
