// package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator/v10"
// )

// // Definisikan struct dengan field yang menggunakan tag validasi dengan rule or
// type User struct {
// 	Username string `validate:"required,email|phone"` // Username bisa berupa email atau nomor telepon
// }

// // Custom validation function for phone numbers
// func isPhone(fl validator.FieldLevel) bool {
// 	// Sederhana: memeriksa apakah string hanya terdiri dari angka dan panjangnya antara 10 hingga 15 karakter
// 	phone := fl.Field().String()
// 	if len(phone) >= 10 && len(phone) <= 15 {
// 		for _, char := range phone {
// 			if char < '0' || char > '9' {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	// Inisialisasi validator
// 	validate := validator.New()

// 	// Daftarkan custom validation untuk phone
// 	validate.RegisterValidation("phone", isPhone)

// 	// Contoh instance dari struct User
// 	user := &User{
// 		// Username: "1234567890", // Ini adalah nomor telepon yang valid
// 		Username: "user@example.com", // Ini adalah email yang valid
// 	}

// 	// Validasi struct
// 	err := validate.Struct(user)
// 	if err != nil {
// 		if _, ok := err.(validator.ValidationErrors); ok {
// 			// Iterasi melalui errors dan tampilkan informasi detail
// 			for _, validationErr := range err.(validator.ValidationErrors) {
// 				fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Namespace(), validationErr.Tag())
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

// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/go-playground/validator/v10"
// )

// // Fungsi validasi khusus untuk memeriksa apakah field Password dan ConfirmPassword sama
// func passwordMatch(fl validator.FieldLevel) bool {
// 	confirmPasswordField, _, _, ok := fl.GetStructFieldOK2()
// 	if !ok {
// 		return false
// 	}
// 	confirmPassword := strings.ToUpper(confirmPasswordField.String())
// 	password := strings.ToUpper(fl.Field().String())
// 	return password == confirmPassword
// }

// // Definisikan struct dengan field yang menggunakan custom tag validasi
// type User struct {
// 	Password        string `validate:"required"`
// 	ConfirmPassword string `validate:"required,password_match=Password"` // Tag custom validation
// }

// func main() {
// 	// Inisialisasi validator
// 	validate := validator.New()

// 	// Daftarkan fungsi validasi khusus
// 	validate.RegisterValidation("password_match", passwordMatch)

// 	// Contoh instance dari struct User
// 	user := &User{
// 		Password:        "secret123",
// 		ConfirmPassword: "secret123", // Harus sama dengan Password untuk validasi berhasil
// 	}

// 	// Validasi struct
// 	err := validate.Struct(user)
// 	if err != nil {
// 		if _, ok := err.(validator.ValidationErrors); ok {
// 			// Iterasi melalui errors dan tampilkan informasi detail
// 			for _, validationErr := range err.(validator.ValidationErrors) {
// 				fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Namespace(), validationErr.Tag())
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

// Struct dengan beberapa field yang perlu divalidasi secara keseluruhan
type User struct {
	Password        string `validate:"required"`
	ConfirmPassword string `validate:"required"`
	Age             int    `validate:"gte=0,lte=130"`
}

// Fungsi validasi untuk seluruh struct User
func userStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	// Memastikan bahwa Password dan ConfirmPassword harus sama
	if user.Password != user.ConfirmPassword {
		sl.ReportError(user.ConfirmPassword, "ConfirmPassword", "ConfirmPassword", "konfirmasi password salah", "")
	}

	// Contoh validasi lainnya: memastikan usia harus lebih dari 18 tahun jika password diisi
	if user.Password != "" && user.Age < 18 {
		sl.ReportError(user.Age, "Age", "Age", "usia kurang dari 18 tahun", "")
	}
}

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Daftarkan fungsi validasi khusus untuk struct User
	validate.RegisterStructValidation(userStructLevelValidation, User{})

	// Contoh instance dari struct User
	user := &User{
		Password:        "secret123",
		ConfirmPassword: "secret124",
		Age:             17, // Ini akan gagal karena usia kurang dari 18 tahun
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
