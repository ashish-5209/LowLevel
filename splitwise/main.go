// +--------------------+
// |      User          |
// |--------------------|
// | - userID: int      |
// | - name: string     |
// | - email: string    |
// | - phone: string    |
// |--------------------|
// | + addFriend(user: User)        |
// | + getBalance()                 |
// +--------------------+              |
//
//	|                                 |
//	|                                 |
//	v                                 v
//
// +--------------------+        +--------------------+
// |    Group           |        |   Expense          |
// |--------------------|        |--------------------|
// | - groupID: int     |        | - expenseID: int   |
// | - name: string     |        | - amount: float    |
// | - users: []User    |        | - description: string|
// |--------------------|        | - paidBy: User     |
// | + addUser(user: User)        | - splitType: string| // Equal, Unequal, Percentage
// | + removeUser(user: User)     | - groupID: int     |
// | + getTotalBalance()          |--------------------|
// +--------------------+        | + calculateSplit()  |
//
//	|                        | + addExpense()      |
//	|                        | + getExpenseDetails() |
//	v                        +--------------------+
//
// +--------------------+
// |   Transaction      |
// |--------------------|
// | - transactionID: int|
// | - fromUser: User    |
// | - toUser: User      |
// | - amount: float     |
// | - status: string    | // Pending, Completed
// | - date: time.Time   |
// |--------------------|
// | + addTransaction()  |
// | + settleTransaction()|
// | + getTransactionDetails() |
// +--------------------+
//
//	|
//	v
//
// +--------------------+
// |   Balance          |
// |--------------------|
// | - balanceID: int   |
// | - user: User       |
// | - amountOwed: float|
// | - amountDue: float |
// |--------------------|
// | + updateBalance()   |
// | + getBalanceDetails()|
// +--------------------+
package main

import "fmt"

func main() {
	// Example users
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	// Example group
	group := Group{
		ID:    1,
		Name:  "Trip Group",
		Users: users,
	}

	// Example expenses
	expenses := []Expense{
		{ID: 1, Amount: 300, Description: "Dinner", PaidBy: users[0], SplitType: "Equal", Group: group},
		{ID: 2, Amount: 200, Description: "Lunch", PaidBy: users[1], SplitType: "Equal", Group: group},
	}

	// Step 1: Calculate net balances
	balances := calculateNetBalances(expenses, users)

	// Step 2: Minimize transactions
	transactions := minimizeTransactions(balances, users)

	// Display the transactions
	for _, transaction := range transactions {
		fmt.Printf("%s owes %s %.2f\n", transaction.FromUser.Name, transaction.ToUser.Name, transaction.Amount)
	}
}
