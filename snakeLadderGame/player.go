package snakeLadderGame

import (
	"math/rand"
	"time"
)

// Player represents a player in the game.
type Player struct {
	ID     int
	Name   string
	Pos    int
	Wallet float64
}

// RollDice simulates a dice roll.
func RollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

// AddPlayer adds a player to the game.
func (g *Game) AddPlayer(id int, name string, wallet float64) {
	player := &Player{
		ID:     id,
		Name:   name,
		Wallet: wallet,
	}
	g.Board.Players = append(g.Board.Players, player)
	g.Board.CurrentPos[id] = 1
}

// MovePlayer moves the player on the board.
func (g *Game) MovePlayer(playerID int, steps int) {
	currentPos := g.Board.CurrentPos[playerID]
	newPos := currentPos + steps

	// Check if the new position has a ladder or a snake
	if ladder, ok := g.Board.Ladders[newPos]; ok {
		newPos = ladder
	} else if snake, ok := g.Board.Snakes[newPos]; ok {
		newPos = snake
	}

	if newPos <= g.Board.Size {
		g.Board.CurrentPos[playerID] = newPos
	}
}
