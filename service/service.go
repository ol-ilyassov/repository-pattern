package service

import "github.com/ol-ilyassov/repository-pattern/repository"

type Product struct {
	storage repository.Repository
}

func New() *Product {
	storage := &Product{
		storage: repository.New(),
	}
	return storage
}
