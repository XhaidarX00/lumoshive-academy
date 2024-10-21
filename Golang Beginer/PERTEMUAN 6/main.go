package main

import (
	"materi/module"
	"materi/utils"
	"time"
)

var dataSlice []module.Data

func main() {
	doneSave := make(chan bool)
	doneSlice := make(chan []module.Data)

	go module.AddSlice(doneSave, dataSlice, 100, doneSlice)
	if <-doneSave {
		go module.PrintSucces(<-doneSlice)
	} else {
		utils.ErrorMessage("terjadi kesalahan menyimpan data")
	}

	time.Sleep(1 * time.Second)
}
