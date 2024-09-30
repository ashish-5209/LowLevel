package main

import (
	"fmt"
	"time"
)

type Payment struct {
	paymentId     int
	amount        float64
	paymentStatus string // e.g., "success", "failed"
	paymentTime   time.Time
}

func (p *Payment) ProcessPayment() {
	p.paymentStatus = "success"
	fmt.Println("Payment processed successfully.")
}

func (p *Payment) RefundPayment() {
	p.paymentStatus = "refunded"
	fmt.Println("Payment has been refunded.")
}
