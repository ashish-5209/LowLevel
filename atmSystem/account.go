package atmSystem

// Account represents a bank account.
type Account struct {
	CardNo  string
	Pin     string
	Balance float64
}

// NewAccount creates a new bank account.
func NewAccount(cardNo, pin string, balance float64) *Account {
	return &Account{
		CardNo:  cardNo,
		Pin:     pin,
		Balance: balance,
	}
}

// VerifyPin verifies the PIN for the account.
func (a *Account) VerifyPin(pin string) bool {
	return a.Pin == pin
}

// GetBalance returns the balance for the account.
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// Withdraw handles the withdrawal from the account.
func (a *Account) Withdraw(amount float64) bool {
	if a.Balance < amount {
		return false
	}
	a.Balance -= amount
	return true
}
