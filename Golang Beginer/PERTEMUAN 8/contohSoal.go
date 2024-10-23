// package main

// import (
// 	"context"
// 	"fmt"
// 	"main/utils"
// 	"net/http"
// 	"time"
// )

// // Struct untuk menyimpan data pengguna
// type User struct {
// 	Username string
// 	Password string
// }

// // Slice untuk menyimpan daftar pengguna
// var users = []User{}

// // Fungsi untuk mendaftarkan pengguna (menggunakan Scan)
// func register() {
// 	var username, password string
// 	for {

// 		fmt.Println("=== Register ===")
// 		fmt.Print("Masukkan username: ")
// 		fmt.Scan(&username)
// 		utils.ClearScreen()

// 		// Cek apakah username sudah terdaftar
// 		if isUserRegistered(username) {
// 			utils.ErrorMessage("Username sudah terdaftar! Silakan login.")
// 			continue
// 		}

// 		for {
// 			fmt.Print("Masukkan password: ")
// 			fmt.Scan(&password)
// 			utils.ClearScreen()
// 			if utils.IsLenVar(password, 6) {
// 				break
// 			} else {
// 				utils.ErrorMessage("Password harus terdiri dari 6 karakter atau lebih. Silakan coba lagi")
// 			}
// 		}

// 		break
// 	}

// 	// Simpan pengguna baru ke dalam slice
// 	users = append(users, User{Username: username, Password: password})
// 	utils.SuccesMessage(fmt.Sprintln("Registrasi berhasil! Username:", username))
// 	utils.SuccesMessage(fmt.Sprintf("Berikut link cli login : \ncurl.exe -X POST \"http://localhost:8080/login\" -d \"username=%s&password=%s\"", username, password))

// }

// // Fungsi untuk login
// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Hanya mendukung POST", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Menggunakan context dengan timeout 2 detik
// 	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
// 	defer cancel()

// 	// Mendekode data login dari request
// 	var username, password string
// 	username = r.FormValue("username")
// 	password = r.FormValue("password")

// 	// Cek login
// 	select {
// 	case <-time.After(1 * time.Second):
// 		if validUser(username, password) {
// 			CapUsername := utils.Capitalize(username)
// 			succes_ := utils.ColorMessage("green", "Succes : ")
// 			fmt.Fprintf(w, "%s Login berhasil, Selamat datang %s!", succes_, CapUsername)
// 		} else if isUserRegistered(username) {
// 			err_ := utils.ColorMessage("red", "Error : ")
// 			msg := fmt.Sprintf("%s Username sudah terdaftar, tapi password salah!", err_)
// 			http.Error(w, msg, http.StatusUnauthorized)
// 		} else {
// 			err_ := utils.ColorMessage("red", "Error : ")
// 			msg := fmt.Sprintf("%s Username tidak ditemukan", err_)
// 			http.Error(w, msg, http.StatusUnauthorized)
// 		}
// 	case <-ctx.Done():
// 		err_ := utils.ColorMessage("red", "Error : ")
// 		msg := fmt.Sprintf("%s Waktu habis", err_)
// 		http.Error(w, msg, http.StatusRequestTimeout)
// 	}
// }

// // Fungsi untuk validasi login
// func validUser(username, password string) bool {
// 	// Cek apakah pengguna dengan username dan password sesuai ada di slice
// 	for _, user := range users {
// 		if user.Username == username && user.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Fungsi untuk cek apakah username sudah terdaftar
// func isUserRegistered(username string) bool {
// 	for _, user := range users {
// 		if user.Username == username {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	// Registrasi user menggunakan scan sebelum menjalankan server
// 	register()

// 	// Menyiapkan API untuk login
// 	http.HandleFunc("/login", loginHandler)

// 	// Menjalankan server
// 	utils.SuccesMessage("Server berjalan di http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

// // curl -X POST "http://localhost:8080/login" -d "username=testuser&password=testpass"

package main

import (
	"fmt"
	"time"
)

func main_contoh() {
	ticker := time.NewTicker(2 * time.Second) // Mengatur ticker dengan interval 2 detik
	defer ticker.Stop()                       // Pastikan ticker dihentikan setelah selesai

	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second) // Simulasi pekerjaan selama 10 detik
		done <- true
	}()

	lastResponseTime := time.Now()

	for {
		select {
		case <-ticker.C:
			// Ketika ticker mengeluarkan sinyal
			fmt.Println("Mengambil data...")
			lastResponseTime = time.Now() // Memperbarui waktu terakhir
			fmt.Println("Waktu terakhir diambil:", lastResponseTime)
			fmt.Println("Waktu after baru:", lastResponseTime.Add(5*time.Second))

		case <-time.After(time.Until(lastResponseTime.Add(5 * time.Second))):
			// Setelah 5 detik dari lastResponseTime, ambil data terakhir
			fmt.Println("Mengambil data terakhir setelah 5 detik dari waktu terakhir...")
			fmt.Println("Waktu after baru:", time.Until(lastResponseTime.Add(5*time.Second)))

		case <-done:
			fmt.Println("Pekerjaan selesai. Keluar dari loop.")
			return
		}
	}
}
