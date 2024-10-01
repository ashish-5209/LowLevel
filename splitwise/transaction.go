package main

import "time"

// Transaction represents a financial transaction between users
type Transaction struct {
	ID       int
	FromUser User
	ToUser   User
	Amount   float64
	Status   string // Pending, Completed
	Date     time.Time
}

// Helper function to minimize transactions
func minimizeTransactions(balanceMap map[int]float64, users []User) []Transaction {
	positiveDebts := []Debt{}
	negativeDebts := []Debt{}

	// Separate users into those who owe money and those who are owed money
	for userID, balance := range balanceMap {
		user := getUserByID(users, userID)
		if balance > 0 {
			positiveDebts = append(positiveDebts, Debt{User: user, Amount: balance})
		} else if balance < 0 {
			negativeDebts = append(negativeDebts, Debt{User: user, Amount: -balance})
		}
	}

	transactions := []Transaction{}
	transactionID := 1

	// Greedily match users who are owed money with users who owe money
	for len(positiveDebts) > 0 && len(negativeDebts) > 0 {
		creditor := positiveDebts[0]
		debtor := negativeDebts[0]

		amount := min(creditor.Amount, debtor.Amount)

		// Create a new transaction
		transaction := Transaction{
			ID:       transactionID,
			FromUser: debtor.User,
			ToUser:   creditor.User,
			Amount:   amount,
			Status:   "Completed",
			Date:     time.Now(),
		}

		transactions = append(transactions, transaction)
		transactionID++

		// Update the debts
		creditor.Amount -= amount
		debtor.Amount -= amount

		// Remove settled users from their respective lists
		if creditor.Amount == 0 {
			positiveDebts = positiveDebts[1:]
		}
		if debtor.Amount == 0 {
			negativeDebts = negativeDebts[1:]
		}
	}

	return transactions
}
