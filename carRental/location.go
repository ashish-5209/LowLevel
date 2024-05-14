package carRental

import (
	"fmt"
	"time"
)

func (c *CarRentalSystem) Pickup(vehicleID int) error {
	for _, vehicle := range c.Vehicles {
		if vehicle.ID == vehicleID && vehicle.Reserved && vehicle.ReservedTill.After(time.Now()) {
			vehicle.Location = "In Transit"
			return nil
		}
	}
	return fmt.Errorf("vehicle with ID %d is not available for pickup", vehicleID)
}

func (c *CarRentalSystem) Drop(vehicleID int, location string) error {
	for _, vehicle := range c.Vehicles {
		if vehicle.ID == vehicleID && vehicle.Location == "In Transit" {
			vehicle.Location = location
			vehicle.Reserved = false
			vehicle.ReservedBy = ""
			vehicle.ReservedAt = time.Time{}
			vehicle.ReservedTill = time.Time{}
			return nil
		}
	}
	return fmt.Errorf("vehicle with ID %d is not available for drop at location %s", vehicleID, location)
}
