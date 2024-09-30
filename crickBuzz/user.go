package main

import "fmt"

// User struct representing a user of the application
type User struct {
	UserID      int
	Name        string
	Email       string
	PhoneNumber string
}

// Methods for User struct
func (u *User) Register() {
	fmt.Printf("User %s registered successfully.\n", u.Name)
}

func (u *User) Subscribe() {
	fmt.Printf("User %s subscribed for notifications.\n", u.Name)
}
