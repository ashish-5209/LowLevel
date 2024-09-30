package main

import "fmt"

// Display represents the display of the vending machine
type Display struct {
	message string
}

func (d *Display) showMessage(message string) {
	fmt.Println("Display Message:", message)
}

func (d *Display) displayPrice(item *Item) {
	fmt.Printf("The price of %s is %.2f\n", item.Name, item.Price)
}
