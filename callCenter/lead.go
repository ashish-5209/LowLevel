package main

import "fmt"

type Lead struct {
	name         string
	isAvalaiable bool
}

func (l *Lead) HandleCall(call Call) bool {
	if l.isAvalaiable {
		l.isAvalaiable = false
		fmt.Println(l.name, "is handling the call from", call.CallerId)
		return true
	}
	return false
}

func (l *Lead) EscalateCall(call Call) bool {
	fmt.Println(l.name, "is escalating the call to the manager")
	return true
}

func (l *Lead) IsAvailable() bool {
	return l.isAvalaiable
}

func (l *Lead) GetName() string {
	return l.name
}
