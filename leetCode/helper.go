package main

import (
	"fmt"
	"math/rand"
)

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func hashPassword(password string) string {
	return fmt.Sprintf("%x", password)
}
