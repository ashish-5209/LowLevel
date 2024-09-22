// +--------------------------------------------------+
// |                  CallCenter                      |
// +--------------------------------------------------+
// | - Freshers: []*Fresher                           |
// | - Leads: []*Lead                                 |
// | - Manager: *Manager                              |
// | - CallQueue: *CallQueue                          |
// +--------------------------------------------------+
// | + RouteCall(call Call)                           |
// | + HandleEscalation(call Call)                    |
// | + ManageQueue()                                  |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                    Fresher                       |
// +--------------------------------------------------+
// | - Name: string                                   |
// | - IsAvailable: bool                              |
// +--------------------------------------------------+
// | + HandleCall(call Call)                          |
// | + EscalateCall(call Call)                        |
// | + IsAvailable() bool                             |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                     Lead                         |
// +--------------------------------------------------+
// | - Name: string                                   |
// | - IsAvailable: bool                              |
// +--------------------------------------------------+
// | + HandleCall(call Call)                          |
// | + EscalateCall(call Call)                        |
// | + IsAvailable() bool                             |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                    Manager                       |
// +--------------------------------------------------+
// | - Name: string                                   |
// | - IsAvailable: bool                              |
// +--------------------------------------------------+
// | + HandleCall(call Call)                          |
// | + IsAvailable() bool                             |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                    Call                          |
// +--------------------------------------------------+
// | - CallerID: string                               |
// | - Issue: string                                  |
// | - EscalationLevel: int                           |
// +--------------------------------------------------+

// +--------------------------------------------------+
// |                  CallQueue                       |
// +--------------------------------------------------+
// | - Queue: []Call                                  |
// +--------------------------------------------------+
// | + AddCall(call Call)                             |
// | + GetNextCall() *Call                            |
// +--------------------------------------------------+

package main

func main() {
	f1 := Freshers{name: "Freshers1", isAvalaiable: true}
	f2 := Freshers{name: "Freshers2", isAvalaiable: true}
	l1 := Lead{name: "Lead1", isAvalaiable: true}
	m := Manager{name: "Manager1", isAvalaiable: true}

	callCenter := CallCenter{
		freshers:  []Freshers{f1, f2},
		leads:     []Lead{l1},
		manager:   m,
		callQueue: CallQueue{},
	}

	call1 := Call{CallerId: "Caller1", Issue: "Issue 1"}
	call2 := Call{CallerId: "Caller2", Issue: "Issue 2"}
	call3 := Call{CallerId: "Caller3", Issue: "Issue 3"}

	// Simulate incoming calls
	callCenter.RouteCall(call1)
	callCenter.RouteCall(call2)
	callCenter.RouteCall(call3)

	// Check if any queued calls can be processed
	callCenter.ManageQueue()
}
