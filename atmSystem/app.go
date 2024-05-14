// +--------------------------------------+
// |                ATM                   |
// +--------------------------------------+
// | - Balance: float64                  |
// | - CardInserted: bool                |
// | - PinEntered: bool                  |
// +--------------------------------------+
// | + InsertCard()                      |
// | + EnterPin(pin string)              |
// | + Withdraw(amount float64)          |
// | + CheckBalance() -> float64         |
// | + EjectCard()                       |
// +--------------------------------------+

// +--------------------------------------+
// |              Bank                    |
// +--------------------------------------+
// | - Accounts: map[string]*Account     |
// | - Notes: map[float64]int            |
// +--------------------------------------+
// | + VerifyPin(cardNo string, pin string) -> bool |
// | + GetBalance(cardNo string) -> float64 |
// | + Withdraw(cardNo string, amount float64) -> bool |
// | + GetNotes() map[float64]int                   |
// | + UpdateNotes(amount float64) bool            |
// +--------------------------------------+

// +--------------------------------------+
// |               Account                |
// +--------------------------------------+
// | - CardNo: string                    |
// | - Pin: string                       |
// | - Balance: float64                  |
// +--------------------------------------+
// | + VerifyPin(pin string) -> bool     |
// | + GetBalance() -> float64           |
// | + Withdraw(amount float64) -> bool |
// +--------------------------------------+

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
