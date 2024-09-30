package main

import "math"

// Bishop Implementation
type Bishop struct {
	color bool
	x, y  int
}

func (b *Bishop) getPosition() (int, int) {
	return b.x, b.y
}

func (b *Bishop) getColor() bool {
	return b.color
}

func (b *Bishop) getName() string {
	return "Bishop"
}

func (b *Bishop) move(toX, toY int, board *Board) bool {
	return math.Abs(float64(b.x-toX)) == math.Abs(float64(b.y-toY)) // Diagonal movement
}

func (b *Bishop) validMoves(board *Board) []Move {
	var moves []Move
	for i := 1; i < 8; i++ {
		if b.x+i < 8 && b.y+i < 8 {
			moves = append(moves, Move{b.x, b.y, b.x + i, b.y + i})
		}
		if b.x-i >= 0 && b.y-i >= 0 {
			moves = append(moves, Move{b.x, b.y, b.x - i, b.y - i})
		}
	}
	return moves
}
