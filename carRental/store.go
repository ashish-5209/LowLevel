package carRental

// Store represents a specific store that holds a fleet of vehicles.
type Store struct {
	StoreID  int
	Name     string
	Address  string
	Vehicles []Vehicle
}