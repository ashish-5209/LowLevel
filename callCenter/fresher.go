package main

import "fmt"

type Freshers struct {
	name         string
	isAvalaiable bool
}

func (f *Freshers) HandleCall(call Call) bool {
	if f.isAvalaiable {
		f.isAvalaiable = false
		fmt.Println(f.name, "is handling the call from", call.CallerId)
		return true
	}
	return false
}

func (f *Freshers) EscalateCall(call Call) bool {
	fmt.Println(f.name, "is escalating the call from", call.CallerId)
	return true
}

func (f *Freshers) IsAvailable() bool {
	return f.isAvalaiable
}

func (f *Freshers) GetName() string {
	return f.name
}
