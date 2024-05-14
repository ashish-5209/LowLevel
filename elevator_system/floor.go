// Functions:
// NewBuilding(floors int, elevators int, algorithm func(*Elevator)): A function to create a new Building instance with the specified number of floors and elevators, and with the given algorithm.
// NewElevator(id int, algorithm func(*Elevator)): A function to create a new Elevator instance with the specified ID and algorithm.
// MoveTo(floor int): A method to move the elevator to the specified floor.
// OpenDoor(): A method to open the door of the elevator.
// CloseDoor(): A method to close the door of the elevator.
// AddToQueue(floor int): A method to add a floor request to the elevator's queue.

package elevator_system

import (
	"fmt"
	"time"
)

func (e *Elevator) MoveTo(floor int) {
	if e.CurrentFloor == floor {
		return
	}

	var direction Direction
	if e.CurrentFloor < floor {
		direction = Up
	} else {
		direction = Down
	}

	e.Status = Moving
	for e.CurrentFloor != floor {
		e.CurrentFloor += int(direction)
		fmt.Printf("Elevator %d is moving to floor %d\n", e.ID, e.CurrentFloor)
		time.Sleep(time.Second)
	}

	fmt.Printf("Elevator %d reached floor %d\n", e.ID, e.CurrentFloor)
	e.Status = Stopped
}

func (e *Elevator) OpenDoor() {
	e.Status = Opening
	fmt.Printf("Elevator %d door is opening\n", e.ID)
	time.Sleep(time.Second)
	fmt.Printf("Elevator %d door is open\n", e.ID)
	e.Status = Stopped
}

func (e *Elevator) CloseDoor() {
	e.Status = Closing
	fmt.Printf("Elevator %d door is closing\n", e.ID)
	time.Sleep(time.Second)
	fmt.Printf("Elevator %d door is closed\n", e.ID)
	e.Status = Stopped
}
