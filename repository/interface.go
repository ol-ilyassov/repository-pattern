package repository

import "github.com/ol-ilyassov/repository-pattern/model"

// Repository is an interface that holds key methods of working with local storage.
type Repository interface {
	Count() int
	GetProducts() []model.Product
	GetProductByID(id string) (int, model.Product, error)
	AddProduct(input model.ProductNew) (int, model.Product)
	UpdateProduct(id string, input model.ProductNew) (int, model.Product, error)
	DeleteProduct(id string) (int, error)
}
