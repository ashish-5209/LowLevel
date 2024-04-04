package parkingSpot

import "sync"

// ParkingSpot represents an individual parking spot
type ParkingSpot struct {
	SpotNumber int
	IsOccupied bool
	Type       ParkingSpotType
	Vehicle    *Vehicle
	Lock       sync.Mutex
}
