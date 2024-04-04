package tic_toc_game

import "fmt"

// GameManager oversees the creation and management of individual games
type GameManager struct {
	Games map[string]*Game
}

// CreateGame creates a new Tic-Tac-Toe game
func (gm *GameManager) CreateGame(gameID string, player1, player2 Player) *Game {
	board := make([][]*Piece, 3)
	for i := range board {
		board[i] = make([]*Piece, 3)
	}
	game := &Game{
		Board:   board,
		Players: []Player{player1, player2},
	}
	gm.Games[gameID] = game
	return game
}

// GetGame retrieves the game with the given ID
func (gm *GameManager) GetGame(gameID string) *Game {
	if game, ok := gm.Games[gameID]; ok {
		return game
	}
	return nil
}

// MakeMove allows a player to make a move in the game
func (gm *GameManager) MakeMove(gameID string, player Player, row, col int) bool {
	game := gm.GetGame(gameID)
	if game == nil {
		fmt.Println("Game not found")
		return false
	}
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		fmt.Println("Invalid move. Row and column must be between 0 and 2")
		return false
	}
	if game.Board[row][col] != nil {
		fmt.Println("Invalid move. Cell already occupied")
		return false
	}
	game.Board[row][col] = player.Piece
	return true
}
