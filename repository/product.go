package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ol-ilyassov/repository-pattern/model"
)

// const errors set.
var (
	ErrNotFound = errors.New("product not found")
)

// Count is a function to return a number of records in database.
func (r *ProductRepository) Count() int {
	return len(r.database)
}

// GetProducts is a function to return all products.
func (r *ProductRepository) GetProducts() []model.Product {
	return r.database
}

// GetProductByID is a function to return a product by id value from database.
func (r *ProductRepository) GetProductByID(id string) (int, model.Product, error) {
	var product model.Product

	for i, v := range r.database {
		if v.ID == id {
			return i, v, nil
		}
	}

	return 0, product, ErrNotFound
}

// AddProduct is a function to add a product into database.
func (r *ProductRepository) AddProduct(input model.ProductNew) (int, model.Product) {
	product := model.Product{
		ID:          uuid.NewString(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}
	r.database = append(r.database, product)
	index := r.Count() - 1

	return index, product
}

// UpdateProduct is a function to edit a product fields in database.
func (r *ProductRepository) UpdateProduct(id string, input model.ProductNew) (int, model.Product, error) {
	index, product, err := r.GetProductByID(id)
	if err != nil {
		return 0, product, err
	}

	product = model.Product{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}
	r.database[index] = product

	return index, product, nil
}

// DeleteProduct is a function to delete a product from database.
func (r *ProductRepository) DeleteProduct(id string) (int, error) {
	index, _, err := r.GetProductByID(id)
	if err != nil {
		return 0, err
	}

	copy(r.database[index:], r.database[index+1:])
	r.database = r.database[:r.Count()-1]

	return index, nil
}
