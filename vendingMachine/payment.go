package main

// Payment represents a payment interface
type Payment interface {
	acceptPayment(amount float64) bool
	processPayment() bool
}
