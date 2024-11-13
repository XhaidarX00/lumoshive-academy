package repository

import (
	"latihan/model"
	"latihan/model/books"

	"go.uber.org/zap"
)

func (r *Repository) GetBookDataRepo(data *[]books.Book) error {
	query := "SELECT * FROM books ORDER BY id"
	rows, err := r.DB.Query(query)
	if err != nil {
		r.Logger.Error("Error GetBooksDataRepo", zap.Error(err))
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var book books.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Type, &book.Author, &book.Price, &book.Discount); err != nil {
			r.Logger.Error("Error GetBooksDataRepo", zap.Error(err))
			return err
		}

		book.Rating, err = r.calculateAverageRating(book.ID)
		if err != nil {
			r.Logger.Error("Error GetBooksDataRepo", zap.Error(err))
			return err
		}

		*data = append(*data, book)
	}

	return nil
}

func (r *Repository) EditBookDataRepo(book books.Book) error {
	query := `
		UPDATE books 
		SET 
			name = $1, 
			type = $2, 
			author = $3, 
			price = $4, 
			discount = $5 
		WHERE id = $6
	`

	_, err := r.DB.Exec(query, book.Name, book.Type, book.Author, book.Price, book.Discount, book.ID)
	if err != nil {
		r.Logger.Error("Error EditBooks", zap.Error(err))
		return err
	}

	return nil
}

func (r *Repository) AddBookDataRepo(book books.Book) error {
	query := `
        INSERT INTO books 
        (id, name, type, author, price, discount)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	_, err := r.DB.Exec(query, book.ID, book.Name, book.Type, book.Author, book.Price, book.Discount)
	if err != nil {
		r.Logger.Error("Error AddBooks", zap.Error(err))
		return err
	}

	return nil
}

func (r *Repository) GetDhasboardDataRepo(data *model.GetDhasboardData) error {
	query := []string{
		"SELECT MAX(rating) AS highest_rating FROM reviews",
		"SELECT COUNT(id) FROM orders",
		"SELECT COUNT(id) FROM books",
	}

	for i, q := range query {
		switch i {
		case 0:
			err := r.DB.QueryRow(q).Scan(&data.Highest_Rating)
			if err != nil {
				r.Logger.Error("Error GetDashboard", zap.Error(err))
				return err
			}
		case 1:
			err := r.DB.QueryRow(q).Scan(&data.Total_Sales)
			if err != nil {
				r.Logger.Error("Error GetDashboard", zap.Error(err))
				return err
			}
		case 2:
			err := r.DB.QueryRow(q).Scan(&data.Total_Books)
			if err != nil {
				r.Logger.Error("Error GetDashboard", zap.Error(err))
				return err
			}
		}
	}

	return nil
}

func (r *Repository) DeleteBookRepo(id string) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		r.Logger.Error("Error DeleteBook", zap.Error(err))
		return err
	}

	return nil
}

// Fungsi untuk menghitung rata-rata rating
func (r *Repository) calculateAverageRating(bookID string) (float64, error) {
	var count int
	var avgRating float64

	// Hitung jumlah data
	err := r.DB.QueryRow("SELECT COUNT(*) FROM reviews").Scan(&count)
	if err != nil {
		return 0, err
	}

	if count <= 10 {
		err = r.DB.QueryRow("SELECT COALESCE(ROUND(AVG(rating), 1), 0) FROM reviews WHERE book_id = $1", bookID).Scan(&avgRating)
	} else {
		err = r.DB.QueryRow(`
			SELECT COALESCE(ROUND(AVG(rating), 1), 0) 
			FROM (
				SELECT rating 
				FROM reviews 
				WHERE book_id = $1
				ORDER BY review_date DESC 
				LIMIT 25
			) AS latest_reviews
		`, bookID).Scan(&avgRating)
	}

	if err != nil {
		return 0, err
	}
	return avgRating, nil
}
