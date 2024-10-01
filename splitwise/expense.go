package main

// Expense represents an expense made by a user
type Expense struct {
	ID          int
	Amount      float64
	Description string
	PaidBy      User
	SplitType   string // e.g., Equal, Unequal, Percentage
	Group       Group
}

// Helper function to calculate net balances for each user in the group
func calculateNetBalances(expenses []Expense, users []User) map[int]float64 {
	balanceMap := make(map[int]float64)

	// Initialize the balance for each user
	for _, user := range users {
		balanceMap[user.ID] = 0
	}

	// Calculate net balances based on expenses
	for _, expense := range expenses {
		splitAmount := expense.Amount / float64(len(expense.Group.Users))

		// Distribute expense among all users
		for _, user := range expense.Group.Users {
			if user.ID == expense.PaidBy.ID {
				balanceMap[user.ID] += expense.Amount - splitAmount
			} else {
				balanceMap[user.ID] -= splitAmount
			}
		}
	}

	return balanceMap
}
