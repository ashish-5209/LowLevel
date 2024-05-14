// ____________________             ____________________
// |      GameManager   |           |       Game        |
// |--------------------|           |-------------------|
// | - games: map[string]*Game|      | - board: [][]Piece|
// |____________________|           | - players: []Player|
// | + CreateGame(gameID: string, player1: Player, player2: Player): Game|
// | + StartGame(gameID: string)|
// | + GetGame(gameID: string): *Game|
// | + MakeMove(gameID: string, player: Player, row: int, col: int): bool|
// |____________________|           |___________________|

//           ^
//           |
//           |
//     ______|______
//    |             |
//    |             |
//   _|___       ___|_
//  |User1|     |User2|
//  |_____|     |_____|
//      |
//      |
//      |
//  ______|______
// |   Piece     |
// |-------------|
// | - symbol: string|
// |_____________|

package tic_toc_game

import (
	"fmt"
)

func App() {
	gm := GameManager{
		Games: make(map[string]*Game),
	}

	fmt.Println(gm)

	// Create players
	player1 := Player{Name: "Player 1", Piece: &Piece{Symbol: "X"}}
	player2 := Player{Name: "Player 2", Piece: &Piece{Symbol: "O"}}

	// Create a new game
	game := gm.CreateGame("game1", player1, player2)

	// Game loop
	currentPlayer := player1
	for {
		game.DisplayBoard()

		var row, col int
		fmt.Printf("%s's turn (row column): ", currentPlayer.Name)
		fmt.Scanln(&row, &col)

		if !gm.MakeMove("game1", currentPlayer, row, col) {
			continue
		}

		if game.CheckWinner() {
			fmt.Printf("%s wins!\n", currentPlayer.Name)
			break
		}

		// Switch to the other player
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}
