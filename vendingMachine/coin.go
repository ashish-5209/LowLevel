package main

import "fmt"

// Coin represents a coin payment
type Coin struct {
	value float64
}

func (c *Coin) acceptPayment(amount float64) bool {
	if c.value >= amount {
		return true
	}
	return false
}

func (c *Coin) processPayment() bool {
	// Processing coin payment
	fmt.Println("Payment processed using coins.")
	return true
}
