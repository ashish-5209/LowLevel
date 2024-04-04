package parkingSpot

import "fmt"

// Vehicle represents a generic vehicle
type Vehicle struct {
	LicensePlate string
	Size         int
}

// ParkVehicle parks the vehicle in the parking lot
func (pl *ParkingLot) ParkVehicle(vehicle Vehicle, spotType ParkingSpotType) bool {
	pl.Lock.Lock()
	defer pl.Lock.Unlock()

	if pl.AvailableSpots == 0 {
		fmt.Println("Parking lot is full")
		return false
	}

	for _, floor := range pl.Floors {
		for _, spot := range floor.Spots {
			spot.Lock.Lock()
			defer spot.Lock.Unlock()
			if !spot.IsOccupied {
				spot.IsOccupied = true
				spot.Vehicle = &vehicle
				pl.AvailableSpots--
				fmt.Printf("Vehicle with license plate %s parked in spot %d on floor %d\n", vehicle.LicensePlate, spot.SpotNumber, floor.Level)

				// Generate bill
				bill := pl.generateBill(vehicle, spotType)
				pl.BillingRecords[bill.BillID] = bill
				return true
			}
		}
	}

	fmt.Println("Parking lot is full")
	return false
}

// RemoveVehicle removes the vehicle from the parking lot
func (pl *ParkingLot) RemoveVehicle(vehicle Vehicle) bool {
	pl.Lock.Lock()
	defer pl.Lock.Unlock()

	for _, floor := range pl.Floors {
		for _, spot := range floor.Spots {
			spot.Lock.Lock()
			defer spot.Lock.Unlock()
			if spot.IsOccupied && spot.Vehicle.LicensePlate == vehicle.LicensePlate {
				spot.IsOccupied = false
				spot.Vehicle = nil
				pl.AvailableSpots++
				fmt.Printf("Vehicle with license plate %s removed from spot on floor %d\n", vehicle.LicensePlate, floor.Level)
				return true
			}
		}
	}

	fmt.Println("Vehicle not found in parking lot")
	return false
}
