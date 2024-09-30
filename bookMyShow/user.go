package main

import "fmt"

type User struct {
	userId      int
	name        string
	email       string
	phoneNumber string
}

func (u *User) SearchMovie() {
	fmt.Println("Searching for movies...")
}

func (u *User) BookTicket() {
	fmt.Println("Booking tikcet...")
}

func (u *User) CancelTicket() {
	fmt.Println("Cancelling ticket...")
}
