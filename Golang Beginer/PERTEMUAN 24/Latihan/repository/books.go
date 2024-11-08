package repository

import (
	"latihan/model"
	"latihan/model/books"
)

func (r *Repository) GetBookDataRepo(data *[]books.Book) error {
	query := "SELECT * FROM books ORDER BY id"
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var book books.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Type, &book.Author, &book.Price, &book.Discount); err != nil {
			panic(err)
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
				return err
			}
		case 1:
			err := r.DB.QueryRow(q).Scan(&data.Total_Sales)
			if err != nil {
				return err
			}
		case 2:
			err := r.DB.QueryRow(q).Scan(&data.Total_Books)
			if err != nil {
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
		return err
	}

	return nil
}
