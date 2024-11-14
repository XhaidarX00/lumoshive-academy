//// validation one variable
// package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator/v10"
// )

// func main() {
// 	validator := validator.New()

// 	// validation one variable
// 	var username string = ""
// 	err := validator.Var(username, "required")

// 	if err != nil {
// 		fmt.Println("username validation failed:", err)
// 	} else {
// 		fmt.Println("username validation passed!")
// 	}
// }

// validation two variable
package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh validasi: Memastikan bahwa nilai konfirmasi password sama dengan password
	password := "secret123"
	confirmPassword := "secret123"

	// Menggunakan VarWithValue untuk memvalidasi confirmPassword terhadap password
	err := validate.VarWithValue(confirmPassword, password, "eqfield")
	if err != nil {
		fmt.Println("Confirm password validation failed:", err)
	} else {
		fmt.Println("Confirm password validation passed!")
	}
}
