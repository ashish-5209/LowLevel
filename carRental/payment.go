package carRental

type PaymentType int

const (
	CreditCard PaymentType = iota
	DebitCard
	Cash
)
