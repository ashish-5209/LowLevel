package main

// Innings struct representing an innings in the match
type Innings struct {
	InningsID   int
	BattingTeam Team
	Runs        int
	Wickets     int
	BallsFaced  int
	Balls       []Ball
}

// Methods for Innings struct
func (i *Innings) AddBall(ball Ball) {
	i.Balls = append(i.Balls, ball)
	i.BallsFaced++
	i.Runs += ball.RunsScored
	if ball.BallType == "Wicket" {
		i.Wickets++
	}
}
