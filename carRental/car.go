package carRental

import "fmt"

// Car represents a car and inherits from BaseVehicle.
type Car struct {
	BaseVehicle
	NumberOfDoors int
}

func (c Car) GetVehicleType() string {
	return "Car"
}

func (c Car) GetPricePerDay() float64 {
	return c.PricePerDay
}

func (c Car) String() string {
	return fmt.Sprintf("Car - %s, NumberOfDoors: %d", c.BaseVehicle.String(), c.NumberOfDoors)
}
