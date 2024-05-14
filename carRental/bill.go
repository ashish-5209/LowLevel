package carRental

type Bill struct {
	ID          int
	Amount      float64
	PaymentType PaymentType
}

func (c *CarRentalSystem) GenerateBill(vehicleID int, amount float64, paymentType PaymentType) error {
	bill := &Bill{
		ID:          len(c.Bills) + 1,
		Amount:      amount,
		PaymentType: paymentType,
	}
	c.Bills = append(c.Bills, bill)
	return nil
}
