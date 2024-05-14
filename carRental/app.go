// +--------------------------------------------------+
// |               CarRentalSystem                    |
// +--------------------------------------------------+
// | - Vehicles: []*Vehicle                          |
// | - Bills: []*Bill                                |
// +--------------------------------------------------+
// | + NewCarRentalSystem()                          |
// | + AddVehicle(id int, vehicleType VehicleType,   |
// |               location string)                   |
// | + ReserveVehicle(vehicleID int, customerID      |
// |               string, pickupTime, dropTime      |
// |               time.Time)                        |
// | + CancelReservation(vehicleID int)              |
// | + Pickup(vehicleID int)                         |
// | + Drop(vehicleID int, location string)          |
// | + GenerateBill(vehicleID int, amount float64,  |
// |               paymentType PaymentType)          |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                    Vehicle                       |
// +--------------------------------------------------+
// | - ID: int                                        |
// | - Type: VehicleType                             |
// | - Status: VehicleStatus                         |
// | - Location: string                              |
// | - Reserved: bool                                |
// | - ReservedBy: string                            |
// | - ReservedAt: time.Time                         |
// | - ReservedTill: time.Time                       |
// +--------------------------------------------------+
// | + NewVehicle(id int, vehicleType VehicleType,   |
// |               location string)                   |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                     Bill                         |
// +--------------------------------------------------+
// | - ID: int                                        |
// | - Amount: float64                               |
// | - PaymentType: PaymentType                      |
// +--------------------------------------------------+
// | + NewBill(id int, amount float64,              |
// |               paymentType PaymentType)          |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                Enumerations                      |
// +--------------------------------------------------+
// | - VehicleType                                   |
// | - VehicleStatus                                 |
// | - PaymentType                                   |
// +--------------------------------------------------+

package carRental

import (
	"fmt"
	"time"
)

func App() {
	fmt.Println("Statrt")
	carRental := NewCarRentalSystem()

	// Adding vehicles
	carRental.AddVehicle(1, Car, "Location A")
	carRental.AddVehicle(2, Car, "Location B")
	carRental.AddVehicle(3, Bike, "Location C")

	// Reserving a vehicle
	err := carRental.ReserveVehicle(1, "Customer1", time.Now(), time.Now().Add(time.Hour*24))
	if err != nil {
		fmt.Println("Error reserving vehicle:", err)
	}

	// Pickup
	err = carRental.Pickup(1)
	if err != nil {
		fmt.Println("Error picking up vehicle:", err)
	}

	// Dropping the vehicle
	err = carRental.Drop(1, "Location D")
	if err != nil {
		fmt.Println("Error dropping vehicle:", err)
	}

	// Generating bill
	err = carRental.GenerateBill(1, 50.0, CreditCard)
	if err != nil {
		fmt.Println("Error generating bill:", err)
	}
	fmt.Println(*carRental.Bills[0])
}
