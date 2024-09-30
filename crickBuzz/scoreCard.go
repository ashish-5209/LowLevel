package main

// Scorecard struct representing the scorecard for the match
type Scorecard struct {
	Innings []Innings
}

// Methods for Scorecard struct
func (s *Scorecard) UpdateScore(inningID int, runs int) {
	for i := range s.Innings {
		if s.Innings[i].InningsID == inningID {
			s.Innings[i].Runs += runs
			break
		}
	}
}
