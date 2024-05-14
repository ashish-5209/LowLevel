package atmSystem

func App() {
	bank := NewBank()

	// Create bank accounts
	account1 := NewAccount("1234", "1234", 1000)
	account2 := NewAccount("5678", "5678", 500)
	bank.Accounts["1234"] = account1
	bank.Accounts["5678"] = account2

	// Create ATM
	atm := &ATM{}

	// Insert card
	atm.InsertCard()

	// Enter PIN
	atm.EnterPin("1234")

	// Withdraw
	atm.Withdraw(200, bank, "1234")

	// Eject card
	atm.EjectCard()
}
