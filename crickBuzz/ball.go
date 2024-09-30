package main

// Ball struct representing a ball bowled in the innings
type Ball struct {
	BallID     int
	Over       int
	RunsScored int
	BallType   string // e.g., Delivery, No Ball, Wide
}
