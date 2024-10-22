package main

import (
	"fmt"
	"materi/module"
	"materi/utils"
	"strconv"
	"time"
)

var dataSlice []module.Data

func main() {
	var inputInt int
	for {
		var input string
		fmt.Printf("Masukan data : ")
		fmt.Scan(&input)

		utils.ClearScreen()

		inputStr, err := strconv.Atoi(input)
		inputInt = inputStr
		if err != nil {
			utils.ErrorMessage("Data harus harus berupa angka")
			continue
		} else {
			break
		}
	}

	doneSave := make(chan bool)
	doneSlice := make(chan []module.Data)

	go module.AddSlice(doneSave, dataSlice, inputInt, doneSlice)
	if <-doneSave {
		go module.PrintSucces(<-doneSlice)
	} else {
		utils.ErrorMessage("terjadi kesalahan menyimpan data")
	}

	time.Sleep(1 * time.Second)
}
