package dataservice

import (
	"database/sql"
	"goTraining/models"
)

func CheckAvailability(db *sql.DB, product_id int, required_quantity int) (bool, error) {
	var quantity int
	err := db.QueryRow("SELECT quantity FROM product WHERE id = ?", product_id).Scan(&quantity)
	if err != nil {
		return false, err
	}
	return quantity >= required_quantity, nil
}

func GetProductId(db *sql.DB, ProductName string) (int, error) {
	var product_id int
	err := db.QueryRow("SELECT id FROM product WHERE name =?", ProductName).Scan(&product_id)
	if err != nil {
		return 0, err
	}
	return product_id, nil
}

func InsertCustomer(db *sql.DB, customer models.Customer) error {
	query := `INSERT INTO customer(name, phone, address, purchased,quantity) VALUES(?, ?, ?, ?, ?)`
	_, err := db.Exec(query, customer.Name, customer.Phone, customer.Address, customer.Purchased, customer.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func UpdateQuantity(db *sql.DB, product_id int, quantity int) error {
	query := `UPDATE product SET quantity = quantity - ? WHERE id = ?`
	_, err := db.Exec(query, quantity, product_id)
	if err != nil {
		return err
	}
	return nil

}
