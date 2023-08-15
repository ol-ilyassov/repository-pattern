package model

// ProductNew is used to store input fields for product data
type ProductNew struct {
	Name        string
	Description string
	Price       float64
	Stock       uint
}
