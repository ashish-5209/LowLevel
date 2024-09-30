package main

import (
	"fmt"
	"sync"
)

type Seat struct {
	seatId     int
	seatNumber string
	isBooked   bool
	price      float64
	mu         sync.Mutex // Mutex to prevent race conditions
}

func (s *Seat) BookSeat() {
	s.mu.Lock()         // Lock the mutex before modifying the seat
	defer s.mu.Unlock() // Ensure the mutex is unlocked after the function finishes

	if s.isBooked {
		fmt.Printf("Seat %s is already booked.\n", s.seatNumber)
		return
	}

	s.isBooked = true
	fmt.Printf("Seat %s has been successfully booked.\n", s.seatNumber)
}

func (s *Seat) CancelBooking() {
	s.mu.Lock()         // Lock the mutex before modifying the seat
	defer s.mu.Unlock() // Ensure the mutex is unlocked after the function finishes

	if !s.isBooked {
		fmt.Printf("Seat %s is not booked yet.\n", s.seatNumber)
		return
	}

	s.isBooked = false
	fmt.Printf("Booking for seat %s has been successfully cancelled.\n", s.seatNumber)

}
