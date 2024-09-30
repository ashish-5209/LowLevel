package main

import "fmt"

// Inventory struct representing the stock of products
type Inventory struct {
	InventoryID int
	ProductID   int
	Quantity    int
	Location    string
}

// Methods for Inventory
func (i *Inventory) AddInventory() {
	fmt.Printf("Inventory for Product ID %d added successfully.\n", i.ProductID)
}

func (i *Inventory) UpdateInventory(quantity int, location string) {
	i.Quantity = quantity
	i.Location = location
	fmt.Printf("Inventory %d updated successfully.\n", i.InventoryID)
}

func (i *Inventory) DeleteInventory() {
	fmt.Printf("Inventory %d deleted successfully.\n", i.InventoryID)
}

func (i *Inventory) GetInventoryDetails() {
	fmt.Printf("Inventory ID: %d, Product ID: %d, Quantity: %d, Location: %s\n", i.InventoryID, i.ProductID, i.Quantity, i.Location)
}
