package main

import (
	"fmt"
	"time"
)

// Transaction represents a transaction in the vending machine
type Transaction struct {
	transactionID int
	dateTime      time.Time
	item          *Item
	amountPaid    float64
	paymentStatus string
}

func (t *Transaction) createReceipt() {
	fmt.Printf("Receipt: Transaction ID: %d, Item: %s, Amount Paid: %.2f\n", t.transactionID, t.item.Name, t.amountPaid)
}

func (t *Transaction) recordTransaction() {
	fmt.Println("Transaction recorded successfully.")
}
