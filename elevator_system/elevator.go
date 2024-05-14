// Elevator:
// ID: An integer representing the unique identifier of the elevator.
// CurrentFloor: An integer representing the current floor where the elevator is located.
// Status: An enum representing the status of the elevator (Stopped, Moving, Opening, Closing).
// Destination: An integer representing the destination floor of the elevator.
// InternalQueue: A slice of integers representing the internal queue of floor requests.
// ExternalQueue: A map where keys are floor numbers and values are slices of Direction representing external requests to the elevator.
// Algorithm: A function pointer representing the algorithm used for moving the elevator.
// mux: A mutex for ensuring thread safety.

package elevator_system

import "sync"

type Elevator struct {
	ID            int
	CurrentFloor  int
	Status        ElevatorStatus
	Destination   int
	InternalQueue []int
	ExternalQueue map[int][]Direction
	Algorithm     func(*Elevator)
	mux           sync.Mutex
}

func NewElevator(id int, algorithm func(*Elevator)) *Elevator {
	return &Elevator{
		ID:            id,
		CurrentFloor:  0,
		Status:        Stopped,
		Destination:   0,
		InternalQueue: []int{},
		ExternalQueue: make(map[int][]Direction),
		Algorithm:     algorithm,
	}
}

func (e *Elevator) AddToQueue(floor int) {
	e.mux.Lock()
	defer e.mux.Unlock()

	if e.Status == Stopped {
		e.InternalQueue = append(e.InternalQueue, floor)
		e.Algorithm(e)
	} else {
		direction := Up
		if floor < e.CurrentFloor {
			direction = Down
		}
		e.ExternalQueue[floor] = append(e.ExternalQueue[floor], direction)
	}
}
