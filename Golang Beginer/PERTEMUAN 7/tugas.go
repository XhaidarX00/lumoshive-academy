package main

import (
	"fmt"
	"main/utils"
	"math/rand"
	"time"
)

func main() {
	suhuChan := make(chan int, 1)
	kelembapanChan := make(chan int, 1)
	tekananChan := make(chan int, 1)
	done := make(chan bool)

	ticker := time.NewTicker(2 * time.Second)
	batasSensor := BatasSensor{suhu: 20, kelembapan: 30, tekanan: 25}

	go sensor("Sensor Suhu", suhuChan, done)
	go sensor("Sensor Kecepatan", kelembapanChan, done)
	go sensor("Sensor Tekanan", tekananChan, done)

	fmt.Println("\n ========= Mengambil data sensor ========= \n ")
	for {
		select {
		case val := <-suhuChan:
			utils.SuccesMessage(fmt.Sprintf("Data dari Sensor Suhu: %d\n", val))
			batasSensor.cekBatasan(val, "suhu")
		case val := <-kelembapanChan:
			utils.SuccesMessage(fmt.Sprintf("Data dari Sensor Kelembapan: %d\n", val))
			batasSensor.cekBatasan(val, "kelembapan")
		case val := <-tekananChan:
			utils.SuccesMessage(fmt.Sprintf("Data dari Sensor Tekanan: %d\n", val))
			batasSensor.cekBatasan(val, "tekanan")
		case <-done:
			fmt.Println("\n ========= Sensor selesai  ========= \n ")
			return
		case <-ticker.C:
			fmt.Println("\n ========= Mengambil data sensor ========= \n ")
		}
	}
}

type BatasSensor struct {
	suhu       int
	kelembapan int
	tekanan    int
}

func (s BatasSensor) cekBatasan(angka int, data string) {
	if data == "suhu" && angka > s.suhu {
		utils.ErrorMessage("Data melebihi batasan Suhu.")
	}

	if data == "kelembapan" && angka > s.kelembapan {
		utils.ErrorMessage("Data melebihi batasan Kelembapan.")
	}

	if data == "tekanan" && angka > s.tekanan {
		utils.ErrorMessage("Data melebihi batasan Tekanan.")
	}
}

func sensor(name string, ch chan<- int, done chan<- bool) {
	for {
		select {
		case <-time.After(5 * time.Second):
			utils.ErrorMessage(fmt.Sprintf("%s timeout", name))
			done <- true
			return
		case ch <- rand.Intn(100):
			time.Sleep(2 * time.Second)
		}
	}
}
