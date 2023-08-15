package client

import (
	"fmt"

	"github.com/ol-ilyassov/repository-pattern/model"
	"github.com/ol-ilyassov/repository-pattern/service"
)

// DisplayMenu is a function to demonstrate interactive functionality of application.
func DisplayMenu(service service.Service) {
	var choice int
	fmt.Println("> Products Manager App")
	fmt.Println("> Please, choose an action:")
	for {
		fmt.Println("========================")
		fmt.Println("[1] Get all products")
		fmt.Println("[2] Get Product by ID")
		fmt.Println("[3] Add new product")
		fmt.Println("[4] Update product")
		fmt.Println("[5] Delete product")
		fmt.Println("[6] Quit")
		fmt.Println("========================")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			getAllProducts(service)
		case 2:
			getProductById(service)
		case 3:
			addProduct(service)
		case 4:
			updateProduct(service)
		case 5:
			deleteProduct(service)
		case 6:
			fmt.Println("========================")
			fmt.Println("Application Quit")
			fmt.Println("========================")
			return
		default:
			fmt.Println("========================")
			fmt.Println("Invalid input!")
			fmt.Println("Please, choose an action from 1 to 6")
		}
	}
}

// getAllProducts is a function to demonstrate all products.
func getAllProducts(service service.Service) {
	var products []model.Product = service.All()

	if service.Count() == 0 {
		fmt.Println("========================")
		fmt.Println("No Products!")
	} else {
		fmt.Println("========================")
		fmt.Println("Products list:")
		for _, product := range products {
			showProductData(product)
		}
	}
}

// getProductById is a function to demonstrate a product by ID.
func getProductById(service service.Service) {
	var productId string
	fmt.Println("========================")
	fmt.Println("Enter product ID: ")
	fmt.Scanln(&productId)

	product, err := service.GetByID(productId)
	if err != nil {
		fmt.Println("========================")
		fmt.Println("Error:", err)
	} else {
		fmt.Println("========================")
		fmt.Println("Product data:")
		showProductData(product)
	}
}

// addProduct is a function to add a new product.
func addProduct(service service.Service) {
	var productData model.ProductNew

	fmt.Println("========================")
	fmt.Println("Enter product name: ")
	fmt.Scanln(&productData.Name)

	fmt.Println("Enter product description: ")
	fmt.Scanln(&productData.Description)

	fmt.Println("Enter product price: ")
	fmt.Scanln(&productData.Price)

	fmt.Println("Enter product stock: ")
	fmt.Scanln(&productData.Stock)

	insertedProduct := service.Add(productData)

	fmt.Println("========================")
	fmt.Println("New product:")
	showProductData(insertedProduct)
}

// updateProduct is a function to update a product data.
func updateProduct(service service.Service) {
	var productId string
	var productData model.ProductNew

	fmt.Println("========================")
	fmt.Println("Enter product id: ")
	fmt.Scanf("%s\n", &productId)

	fmt.Println("Enter product name: ")
	fmt.Scanln(&productData.Name)

	fmt.Println("Enter product description: ")
	fmt.Scanln(&productData.Description)

	fmt.Println("Enter product price: ")
	fmt.Scanln(&productData.Price)

	fmt.Println("Enter product stock: ")
	fmt.Scanln(&productData.Stock)

	updatedProduct, err := service.Update(productId, productData)
	if err != nil {
		fmt.Println("========================")
		fmt.Println("Error:", err)
	} else {
		fmt.Println("========================")
		fmt.Println("Updated product:")
		showProductData(updatedProduct)
	}
}

// deleteProduct is a function to delete a product.
func deleteProduct(service service.Service) {
	var productId string
	fmt.Println("========================")
	fmt.Println("Enter product id: ")
	fmt.Scanf("%s\n", &productId)

	err := service.Delete(productId)
	if err != nil {
		fmt.Println("========================")
		fmt.Println("Error:", err)
	} else {
		fmt.Println("========================")
		fmt.Println("Product deleted!")
	}
}

// showProductData is a function to print a product data in console.
func showProductData(product model.Product) {
	fmt.Println("========================")
	fmt.Println("ID: ", product.ID)
	fmt.Println("Name: ", product.Name)
	fmt.Println("Description: ", product.Description)
	fmt.Println("Price: ", product.Price)
	fmt.Println("Stock: ", product.Stock)
	fmt.Println("========================")
}
