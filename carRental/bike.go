package carRental

import "fmt"

// Bike represents a bike and inherits from BaseVehicle.
type Bike struct {
	BaseVehicle
	HasGear bool
}

func (b Bike) GetVehicleType() string {
	return "Bike"
}

func (b Bike) GetPricePerDay() float64 {
	return b.PricePerDay
}

func (b Bike) String() string {
	return fmt.Sprintf("Bike - %s, HasGear: %t", b.BaseVehicle.String(), b.HasGear)
}
