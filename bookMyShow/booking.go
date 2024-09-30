package main

import (
	"fmt"
	"time"
)

type Booking struct {
	bookingId   int
	bookingTime time.Time
	totalPrice  float64
	status      string // e.g., "confirmed", "cancelled"
}

func (b *Booking) ConfirmBooking() {
	b.status = "confirmed"
	fmt.Println("Booking confirmed.")
}

func (b *Booking) CancelBooking() {
	b.status = "cancelled"
	fmt.Println("Booking cancelled.")
}
