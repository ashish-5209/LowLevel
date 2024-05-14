package atmSystem

// Bank represents the bank system.
type Bank struct {
	Accounts map[string]*Account
	Notes    map[float64]int
}

// NewBank creates a new bank system.
func NewBank() *Bank {
	return &Bank{
		Accounts: make(map[string]*Account),
		Notes: map[float64]int{
			100.0: 10,
			50.0:  10,
			20.0:  10,
			10.0:  10,
		},
	}
}

// VerifyPin verifies the PIN for a given card number.
func (b *Bank) VerifyPin(cardNo, pin string) bool {
	if acc, ok := b.Accounts[cardNo]; ok {
		return acc.VerifyPin(pin)
	}
	return false
}

// GetBalance returns the balance for a given card number.
func (b *Bank) GetBalance(cardNo string) float64 {
	if acc, ok := b.Accounts[cardNo]; ok {
		return acc.GetBalance()
	}
	return 0
}

// Withdraw handles the withdrawal from the bank.
func (b *Bank) Withdraw(cardNo string, amount float64) []float64 {
	if acc, ok := b.Accounts[cardNo]; ok {
		if acc.GetBalance() < amount {
			return nil
		}
		if b.CheckNotes(amount) {
			if acc.Withdraw(amount) {
				return b.DispenseNotes(amount)
			}
		}
	}
	return nil
}

// CheckNotes checks if the bank has enough notes to fulfill the withdrawal.
func (b *Bank) CheckNotes(amount float64) bool {
	total := amount
	for note, count := range b.Notes {
		for total >= note && count > 0 {
			total -= note
			count--
		}
	}
	return total == 0
}

// DispenseNotes dispenses the notes for the withdrawal.
func (b *Bank) DispenseNotes(amount float64) []float64 {
	notes := make([]float64, 0)
	for note, count := range b.Notes {
		for amount >= note && count > 0 {
			notes = append(notes, note)
			amount -= note
			count--
		}
	}
	return notes
}

// UpdateNotes updates the notes count after a withdrawal.
func (b *Bank) UpdateNotes(amount float64) bool {
	for note, count := range b.Notes {
		for amount >= note && count > 0 {
			amount -= note
			count--
		}
	}
	return amount == 0
}
