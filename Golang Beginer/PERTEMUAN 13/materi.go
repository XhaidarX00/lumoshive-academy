package main

// import (
// 	"database/sql"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	connStr := "user=postgres dbname=latihan13 sslmode=disable password=postgres host=localhost"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Mengatur parameter pool koneksi
// 	db.SetMaxOpenConns(25)                  // Jumlah maksimal koneksi terbuka
// 	db.SetMaxIdleConns(25)                  // Jumlah maksimal koneksi idle
// 	db.SetConnMaxLifetime(30 * time.Minute) // Durasi maksimal koneksi (misalnya, 30 menit)
// 	db.SetConnMaxIdleTime(5 * time.Minute)  // Durasi maksimal koneksi idle (misalnya, 5 menit)

// 	// Test the connection to the database
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println("Successfully Connected")
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=latihan13 sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	// Insert data baru ke tabel users
// 	result, err := db.Exec("INSERT INTO users (first_name, last_name, email, birth_date, registration_date) VALUES ('Budi', 'Santoso', 'budi@example.com', '1990-05-15', NOW())")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Rows inserted: %d\n", rowsAffected)

// 	// call function get all user
// 	GetAllUser(db)
// }

// // get all user
// func GetAllUser(db *sql.DB) {
// 	// Mengambil data dari tabel users dengan Query
// 	rows, err := db.Query("SELECT id, first_name, last_name, email, birth_date, registration_date FROM users")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// Membaca hasil query
// 	for rows.Next() {
// 		var id int
// 		var firstName, lastName, email string
// 		var birthDate, registrationDate sql.NullTime

// 		err := rows.Scan(&id, &firstName, &lastName, &email, &birthDate, &registrationDate)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("ID: %d\n", id)
// 		fmt.Printf("Nama: %s %s\n", firstName, lastName)
// 		fmt.Printf("Email: %s\n", email)
// 		if birthDate.Valid {
// 			fmt.Printf("Tanggal Lahir: %s\n", birthDate.Time.Format("2006-01-02"))
// 		} else {
// 			fmt.Println("Tanggal Lahir: <NULL>")
// 		}
// 		if registrationDate.Valid {
// 			fmt.Printf("Tanggal Pendaftaran: %s\n", registrationDate.Time.Format("2006-01-02 15:04:05"))
// 		} else {
// 			fmt.Println("Tanggal Pendaftaran: <NULL>")
// 		}
// 		fmt.Println("-------------------")
// 	}

// 	// Periksa apakah ada error selama proses pembacaan
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// type User struct {
// 	ID               int64
// 	Username         sql.NullString
// 	Age              sql.NullInt64
// 	Email            sql.NullString
// 	BirthDate        sql.NullTime
// 	RegistrationDate sql.NullTime
// }

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=latihan13 sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Query untuk mengambil semua data dari tabel users
// 	rows, err := db.Query("SELECT id, username, age, email, birth_date, registration_date FROM users")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// Membaca hasil query dan memetakkannya ke struct User
// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.Username, &user.Age, &user.Email, &user.BirthDate, &user.RegistrationDate)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, user)
// 	}

// 	// Memeriksa error selama pembacaan
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Menampilkan data pengguna
// 	for _, u := range users {
// 		fmt.Printf("ID: %d\n", u.ID)
// 		if u.Username.Valid {
// 			fmt.Printf("Username: %s\n", u.Username.String)
// 		} else {
// 			fmt.Println("Username: <NULL>")
// 		}
// 		if u.Age.Valid {
// 			fmt.Printf("Age: %d\n", u.Age.Int64)
// 		} else {
// 			fmt.Println("Age: <NULL>")
// 		}

// 		if u.Email.Valid {
// 			fmt.Printf("Email: %s\n", u.Email.String)
// 		} else {
// 			fmt.Println("Email: <NULL>")
// 		}
// 		if u.BirthDate.Valid {
// 			fmt.Printf("Birth Date: %s\n", u.BirthDate.Time.Format("2006-01-02"))
// 		} else {
// 			fmt.Println("Birth Date: <NULL>")
// 		}
// 		if u.RegistrationDate.Valid {
// 			fmt.Printf("Registration Date: %s\n", u.RegistrationDate.Time.Format("2006-01-02 15:04:05"))
// 		} else {
// 			fmt.Println("Regis Date: <NULL>")
// 		}
// 		fmt.Println("-------------------")
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// type User struct {
// 	ID       int
// 	Username string
// 	Password string
// 	Email    string
// }

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Username dan password yang diterima dari input pengguna (misalnya dari form login)
// 	username := "admin"
// 	password := "password' OR '1'='1"

// 	// Query yang tidak aman
// 	// query := "SELECT id, username, password, email FROM customers WHERE username='" + username + "' AND password='" + password + "'"
// 	query := fmt.Sprintf("SELECT id, username, password, email FROM customers WHERE username='%s' AND password='%s'", username, password)
// 	fmt.Println("Executing query:", query)

// 	row := db.QueryRow(query)

// 	var user User
// 	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No user found with the provided credentials.")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("User found: %+v\n", user)
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// type Customer struct {
// 	ID       int
// 	Username string
// 	Password string
// 	Email    string
// }

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=latihan13 sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Input pengguna yang aman dari SQL Injection
// 	username := "admin"
// 	password := "adminpassword"

// 	// Query yang aman menggunakan parameterized query
// 	query := "SELECT id, username, password, email FROM customers WHERE username=$1 AND password=$2"
// 	fmt.Println("Executing query:", query)

// 	row := db.QueryRow(query, username, password)

// 	var customer Customer
// 	err = row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No customer found with the provided credentials.")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("Customer found: %+v\n", customer)
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq" // PostgreSQL driver
// )

// type Customer struct {
// 	ID       int
// 	Username string
// 	Password string
// 	Email    string
// }

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=latihan13 sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// untuk mendapatkan ID dari baris yang dimasukkan
// 	var lastInsertId int
// 	err = db.QueryRow("INSERT INTO customers (username, password, email) VALUES ($1, $2, $3) RETURNING id",
// 		"dikcy", "123456", "dicky@example.com").Scan(&lastInsertId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Last Inserted ID: %d\n", lastInsertId)

// 	stat, err := db.Prepare("SELECT id, username, password, email FROM customers WHERE username=$1 AND password=$2")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stat.Close()

// 	username := "admin"
// 	password := "adminpassword"

// 	row := stat.QueryRow(username, password)
// 	var customer Customer

// 	err = row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No customers found with the provided credentials")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("Customer found : %v\n", customer)
// 	}

// 	username = "customer1"
// 	password = "password1"

// 	row = stat.QueryRow(username, password)

// 	err = row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No customers found with the provided credentials")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("Customer found : %v\n", customer)
// 	}
// }

// database transaction
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"main/auth"

// 	_ "github.com/lib/pq"
// )

// // OrderItem mewakili item dalam pesanan

// type OrderItem struct {
// 	ProductID int
// 	Quantity  int
// 	Price     float64
// }

// func main() {

// 	// Koneksi ke database PostgreSQL
// 	authInstance := &auth.Auth{}
// 	db, err := authInstance.Inisialisasi()
// 	if err != nil {
// 		log.Fatal("Gagal menghubungkan ke database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	fmt.Println("Berhasil terhubung ke database!")

// 	// Membuat pesanan
// 	userID := 1
// 	orderItems := []OrderItem{
// 		{ProductID: 1, Quantity: 2, Price: 50.00},
// 		{ProductID: 2, Quantity: 1, Price: 30.00},
// 	}

// 	orderID, err := createOrder(db, userID, orderItems)
// 	if err != nil {

// 		log.Fatal(err)

// 	}
// 	fmt.Printf("Order created successfully with ID: %d\n", orderID)

// }

// func createOrder(db *sql.DB, userID int, items []OrderItem) (int, error) {
// 	// Memulai transaksi
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Membuat entri baru dalam tabel orders
// 	var orderID int
// 	err = tx.QueryRow("INSERT INTO orders (user_id) VALUES ($1) RETURNING id", userID).Scan(&orderID)
// 	if err != nil {
// 		tx.Rollback()
// 		return 0, err
// 	}

// 	// Memproses setiap item dalam pesanan
// 	for _, item := range items {
// 		// Memperbarui stok produk
// 		_, err := tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1", item.Quantity, item.ProductID)
// 		if err != nil {
// 			tx.Rollback()
// 			return 0, err
// 		}

// 		// Menambahkan item ke tabel order_items
// 		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderID, item.ProductID, item.Quantity, item.Price)
// 		if err != nil {
// 			tx.Rollback()
// 			return 0, err
// 		}

// 	}

// 	// Menyelesaikan transaksi
// 	err = tx.Commit()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return orderID, nil
// }

// repository pattern
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// // User merupakan entitas yang akan disimpan di database
// type User struct {
// 	ID               int
// 	FirstName        string
// 	LastName         string
// 	Email            string
// 	BirthDate        time.Time
// 	RegistrationDate time.Time
// }

// // UserRepository merupakan antarmuka untuk mengakses data User
// type UserRepository interface {
// 	Create(user *User) error
// 	Update(user *User) error
// 	Delete(id int) error
// 	GetByID(id int) (*User, error)
// 	GetAll() ([]*User, error)
// }

// // UserRepositoryDB adalah implementasi dari UserRepository menggunakan database SQL
// type UserRepositoryDB struct {
// 	DB *sql.DB
// }

// // Create akan membuat user baru di database
// func (r *UserRepositoryDB) Create(user *User) error {
// 	query := `INSERT INTO users (first_name, last_name, email, birth_date, registration_date)
//               VALUES ($1, $2, $3, $4, $5) RETURNING id`
// 	err := r.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.BirthDate, user.RegistrationDate).Scan(&user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Update akan memperbarui data user di database berdasarkan ID
// func (r *UserRepositoryDB) Update(user *User) error {
// 	query := `UPDATE users SET first_name=$1, last_name=$2, email=$3, birth_date=$4, registration_date=$5 WHERE id=$6`
// 	_, err := r.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.BirthDate, user.RegistrationDate, user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Delete akan menghapus data user dari database berdasarkan ID
// func (r *UserRepositoryDB) Delete(id int) error {
// 	query := "DELETE FROM users WHERE id=$1"
// 	_, err := r.DB.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // GetByID akan mengembalikan data user dari database berdasarkan ID
// func (r *UserRepositoryDB) GetByID(id int) (*User, error) {
// 	query := "SELECT id, first_name, last_name, email, birth_date, registration_date FROM users WHERE id=$1"
// 	row := r.DB.QueryRow(query, id)
// 	user := &User{}
// 	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.RegistrationDate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// // GetAll akan mengembalikan semua data user dari database
// func (r *UserRepositoryDB) GetAll() ([]*User, error) {
// 	query := "SELECT id, first_name, last_name, email, birth_date, registration_date FROM users"
// 	rows, err := r.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []*User{}
// 	for rows.Next() {
// 		user := &User{}
// 		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.RegistrationDate)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

// // Inisialisasi repository
// func NewUserRepository(db *sql.DB) UserRepository {
// 	return &UserRepositoryDB{DB: db}
// }

// func main() {
// 	// Koneksi ke database
// 	authInstance := &auth.Auth{}
// 	db, err := authInstance.Inisialisasi()
// 	if err != nil {
// 		log.Fatal("Gagal menghubungkan ke database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	fmt.Println("Berhasil terhubung ke database!")

// 	// Menginisialisasi repository
// 	userRepo := NewUserRepository(db)

// 	// Contoh penggunaan repository
// 	// Membuat user baru
// 	user := &User{
// 		FirstName:        "John",
// 		LastName:         "Doe",
// 		Email:            "john.doe@example.com",
// 		BirthDate:        time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
// 		RegistrationDate: time.Now(),
// 	}
// 	err = userRepo.Create(user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("User created successfully with ID: %d\n", user.ID)

// 	// Mendapatkan semua user
// 	users, err := userRepo.GetAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("All users:")
// 	for _, u := range users {
// 		fmt.Printf("ID: %d, Name: %s %s, Email: %s\n", u.ID, u.FirstName, u.LastName, u.Email)
// 	}

// 	// Mendapatkan user berdasarkan ID
// 	userByID, err := userRepo.GetByID(1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("User found by ID: %d - Name: %s %s, Email: %s\n", userByID.ID, userByID.FirstName, userByID.LastName, userByID.Email)

// 	// Menghapus user
// 	err = userRepo.Delete(user.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("User with ID %d has been deleted\n", user.ID)
// }
