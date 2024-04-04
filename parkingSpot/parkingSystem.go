package parkingSpot

import "fmt"

func ParkingSystem() {
	// Initialize parking lot with a capacity of 10 and 3 floors
	parkingLot := NewParkingLot(10, 3)

	// Park a vehicle
	vehicle1 := Vehicle{LicensePlate: "ABC123", Size: 1}
	spotType := ParkingSpotType{Name: "Regular", HourlyRate: 5.0}
	parkingLot.ParkVehicle(vehicle1, spotType)

	// Generate bill for the parked vehicle
	fmt.Println("Billing Records:")
	for _, bill := range parkingLot.BillingRecords {
		fmt.Printf("Bill ID: %s, Amount: %.2f\n", bill.BillID, bill.Amount)
	}

	// Make a payment for the bill
	parkingLot.makePayment(vehicle1, "BILL_ABC123", 10.0)

	// Remove the vehicle
	parkingLot.RemoveVehicle(vehicle1)
}
