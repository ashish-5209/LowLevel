package main

import "fmt"

type Manager struct {
	name         string
	isAvalaiable bool
}

func (m *Manager) HandleCall(call Call) bool {
	if m.isAvalaiable {
		m.isAvalaiable = false
		fmt.Println(m.name, "is handling the call from", call.CallerId)
		return true
	}
	return false
}

func (m *Manager) IsAvailable() bool {
	return m.isAvalaiable
}

func (m *Manager) GetName() string {
	return m.name
}
