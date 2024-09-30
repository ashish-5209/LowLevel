package main

import (
	"fmt"
	"time"
)

type Show struct {
	showId    int
	startTime time.Time
	endTime   time.Time
	date      time.Time
}

func (s *Show) GetAvailableSeats() {
	fmt.Println("Fetching available seats for the show...")
}

func (s *Show) ReserveSeats() {
	fmt.Println("Reserving seats for the show...")
}
