package carRental

import (
	"fmt"
	"time"
)

func (c *CarRentalSystem) ReserveVehicle(vehicleID int, customerID string, pickupTime, dropTime time.Time) error {
	for _, vehicle := range c.Vehicles {
		if vehicle.ID == vehicleID && vehicle.Status == Active && !vehicle.Reserved {
			vehicle.Reserved = true
			vehicle.ReservedBy = customerID
			vehicle.ReservedAt = time.Now()
			vehicle.ReservedTill = dropTime
			return nil
		}
	}
	return fmt.Errorf("vehicle with ID %d is not available for reservation", vehicleID)
}

func (c *CarRentalSystem) CancelReservation(vehicleID int) error {
	for _, vehicle := range c.Vehicles {
		if vehicle.ID == vehicleID && vehicle.Reserved && vehicle.ReservedTill.After(time.Now()) {
			vehicle.Reserved = false
			vehicle.ReservedBy = ""
			vehicle.ReservedAt = time.Time{}
			vehicle.ReservedTill = time.Time{}
			return nil
		}
	}
	return fmt.Errorf("reservation for vehicle with ID %d cannot be cancelled", vehicleID)
}
