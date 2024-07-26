// CarRentalSystem
// +--------------------------------------------------+
// |               CarRentalSystem                    |
// +--------------------------------------------------+
// | - Locations: []Location                         |
// | - Users: []User                                 |
// | - Bills: []Bill                                 |
// +--------------------------------------------------+
// | + SearchVehicle(location Location, startDate time.Time, endDate time.Time) []Vehicle |
// | + MakeReservation(user User, vehicle Vehicle, store Store, location Location, startDate time.Time, endDate time.Time) Reservation |
// | + GenerateBill(reservation Reservation) Bill   |
// | + MakePayment(bill Bill, paymentMethod string) Payment |
// +--------------------------------------------------+

// User
// +--------------------------------------------------+
// |                    User                           |
// +--------------------------------------------------+
// | - UserID: int                                    |
// | - Name: string                                  |
// | - Email: string                                 |
// | - PhoneNumber: string                           |
// | - LicenseNumber: string                         |
// +--------------------------------------------------+
// | + NewUser(id int, name string, email string, phone string, license string) |
// +--------------------------------------------------+

// Location
// +--------------------------------------------------+
// |                  Location                        |
// +--------------------------------------------------+
// | - LocationID: int                               |
// | - Name: string                                  |
// | - Stores: []Store                               |
// +--------------------------------------------------+
// | + NewLocation(id int, name string)              |
// | + AddStore(store Store)                         |
// +--------------------------------------------------+

// Store
// +--------------------------------------------------+
// |                    Store                         |
// +--------------------------------------------------+
// | - StoreID: int                                  |
// | - Name: string                                 |
// | - Address: string                              |
// | - Vehicles: []Vehicle                           |
// +--------------------------------------------------+
// | + NewStore(id int, name string, address string) |
// | + AddVehicle(vehicle Vehicle)                   |
// +--------------------------------------------------+

// Vehicle (Interface)
// +--------------------------------------------------+
// |                    Vehicle (interface)           |
// +--------------------------------------------------+
// | + GetVehicleType() string                       |
// | + GetPricePerDay() float64                      |
// | + GetID() int                                   |
// | + IsAvailable() bool                            |
// | + SetStatus(status string)                      |
// +--------------------------------------------------+

// BaseVehicle
// +--------------------------------------------------+
// |                  BaseVehicle                     |
// +--------------------------------------------------+
// | - VehicleID: int                                |
// | - Model: string                                |
// | - Make: string                                 |
// | - Year: int                                    |
// | - PricePerDay: float64                         |
// | - Status: string // Available, Reserved, Out of Service |
// +--------------------------------------------------+
// | + GetID() int                                  |
// | + IsAvailable() bool                           |
// | + SetStatus(status string)                     |
// +--------------------------------------------------+

// Car
// +--------------------------------------------------+
// |                    Car                           |
// +--------------------------------------------------+
// | - NumberOfDoors: int                            |
// +--------------------------------------------------+
// | + GetVehicleType() string                       |
// | + GetPricePerDay() float64                      |
// +--------------------------------------------------+

// Bike
// +--------------------------------------------------+
// |                    Bike                          |
// +--------------------------------------------------+
// | - HasGear: bool                                 |
// +--------------------------------------------------+
// | + GetVehicleType() string                       |
// | + GetPricePerDay() float64                      |
// +--------------------------------------------------+

// Reservation
// +--------------------------------------------------+
// |                 Reservation                      |
// +--------------------------------------------------+
// | - ReservationID: int                            |
// | - UserID: int                                   |
// | - VehicleID: int                                |
// | - StoreID: int                                  |
// | - LocationID: int                               |
// | - StartDate: time.Time                          |
// | - EndDate: time.Time                            |
// | - TotalAmount: float64                          |
// | - Status: string // Reserved, Completed, Cancelled |
// +--------------------------------------------------+
// | + NewReservation(id int, userID int, vehicleID int, storeID int, locationID int, startDate time.Time, endDate time.Time) |
// +--------------------------------------------------+

// Bill
// +--------------------------------------------------+
// |                    Bill                          |
// +--------------------------------------------------+
// | - BillID: int                                   |
// | - ReservationID: int                           |
// | - UserID: int                                  |
// | - VehicleID: int                               |
// | - TotalAmount: float64                         |
// | - DiscountAmount: float64                      |
// | - TaxAmount: float64                           |
// | - FinalAmount: float64                         |
// | - PaymentStatus: string // Paid, Unpaid         |
// +--------------------------------------------------+
// | + NewBill(id int, reservationID int, amount float64, discountAmount float64, taxAmount float64, finalAmount float64) |
// +--------------------------------------------------+

// Payment
// +--------------------------------------------------+
// |                    Payment                       |
// +--------------------------------------------------+
// | - PaymentID: int                                |
// | - BillID: int                                  |
// | - UserID: int                                  |
// | - Amount: float64                              |
// | - PaymentMethod: string // CreditCard, DebitCard, NetBanking |
// | - PaymentStatus: string // Success, Failed     |
// | - PaymentDate: time.Time                        |
// +--------------------------------------------------+
// | + NewPayment(id int, billID int, userID int, amount float64, paymentMethod string) |
// +--------------------------------------------------+

package carRental

import (
	"fmt"
	"time"
)

func App() {
	// Example usage
	user := User{
		UserID:        1,
		Name:          "John Doe",
		Email:         "john.doe@example.com",
		PhoneNumber:   "123-456-7890",
		LicenseNumber: "XYZ12345",
	}

	location := Location{
		LocationID: 1,
		Name:       "City Center",
		Stores: []Store{
			{
				StoreID: 1,
				Name:    "Main Store",
				Address: "123 Main St",
				Vehicles: []Vehicle{
					&Car{
						BaseVehicle: BaseVehicle{
							VehicleID:   1,
							Model:       "Sedan",
							Make:        "Toyota",
							Year:        2020,
							PricePerDay: 50.0,
							Status:      "Available",
						},
						NumberOfDoors: 4,
					},
					&Bike{
						BaseVehicle: BaseVehicle{
							VehicleID:   2,
							Model:       "Sport",
							Make:        "Yamaha",
							Year:        2021,
							PricePerDay: 20.0,
							Status:      "Available",
						},
						HasGear: true,
					},
				},
			},
		},
	}

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 3) // 3 days rental

	// Search for available vehicles
	availableVehicles := SearchVehicle(location, startDate, endDate)
	fmt.Printf("Available vehicles:\n")
	for _, vehicle := range availableVehicles {
		fmt.Println(vehicle)
	}

	if len(availableVehicles) > 0 {
		// Make a reservation for the first available vehicle
		reservation := MakeReservation(user, availableVehicles[0], location.Stores[0], location, startDate, endDate)
		fmt.Printf("Reservation made: %+v\n", reservation)

		// Generate bill for the reservation
		bill := GenerateBill(reservation)
		fmt.Printf("Generated bill: %+v\n", bill)

		// Make payment for the bill
		payment := MakePayment(bill, "CreditCard")
		fmt.Printf("Payment made: %+v\n", payment)
	}
}
