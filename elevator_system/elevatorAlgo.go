// Algos:
// OddDoors(e *Elevator): A function that defines the algorithm for moving the elevator with odd door opening behavior.
// EvenDoors(e *Elevator): A function that defines the algorithm for moving the elevator with even door opening behavior.
// MinTime(e *Elevator): A function that defines the algorithm for moving the elevator with minimum time to destination.

package elevator_system

import (
	"math"
	"sort"
)

func OddDoors(e *Elevator) {
	sort.Ints(e.InternalQueue)
	for len(e.InternalQueue) > 0 {
		nextFloor := e.InternalQueue[0]
		e.InternalQueue = e.InternalQueue[1:]
		e.MoveTo(nextFloor)
		e.OpenDoor()
		e.CloseDoor()
	}
}

func EvenDoors(e *Elevator) {
	sort.Sort(sort.Reverse(sort.IntSlice(e.InternalQueue)))
	for len(e.InternalQueue) > 0 {
		nextFloor := e.InternalQueue[0]
		e.InternalQueue = e.InternalQueue[1:]
		e.MoveTo(nextFloor)
		e.OpenDoor()
		e.CloseDoor()
	}
}

func MinTime(e *Elevator) {
	sort.Slice(e.InternalQueue, func(i, j int) bool {
		return math.Abs(float64(e.CurrentFloor-e.InternalQueue[i])) < math.Abs(float64(e.CurrentFloor-e.InternalQueue[j]))
	})

	for len(e.InternalQueue) > 0 {
		nextFloor := e.InternalQueue[0]
		e.InternalQueue = e.InternalQueue[1:]
		e.MoveTo(nextFloor)
		e.OpenDoor()
		e.CloseDoor()
	}
}
