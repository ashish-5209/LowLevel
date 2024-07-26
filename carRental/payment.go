package carRental

import "time"

// Payment handles the payment process.
type Payment struct {
	PaymentID     int
	BillID        int
	UserID        int
	Amount        float64
	PaymentMethod string // CreditCard, DebitCard, NetBanking, etc.
	PaymentStatus string // Success, Failed
	PaymentDate   time.Time
}

// MakePayment processes the payment for the bill.
func MakePayment(bill Bill, paymentMethod string) Payment {
	payment := Payment{
		PaymentID:     1, // This would normally be generated by the system
		BillID:        bill.BillID,
		UserID:        bill.UserID,
		Amount:        bill.FinalAmount,
		PaymentMethod: paymentMethod,
		PaymentStatus: "Success", // Assume success for this example
		PaymentDate:   time.Now(),
	}

	// Update the bill status
	bill.PaymentStatus = "Paid"

	return payment
}
