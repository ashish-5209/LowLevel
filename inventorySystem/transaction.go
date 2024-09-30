package main

import (
	"fmt"
	"time"
)

// Transaction struct representing a product transaction
type Transaction struct {
	TransactionID   int
	ProductID       int
	Quantity        int
	TransactionType string // e.g., "In", "Out"
	TransactionDate time.Time
}

// Methods for Transaction
func (t *Transaction) AddTransaction() {
	fmt.Printf("Transaction %d for Product ID %d added successfully.\n", t.TransactionID, t.ProductID)
}

func (t *Transaction) GetTransactionDetails() {
	fmt.Printf("Transaction ID: %d, Product ID: %d, Quantity: %d, Type: %s, Date: %s\n", t.TransactionID, t.ProductID, t.Quantity, t.TransactionType, t.TransactionDate)
}
