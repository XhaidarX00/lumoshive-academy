// package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator/v10"
// )

// // Definisikan struct dengan tag validasi
// type User struct {
// 	Name     string `validate:"required"`
// 	Email    string `validate:"required,email"`
// 	Age      int    `validate:"gte=0,lte=130"`
// 	Password string `validate:"required,min=8"`
// 	Address  string `validate:"required"`
// 	Phone    string `validate:"required,e164"` // e164 adalah format internasional untuk nomor telepon
// }

// func main() {
// 	// Inisialisasi validator
// 	validate := validator.New()

// 	// Contoh instance dari struct User
// 	user := &User{
// 		Name:     "Lumoshive Academy",
// 		Email:    "lumoshive@example.com",
// 		Age:      28,
// 		Password: "password123",
// 		Address:  "Grand Garden",
// 		Phone:    "+628709090930",
// 	}

// 	// // Validasi struct
// 	// err := validate.Struct(user)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// } else {
// 	// 	fmt.Println("Validation passed!")
// 	// }

// 	// Validasi struct
// 	err := validate.Struct(user)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			fmt.Printf("Error: Field %s failed validation with tag %s\n", err.Field(), err.Tag())
// 		}
// 	} else {
// 		fmt.Println("Validation passed!")
// 	}

// }

// package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator/v10"
// )

// // Definisikan struct dengan tag validasi
// type Event struct {
// 	Name            string `validate:"required"`
// 	Password        string `validate:"required"`
// 	ConfrimPassword string `validate:"required,eqfield=Password"`
// }

// func main() {
// 	// Inisialisasi validator
// 	validate := validator.New()

// 	// Contoh instance dari struct account
// 	event := &Event{
// 		Name:            "Lumoshive",
// 		Password:        "123",
// 		ConfrimPassword: "123",
// 	}

// 	// Validasi struct
// 	err := validate.Struct(event)
// 	if err != nil {
// 		if _, ok := err.(validator.ValidationErrors); ok {
// 			// Iterasi melalui errors dan tampilkan informasi detail
// 			for _, validationErr := range err.(validator.ValidationErrors) {
// 				fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Field(), validationErr.Tag())
// 				fmt.Printf("  Value: '%v'\n", validationErr.Value())
// 				fmt.Printf("  Condition: '%s'\n", validationErr.Param())
// 			}
// 		} else {
// 			// Error lainnya
// 			fmt.Println("Validation failed:", err)
// 		}
// 	} else {
// 		fmt.Println("Validation passed!")
// 	}

// }

package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan nested struct
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Zip    string `validate:"required"`
}

// Definisikan struct utama yang memiliki nested struct
type User struct {
	Name    string  `validate:"required"`
	Email   string  `validate:"required,email"`
	Age     int     `validate:"gte=0,lte=130"`
	Address Address `validate:"required"`
}

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh instance dari struct User dengan nested struct Address
	user := &User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Age:   30,
		Address: Address{
			Street: "123 Main St",
			City:   "",
			Zip:    "12345",
		},
	}

	// Validasi struct
	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// Iterasi melalui errors dan tampilkan informasi detail
			for _, validationErr := range err.(validator.ValidationErrors) {
				fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Namespace(), validationErr.Tag())
				fmt.Printf("  Value: '%v'\n", validationErr.Value())
				fmt.Printf("  Condition: '%s'\n", validationErr.Param())
			}
		} else {
			// Error lainnya
			fmt.Println("Validation failed:", err)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
