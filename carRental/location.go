package carRental

import "time"

// Location represents a geographical location where stores are present.
type Location struct {
	LocationID int
	Name       string
	Stores     []Store
}

// SearchVehicle allows a user to search for available vehicles in a location within a date range.
func SearchVehicle(location Location, startDate, endDate time.Time) []Vehicle {
	var availableVehicles []Vehicle
	for _, store := range location.Stores {
		for _, vehicle := range store.Vehicles {
			if vehicle.IsAvailable() {
				availableVehicles = append(availableVehicles, vehicle)
			}
		}
	}
	return availableVehicles
}
