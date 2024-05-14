// +--------------------------------------------------+
// |                     Building                     |
// +--------------------------------------------------+
// | - Floors: int                                    |
// | - Elevators: []*Elevator                        |
// +--------------------------------------------------+
// | + NewBuilding(floors int, elevators int,         |
// |               algorithm func(*Elevator))         |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                    Elevator                     |
// +--------------------------------------------------+
// | - ID: int                                        |
// | - CurrentFloor: int                              |
// | - Status: ElevatorStatus                        |
// | - Destination: int                               |
// | - InternalQueue: []int                           |
// | - ExternalQueue: map[int][]Direction             |
// | - Algorithm: func(*Elevator)                    |
// | - mux: sync.Mutex                                |
// +--------------------------------------------------+
// | + NewElevator(id int, algorithm func(*Elevator)) |
// | + MoveTo(floor int)                              |
// | + OpenDoor()                                     |
// | + CloseDoor()                                    |
// | + AddToQueue(floor int)                          |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                      Algos                       |
// +--------------------------------------------------+
// | - OddDoors(e *Elevator)                          |
// | - EvenDoors(e *Elevator)                         |
// | - MinTime(e *Elevator)                           |
// +--------------------------------------------------+

package elevator_system

func App() {
	building := NewBuilding(10, 1, MinTime)

	// Example: Elevator 1
	elevator := building.Elevators[0]

	// Internal Buttons Pressed
	elevator.AddToQueue(8)
	elevator.AddToQueue(5)
	elevator.AddToQueue(1)
	elevator.AddToQueue(4)

	// Someone pressed a button on the 3rd floor
	elevator.AddToQueue(3)
}
