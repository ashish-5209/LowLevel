package main

// User represents a user in the system
type User struct {
	ID    int
	Name  string
	Email string
	Phone string
}

// Helper function to find a user by ID
func getUserByID(users []User, id int) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return User{}
}

// Helper function to get the minimum of two values
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
