package main

import (
	"fmt"
	"time"
)

// Match struct representing a cricket match
type Match struct {
	MatchID   int
	MatchType string // e.g., Test, ODI, T20
	Date      time.Time
	Venue     string
	Status    string // e.g., Live, Finished
	TeamA     Team
	TeamB     Team
	Scorecard Scorecard
}

// Methods for Match struct
func (m *Match) StartMatch() {
	m.Status = "Live"
	fmt.Printf("Match %d has started.\n", m.MatchID)
}

func (m *Match) EndMatch() {
	m.Status = "Finished"
	fmt.Printf("Match %d has ended.\n", m.MatchID)
}
