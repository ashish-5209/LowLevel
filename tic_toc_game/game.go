package tic_toc_game

import "fmt"

// Game represents an individual Tic-Tac-Toe game
type Game struct {
	Board   [][]*Piece
	Players []Player
}

// CheckWinner checks if there's a winner in the game
func (game *Game) CheckWinner() bool {
	for i := 0; i < 3; i++ {
		// Check rows
		if game.Board[i][0] != nil && game.Board[i][0] == game.Board[i][1] && game.Board[i][0] == game.Board[i][2] {
			return true
		}
		// Check columns
		if game.Board[0][i] != nil && game.Board[0][i] == game.Board[1][i] && game.Board[0][i] == game.Board[2][i] {
			return true
		}
	}
	// Check diagonals
	if game.Board[0][0] != nil && game.Board[0][0] == game.Board[1][1] && game.Board[0][0] == game.Board[2][2] {
		return true
	}
	if game.Board[0][2] != nil && game.Board[0][2] == game.Board[1][1] && game.Board[0][2] == game.Board[2][0] {
		return true
	}
	return false
}

// DisplayBoard displays the game board
func (game *Game) DisplayBoard() {
	fmt.Println("Game board:")
	for _, row := range game.Board {
		for _, cell := range row {
			if cell == nil {
				fmt.Print("- ")
			} else {
				fmt.Printf("%s ", cell.Symbol)
			}
		}
		fmt.Println()
	}
}
