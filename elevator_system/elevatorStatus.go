package elevator_system

type ElevatorStatus int

const (
	Stopped ElevatorStatus = iota
	Moving
	Opening
	Closing
)
