package main

type Employee interface {
	HandleCall(call Call) bool
	EscalateCall(call Call) bool
	IsAvailable() bool
	GetName() string
}
