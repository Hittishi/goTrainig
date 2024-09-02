package api

import (
	"database/sql"
	"errors"
	"goTraining/dataservice"
	"goTraining/models"
)

func CreateBookLogic(db *sql.DB, book models.Book) error {
	dataservice.CreateBook(db, book)
	return nil

}
func ProcessOrderLogic(db *sql.DB, o models.Order) error {
	product_id, err := dataservice.GetProductId(db, o.ProductName)
	if err != nil {
		return err
	}

	available, err := dataservice.CheckAvailability(db, product_id, o.Quantity)
	if err != nil {
		return err
	}
	if !available {
		return errors.New("product not available in the required quantity")
	}

	var customer models.Customer
	customer = models.Customer{
		Name:      o.Name,
		Phone:     o.Phone,
		Address:   o.Address,
		Purchased: o.ProductName,
		Quantity:  o.Quantity,
	}

	if err := dataservice.InsertCustomer(db, customer); err != nil {
		return err
	}

	if err := dataservice.UpdateQuantity(db, product_id, o.Quantity); err != nil {
		return err
	}

	return nil
}
