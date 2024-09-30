package main

// Player struct representing a player in the cricket match
type Player struct {
	PlayerID int
	Name     string
	Role     string // e.g., Batsman, Bowler, All-rounder
	Stats    PlayerStats
}
