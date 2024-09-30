package main

// Notification struct representing a notification for a user
type Notification struct {
	NotificationID int
	UserID         int
	Message        string
	Type           string // e.g., Score Update, Match Alert
}
