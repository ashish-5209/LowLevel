package parkingSpot

// Floor represents a floor within the parking lot
type Floor struct {
	Level int
	Spots []*ParkingSpot
}
