package main

import "fmt"

type Screen struct {
	screenId   int
	screenName string
	totalSeats int
}

func (s *Screen) GetSeatAvailability() {
	fmt.Println("Getting seat availability for screen", s.screenName)
}
