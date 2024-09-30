// +------------------+           +------------------+           +-----------------+
// |      Game        |           |      Board       |           |     Player      |
// |------------------|           |------------------|           |-----------------|
// | - players: []Player |           | - squares: [8][8]Square|           | - name: string      |
// | - board: Board   |           |------------------|           | - isWhite: bool  |
// | - status: string |           | + getSquare(x, y): Square |   |-----------------|
// |------------------|           | + resetBoard(): void  |      | + makeMove(): bool|
// | + startGame(): void |         | + displayBoard(): void |     | + getPieces(): []Piece |
// | + makeMove(): bool|           +------------------+           +-----------------+
// | + endGame(): void|
// +------------------+

//      |                                                       |
//      v                                                       v
// +-------------------+                                    +-------------------+
// |      Move         |                                    |     Square        |
// |-------------------|                                    |-------------------|
// | - from: Square    |                                    | - x: int          |
// | - to: Square      |                                    | - y: int          |
// | - piece: Piece    |                                    | - isOccupied: bool|
// |-------------------|                                    | - piece: Piece    |
// | + execute(): bool |                                    |-------------------|
// +-------------------+                                    | + isEmpty(): bool |
//                                                          | + getPiece(): Piece|
//                                                          +-------------------+

//      |                                                       |
//      v                                                       v
// +-------------------+         +-------------------+         +--------------------+
// |      Piece        |<>-------|     King          |         |     Queen           |
// |-------------------|         |-------------------|         |--------------------|
// | - color: bool     |         | + isChecked(): bool|         | + move(): bool     |
// | - position: Square|         +-------------------+         +--------------------+
// |-------------------|         +-------------------+         +--------------------+
// | + move(): bool    |         |     Bishop        |         |     Rook            |
// | + validMoves(): []Move |     +-------------------+         +--------------------+
// |-------------------|         | + move(): bool    |         | + move(): bool     |
// |     ^             |         +-------------------+         +--------------------+
// |     |             |         |     Knight        |         |     Pawn            |
// |     |             |         +-------------------+         +--------------------+
// |     |             |         | + move(): bool    |         | + move(): bool     |
// |-------------------|         +-------------------+         +--------------------+
// |   Subclasses of Pieces represent each specific type of chess piece               |
// +-----------------------------------------------------------------------------------+

package main

import (
	"fmt"
)

// Piece represents a chess piece (abstract)
type Piece interface {
	move(toX, toY int, board *Board) bool
	getPosition() (int, int)
	getColor() bool
	getName() string
	validMoves(board *Board) []Move
}

// Move represents a chess move
type Move struct {
	fromX, fromY, toX, toY int
}

// Main function
func main() {
	board := &Board{}
	board.initBoard()

	// Display valid moves for a few pieces
	rook := board.squares[0][0].(*Rook)
	fmt.Println("Valid moves for Rook:", rook.validMoves(board))

	knight := board.squares[0][1].(*Knight)
	fmt.Println("Valid moves for Knight:", knight.validMoves(board))
}
