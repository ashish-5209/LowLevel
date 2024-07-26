package carRental

import "fmt"

// BaseVehicle provides common attributes for all vehicles.
type BaseVehicle struct {
	VehicleID   int
	Model       string
	Make        string
	Year        int
	PricePerDay float64
	Status      string // Available, Reserved, Out of Service
}

// Vehicle is an interface for different types of vehicles.
type Vehicle interface {
	GetVehicleType() string
	GetPricePerDay() float64
	GetID() int
	IsAvailable() bool
	SetStatus(status string)
	String() string // Added String method for custom output
}

func (bv *BaseVehicle) GetID() int {
	return bv.VehicleID
}

func (bv *BaseVehicle) IsAvailable() bool {
	return bv.Status == "Available"
}

func (bv *BaseVehicle) SetStatus(status string) {
	bv.Status = status
}

func (bv BaseVehicle) String() string {
	return fmt.Sprintf("ID: %d, Model: %s, Make: %s, Year: %d, PricePerDay: %.2f, Status: %s",
		bv.VehicleID, bv.Model, bv.Make, bv.Year, bv.PricePerDay, bv.Status)
}
