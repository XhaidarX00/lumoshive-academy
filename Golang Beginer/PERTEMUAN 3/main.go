package main

import (
	"fmt"
	// "math"
)

func main_() {
	fmt.Println("Testing")

	// Pointer pendeklarasian
	// x := 10
	// y := &x
	// p := new(int)
	// *p = 42

	// fmt.Println("Nilai pointer p : ", *p)
	// fmt.Println("Alamat memori p : ", p)
	// fmt.Println("Nilai pointer x : ", *y)
	// fmt.Println("Alamat memori x : ", y)

	// pointer pada parameter function
	// z := 15
	// fmt.Println("Nilai Z : ", z)
	// changeValue(&z)
	// fmt.Println("Change Z : ", z)

	// struct adalah kumpulan dari satu atau lebih tipe data yang disimpan bersama sebagai satu unit
	// Person1 := Person{name: "haidar", age: 20}

	// fmt.Println("Name : ", Person1.name)
	// fmt.Println("Name : ", Person1.age)

	// // struct variabel objeck pointer
	// personPtr := &Person1
	// personPtr.name = "Iskandar"
	// personPtr.age = 24

	// fmt.Println("Name : ", personPtr.name)
	// fmt.Println("Name : ", personPtr.age)

	// struct methode
	// Rectangle_ := Rectangle{width: 5, height: 20}
	// fmt.Println("Hasil : ", Rectangle_.hitung_luas())

	// struct methode pointer
	// Rectangle_ := Rectangle{width: 5, height: 20}
	// fmt.Println("lebar : ", Rectangle_.width)

	// Rectangle_.setwidth(200)
	// fmt.Println("lebar : ", Rectangle_.width)

	// embeding struct -> memungkinkan untuk menyimpan satu struct didalam struct lain
	// emp := Employes{
	// 	Person:   Person{name: "Haidar", age: 24},
	// 	position: "Developer Golang",
	// }

	// fmt.Println("Name : ", emp.name)
	// fmt.Println("Age : ", emp.age)
	// fmt.Println("Position : ", emp.position)

	// anosymus struct -> struct tanpa nama dan digunakan langsung ditempat
	// person := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "Darmi",
	// 	age:  20,
	// }

	// fmt.Println("Name : ", person.name)
	// fmt.Println("Age : ", person.age)

	// kombinasi slice dan struct -> untuk menyimpan banyak instansi dari tipe struct yang sama
	// var people []Person

	// people = append(people, Person{name: "damri", age: 20})
	// people = append(people, Person{name: "damri2", age: 22})

	// for _, p := range people {
	// 	fmt.Println("Name : ", p.name)
	// 	fmt.Println("Age : ", p.age)
	// }

	// interface -> kumpulan dari methode-methode yang mendeklarasikan tanpa implementasi
	// dog := dog{}
	// cat := cat{}

	// makeAnimalSpeak(dog)
	// makeAnimalSpeak(cat)

	// embed interface didalam interface lain
	// c := Cylinder{Radius: 3, Height: 5}
	// printSolidDetails(c)

	// interface any -> dapat menampung tipe data apapun
	printValue(47)
	printValue("Hello")
	printValue("3.14")
	printValue([]int{1, 2, 3, 4, 5})
}

// func changeValue(val *int) {
// 	*val = 30
// }

// type Person struct {
// 	name string
// 	age  int
// }

// type Rectangle struct {
// 	width  float64
// 	height float64
// }

// methode untuk menghitung luas persegi panjang
// func (name_var Rectangle) hitung_luas() float64 {
// 	return name_var.height * name_var.width
// }

// methode untuk menghitung luas persegi panjang menggunaka pointer
// func (name_var *Rectangle) setwidth(width float64) {
// 	name_var.width = width
// }

// type Employes struct {
// 	Person
// 	position string
// }

// type animal interface {
// 	speak() string
// }

// type dog struct{}

// func (d dog) speak() string {
// 	return "Woof!"
// }

// type cat struct{}

// func (c cat) speak() string {
// 	return "Meow!"
// }

// func makeAnimalSpeak(a animal) {
// 	fmt.Println(a.speak())
// }

// type Shape interface {
// 	Area() float64
// }

// type Volume interface {
// 	Volume() float64
// }

// type Solid interface {
// 	Shape
// 	Volume
// }

// type Cylinder struct {
// 	Radius, Height float64
// }

// func (c Cylinder) Area() float64 {
// 	return 2*math.Pi*c.Radius*c.Height + 2*math.Pi*c.Radius*c.Radius
// }

// func (c Cylinder) Volume() float64 {
// 	return math.Pi * c.Radius * c.Radius * c.Height
// }

// func printSolidDetails(s Solid) {
// 	fmt.Printf("Area : %.2f\n", s.Area())
// 	fmt.Printf("Volume : %.2f\n", s.Volume())
// }

func printValue(val interface{}) {
	fmt.Println("Value : ", val)
}
