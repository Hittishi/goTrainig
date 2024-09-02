package models

type Customer struct {
	Name      string `json:"name"`
	Phone     int    `json:"phone"`
	Address   string `json:"address"`
	Purchased string `json:"purchased"`
	Quantity  int    `json:"quantity"`
}
type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
type Order struct {
	Name        string `json:"name"`
	Phone       int    `json:"phone"`
	Address     string `json:"address"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
