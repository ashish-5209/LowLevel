package main

import "sort"

type LeaderBoardEntry struct {
	UserId string
	Score  int
}

type LeaderBoardService struct {
	leaderBoard map[string]*LeaderBoardEntry
}

func NewLeaderBoardService() *LeaderBoardService {
	return &LeaderBoardService{
		leaderBoard: make(map[string]*LeaderBoardEntry),
	}
}

func (l *LeaderBoardService) UpdateLeaderBoard(userID string, score int) error {
	entry, exists := l.leaderBoard[userID]
	if exists {
		entry.Score += score
	} else {
		l.leaderBoard[userID] = &LeaderBoardEntry{
			UserId: userID,
			Score:  score,
		}
	}
	return nil
}

func (l *LeaderBoardService) GetLeaderBoard() []*LeaderBoardEntry {
	entries := []*LeaderBoardEntry{}
	for _, entry := range l.leaderBoard {
		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Score > entries[j].Score
	})

	return entries
}
