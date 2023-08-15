package repository

import "github.com/ol-ilyassov/repository-pattern/model"

// Repository holds local storage instance.
type ProductRepository struct {
	database []model.Product
}

// New is a function that returns Repository intance.
func New() *ProductRepository {
	r := &ProductRepository{
		database: make([]model.Product, 0),
	}
	return r
}
