package atmSystem

import "fmt"

// ATM represents the ATM machine.
type ATM struct {
	Balance      float64
	CardInserted bool
	PinEntered   bool
}

// InsertCard simulates inserting a card into the ATM.
func (a *ATM) InsertCard() {
	a.CardInserted = true
	fmt.Println("Card inserted")
}

// EnterPin simulates entering a PIN into the ATM.
func (a *ATM) EnterPin(pin string) {
	a.PinEntered = true
	fmt.Println("PIN entered")
}

// Withdraw simulates a withdrawal from the ATM.
func (a *ATM) Withdraw(amount float64, bank *Bank, cardNo string) {
	if !a.CardInserted {
		fmt.Println("Please insert your card")
		return
	}
	if !a.PinEntered {
		fmt.Println("Please enter your PIN")
		return
	}
	if amount > a.Balance {
		fmt.Println("Insufficient balance")
		return
	}
	notes := bank.Withdraw(cardNo, amount)
	if len(notes) == 0 {
		fmt.Println("Unable to dispense amount")
		return
	}
	for _, note := range notes {
		fmt.Printf("Dispensing %.2f\n", note)
	}
	a.Balance -= amount
	fmt.Printf("Remaining balance: %.2f\n", a.Balance)
}

// CheckBalance returns the current balance in the ATM.
func (a *ATM) CheckBalance() float64 {
	return a.Balance
}

// EjectCard simulates ejecting the card from the ATM.
func (a *ATM) EjectCard() {
	a.CardInserted = false
	a.PinEntered = false
	fmt.Println("Card ejected")
}
