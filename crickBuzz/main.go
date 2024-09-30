// +-------------------+         +-------------------+
// | Match             |<------->| Team              |
// |-------------------|         |-------------------|
// | matchID           |         | teamID            |
// | matchType         |         | name              |
// | date              |         | players[]: Player |
// | venue             |         |-------------------|
// | status            |         | +addPlayer()      |
// |-------------------|         | +removePlayer()   |
// | +startMatch()     |         | +getTeamStats()   |
// | +endMatch()       |         +-------------------+
// | +getScorecard()   |
// | +getMatchDetails()|
// +-------------------+
//       |
//       |
//       v
// +-------------------+         +-------------------+
// | Innings           |<------->| Ball              |
// |-------------------|         |-------------------|
// | inningsID         |         | ballID            |
// | battingTeam       |         | over              |
// | runs              |         | runsScored        |
// | wickets           |         | ballType          |
// | ballsFaced        |         |-------------------|
// |-------------------|         | +getBallDetails() |
// | +addBall()        |         +-------------------+
// | +getInningsStats()|
// +-------------------+
//       |
//       |
//       v
// +-------------------+         +-------------------+
// | Scorecard         |         | Commentary         |
// |-------------------|         |-------------------|
// | scorecardID       |         | commentaryID      |
// | match: Match      |         | matchID           |
// | innings: Innings  |         | userID            |
// |-------------------|         | text              |
// | +updateScore()    |         | timestamp         |
// | +getScoreDetails()|         |-------------------|
// +-------------------+         | +addCommentary()  |
//                               +-------------------+
//       |
//       |
//       v
// +-------------------+         +-------------------+
// | User              |         | Notification      |
// |-------------------|         |-------------------|
// | userID            |         | notificationID    |
// | name              |         | userID            |
// | email             |         | message           |
// | phoneNumber       |         | type              |
// |-------------------|         |-------------------|
// | +register()       |         | +sendNotification()|
// | +login()          |         +-------------------+
// | +subscribe()      |
// +-------------------+
//       |
//       |
//       v
// +-------------------+
// | Statistics        |
// |-------------------|
// | playerStats       |
// | matchStats        |
// |-------------------|
// | +getPlayerStats() |
// | +getMatchStats()  |
// +-------------------+

package main

import "time"

// Sample Usage
func main() {
	// Create players
	player1 := Player{PlayerID: 1, Name: "Player 1", Role: "Batsman"}
	player2 := Player{PlayerID: 2, Name: "Player 2", Role: "Bowler"}

	// Create teams
	teamA := Team{TeamID: 1, Name: "Team A", Players: []Player{player1, player2}}
	teamB := Team{TeamID: 2, Name: "Team B"}

	// Create a match
	match := Match{
		MatchID:   1,
		MatchType: "T20",
		Date:      time.Now(),
		Venue:     "Stadium A",
		Status:    "Upcoming",
		TeamA:     teamA,
		TeamB:     teamB,
		Scorecard: Scorecard{Innings: []Innings{}},
	}

	// Start the match
	match.StartMatch()

	// Create innings
	innings := Innings{InningsID: 1, BattingTeam: teamA}

	// Add balls to innings
	ball1 := Ball{BallID: 1, Over: 1, RunsScored: 4, BallType: "Delivery"}
	innings.AddBall(ball1)

	// Update score
	match.Scorecard.UpdateScore(innings.InningsID, ball1.RunsScored)

	// End the match
	match.EndMatch()

	// Register a user
	user := User{UserID: 1, Name: "John Doe", Email: "john@example.com", PhoneNumber: "1234567890"}
	user.Register()
	user.Subscribe()
}
