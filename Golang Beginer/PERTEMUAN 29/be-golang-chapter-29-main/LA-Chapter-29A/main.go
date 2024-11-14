package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	validator := validator.New()

	if validator != nil {
		fmt.Println("success create validator")
	}
}
