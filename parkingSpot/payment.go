package parkingSpot

import "fmt"

// Payment represents a payment made for a bill
type Payment struct {
	PaymentID string
	BillID    string
	Amount    float64
	IsPaid    bool
}

// makePayment makes a payment for the given bill
func (pl *ParkingLot) makePayment(vehicle Vehicle, billID string, amount float64) bool {
	pl.Lock.Lock()
	defer pl.Lock.Unlock()

	// Check if bill exists
	if _, ok := pl.BillingRecords[billID]; !ok {
		fmt.Println("Bill not found")
		return false
	}

	// Mark bill as paid
	payment := Payment{
		PaymentID: fmt.Sprintf("PAY_%s", billID),
		BillID:    billID,
		Amount:    amount,
		IsPaid:    true,
	}

	// Record payment
	pl.Payments[payment.PaymentID] = payment

	fmt.Printf("Payment of %.2f made for bill %s\n", amount, billID)
	return true
}
