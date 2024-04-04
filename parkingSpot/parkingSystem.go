// _________________________            ___________________
// |         ParkingLot      |          |      ParkingSpot  |
// |-------------------------|          |------------------|
// | - capacity: int         |          | - spotNumber: int|
// | - availableSpots: int   |          | - isOccupied: bool|
// | - floors: List<Floor>   |          | - type: SpotType |
// | - lock: Lock            |          | - vehicle: Vehicle |
// |-------------------------|          |__________________|
// | + ParkingLot(capacity: int)|
// | + parkVehicle(vehicle: Vehicle): bool|
// | + removeVehicle(vehicle: Vehicle): bool|
// | + getAvailableSpots(): int|
// | + getOccupiedSpots(): int|
// | + calculateFee(vehicle: Vehicle): float64|
// | + makePayment(vehicle: Vehicle, amount: float64): bool|
// | + generateBill(vehicle: Vehicle): Bill|
// |_________________________|

//          __________________
//         |      Vehicle     |
//         |------------------|
//         | - licensePlate: string|
//         | - size: int          |
//         |____________________|

//                ^
//                |
//                |
//          ______|______
//         |             |
//         |             |
//  _______|______       |    ___________________
// |   Car        |      |   |     Bus           |
// |-------------|      |   |-------------------|
// |             |      |   |                   |
// |             |      |   |                   |
// |_____________|      |   |___________________|

//                ^
//                |
//                |
//          ______|______
//         |             |
//         |             |
//  _______|______       |
// |  Motorbike  |      |
// |-------------|      |
// |             |      |
// |             |      |
// |_____________|      |

//  ___________________
// |     SpotType     |
// |-------------------|
// | - name: string    |
// | - hourlyRate: float64|
// |___________________|

//  ___________________
// |       Bill        |
// |-------------------|
// | - billID: string  |
// | - vehicle: Vehicle|
// | - amount: float64 |
// |___________________|

//  ___________________
// |     Payment       |
// |-------------------|
// | - paymentID: string|
// | - billID: string  |
// | - amount: float64 |
// | - isPaid: bool    |
// |___________________|

//  ___________________
// |       Floor       |
// |-------------------|
// | - level: int      |
// | - spots: List<ParkingSpot> |
// |___________________|

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
