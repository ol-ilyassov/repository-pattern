package main

import (
	"github.com/ol-ilyassov/repository-pattern/client"
	"github.com/ol-ilyassov/repository-pattern/service"
)

func main() {
	productService := service.New()
	client.DisplayMenu(productService)
}
