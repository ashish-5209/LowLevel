package carRental

import "time"

type Vehicle struct {
	ID           int
	Type         VehicleType
	Status       VehicleStatus
	Location     string
	Reserved     bool
	ReservedBy   string
	ReservedAt   time.Time
	ReservedTill time.Time
}

func (c *CarRentalSystem) AddVehicle(id int, vehicleType VehicleType, location string) {
	vehicle := &Vehicle{
		ID:       id,
		Type:     vehicleType,
		Status:   Active,
		Location: location,
	}
	c.Vehicles = append(c.Vehicles, vehicle)
}
