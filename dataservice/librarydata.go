package dataservice

import (
	"database/sql"
	"goTraining/models"
)

func CreateBook(db *sql.DB, book models.Book) error {

	query := `INSERT INTO books(title, author, year) VALUES(?, ?, ?)`
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil

}
