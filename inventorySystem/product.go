package main

import "fmt"

// Product struct representing a product in the inventory
type Product struct {
	ProductID   int
	Name        string
	Description string
	Price       float64
	Quantity    int
	SupplierID  int
}

// Methods for Product
func (p *Product) AddProduct() {
	fmt.Printf("Product %s added successfully.\n", p.Name)
}

func (p *Product) UpdateProduct(name string, description string, price float64, quantity int) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.Quantity = quantity
	fmt.Printf("Product %d updated successfully.\n", p.ProductID)
}

func (p *Product) DeleteProduct() {
	fmt.Printf("Product %d deleted successfully.\n", p.ProductID)
}

func (p *Product) GetProductDetails() {
	fmt.Printf("Product ID: %d, Name: %s, Price: %.2f, Quantity: %d\n", p.ProductID, p.Name, p.Price, p.Quantity)
}
