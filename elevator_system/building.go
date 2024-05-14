// Building:
// Floors: An integer representing the number of floors in the building.
// Elevators: A slice of pointers to Elevator instances.

package elevator_system

type Building struct {
	Floors    int
	Elevators []*Elevator
}

func NewBuilding(floors, elevators int, algorithm func(*Elevator)) *Building {
	building := &Building{
		Floors:    floors,
		Elevators: make([]*Elevator, elevators),
	}
	for i := 0; i < elevators; i++ {
		building.Elevators[i] = NewElevator(i+1, algorithm)
	}
	return building
}
