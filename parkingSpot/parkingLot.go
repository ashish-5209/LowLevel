package parkingSpot

import "sync"

// ParkingLot represents the parking lot
type ParkingLot struct {
	Capacity       int
	AvailableSpots int
	Floors         []*Floor
	BillingRecords map[string]Bill
	Payments       map[string]Payment
	Lock           sync.Mutex
}

// NewParkingLot initializes a new parking lot with the given capacity and number of floors
func NewParkingLot(capacity, numFloors int) *ParkingLot {
	floors := make([]*Floor, numFloors)
	for i := 0; i < numFloors; i++ {
		floors[i] = &Floor{
			Level: i + 1,
			Spots: make([]*ParkingSpot, capacity),
		}
		for j := 0; j < capacity; j++ {
			floors[i].Spots[j] = &ParkingSpot{SpotNumber: j + 1}
		}
	}
	return &ParkingLot{
		Capacity:       capacity,
		AvailableSpots: capacity * numFloors,
		Floors:         floors,
		BillingRecords: make(map[string]Bill),
		Payments:       make(map[string]Payment),
	}
}
