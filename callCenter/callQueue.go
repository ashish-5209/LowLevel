package main

type CallQueue struct {
	queue []Call
}

func (cq *CallQueue) AddCall(call Call) {
	cq.queue = append(cq.queue, call)
}

func (cq *CallQueue) GetNextCall() *Call {
	if len(cq.queue) == 0 {
		return nil
	}
	call := cq.queue[0]
	cq.queue = cq.queue[1:]
	return &call
}
