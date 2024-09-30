package main

import "time"

// Commentary struct representing commentary on the match
type Commentary struct {
	CommentaryID int
	MatchID      int
	UserID       int
	Text         string
	Timestamp    time.Time
}
