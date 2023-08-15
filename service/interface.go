package service

import "github.com/ol-ilyassov/repository-pattern/model"

// Service is an interface that holds key methods of working with ProductRepository.
type Service interface {
	Count() int
	All() []model.Product
	GetByID(id string) (model.Product, error)
	Add(input model.ProductNew) model.Product
	Update(id string, input model.ProductNew) (model.Product, error)
	Delete(id string) error
}
