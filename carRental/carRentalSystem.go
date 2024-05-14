package carRental

type CarRentalSystem struct {
	Vehicles []*Vehicle
	Bills    []*Bill
}

func NewCarRentalSystem() *CarRentalSystem {
	return &CarRentalSystem{
		Vehicles: make([]*Vehicle, 0),
		Bills:    make([]*Bill, 0),
	}
}
