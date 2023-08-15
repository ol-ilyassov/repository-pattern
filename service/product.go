package service

import (
	"github.com/ol-ilyassov/repository-pattern/model"
)

// Count() returns a number of products in repository.
func (s *Product) Count() int {
	return s.storage.Count()
}

// GetProducts returns all products from repository.
func (s *Product) All() []model.Product {
	return s.storage.GetProducts()
}

// GetProductByID returns product from repository.
func (s *Product) GetByID(id string) (model.Product, error) {
	_, product, error := s.storage.GetProductByID(id)
	return product, error
}

// AddProduct adds a new product into repository.
func (s *Product) Add(input model.ProductNew) model.Product {
	_, product := s.storage.AddProduct(input)
	return product
}

// UpdateProduct updates a product in repository and returns updated product.
func (s *Product) Update(id string, input model.ProductNew) (model.Product, error) {
	_, product, error := s.storage.UpdateProduct(id, input)
	return product, error
}

// DeleteProduct deletes a product in repository.
func (s *Product) Delete(id string) error {
	_, err := s.storage.DeleteProduct(id)
	return err
}
