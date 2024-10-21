package module

import (
	"fmt"
	"materi/utils"
	"reflect"
)

type Data struct {
	angka any
}

func AddSlice(doneSave chan<- bool, dataSlice []Data, count int, doneSlice chan<- []Data) {
	for i := 0; i < count; i++ {
		dataSlice = append(dataSlice, Data{angka: i + 1})
	}
	dataSlice = append(dataSlice, Data{angka: "madu"})
	dataSlice = append(dataSlice, Data{angka: "macan"})
	doneSave <- true
	doneSlice <- dataSlice
}

func PrintSucces(dataSlice []Data) {
	if len(dataSlice) == 0 {
		utils.ErrorMessage("dataSlice kosong, tidak ada data yang disimpan")
		return
	}

	for _, data := range dataSlice {
		if reflect.TypeOf(data.angka).Kind() != reflect.Int {
			utils.ErrorMessage(fmt.Sprintf("tipe %v bukan int, tetapi %s", data.angka, reflect.TypeOf(data.angka).Kind()))
		} else {
			utils.SuccesMessage(fmt.Sprintf("Data berhasil %d disimpan didalam dataSlice", data.angka))
		}
	}
}
