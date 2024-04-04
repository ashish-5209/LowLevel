package parkingSpot

import "fmt"

// Bill represents a bill generated for a parked vehicle
type Bill struct {
	BillID  string
	Vehicle *Vehicle
	Amount  float64
}

// generateBill generates a bill for the parked vehicle
func (pl *ParkingLot) generateBill(vehicle Vehicle, spotType ParkingSpotType) Bill {
	// Calculate amount based on hourly rate and duration parked (for simplicity, assuming a fixed hourly rate)
	amount := spotType.HourlyRate * 2 // Assuming 2 hours parked

	bill := Bill{
		BillID:  fmt.Sprintf("BILL_%s", vehicle.LicensePlate),
		Vehicle: &vehicle,
		Amount:  amount,
	}

	return bill
}
