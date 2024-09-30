package main

// Item represents an item in the vending machine
type Item struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (i *Item) updateQuantity(quantity int) {
	i.Quantity += quantity
}
