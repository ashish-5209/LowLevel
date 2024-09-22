package main

import "fmt"

type CallCenter struct {
	freshers  []Freshers
	leads     []Lead
	manager   Manager
	callQueue CallQueue
}

func (cc *CallCenter) RouteCall(call Call) {
	for i := range cc.freshers {
		if cc.freshers[i].IsAvailable() {
			cc.freshers[i].HandleCall(call)
			return
		}
	}

	for i := range cc.leads {
		if cc.leads[i].IsAvailable() {
			cc.leads[i].HandleCall(call)
			return
		}
	}

	if cc.manager.IsAvailable() {
		cc.manager.HandleCall(call)
	} else {
		fmt.Println("All employees are busy. Adding call to queue.")
		cc.callQueue.AddCall(call)
	}
}

func (cc *CallCenter) ManageQueue() {
	call := cc.callQueue.GetNextCall()
	if call != nil {
		cc.RouteCall(*call)
	}
}
