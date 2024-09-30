package main

import "fmt"

// Card represents a card payment
type Card struct {
	cardNumber string
	expiryDate string
}

func (card *Card) acceptPayment(amount float64) bool {
	// For simplicity, assume the card is valid and has sufficient balance
	fmt.Println("Card accepted.")
	return true
}

func (card *Card) processPayment() bool {
	// Processing card payment
	fmt.Println("Payment processed using card.")
	return true
}
