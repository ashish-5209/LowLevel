package snakeLadderGame

import "fmt"

// Game represents the Snake and Ladder game.
type Game struct {
	Board *Board
}

// NewGame creates a new Snake and Ladder game.
func NewGame(size int, snakes, ladders map[int]int) *Game {
	board := &Board{
		Size:       size,
		Snakes:     snakes,
		Ladders:    ladders,
		Players:    make([]*Player, 0),
		CurrentPos: make(map[int]int),
	}
	return &Game{board}
}

// PlayGame simulates the Snake and Ladder game.
func (g *Game) PlayGame() {
	for {
		for _, player := range g.Board.Players {
			steps := RollDice()
			fmt.Printf("%s rolled a %d\n", player.Name, steps)
			g.MovePlayer(player.ID, steps)

			fmt.Printf("%s is now at position %d\n", player.Name, g.Board.CurrentPos[player.ID])

			if g.Board.CurrentPos[player.ID] == g.Board.Size {
				fmt.Printf("%s has won the game!\n", player.Name)
				return
			}
		}
	}
}
