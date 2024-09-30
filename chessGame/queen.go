package main

import "math"

// Queen Implementation
type Queen struct {
	color bool
	x, y  int
}

func (q *Queen) getPosition() (int, int) {
	return q.x, q.y
}

func (q *Queen) getColor() bool {
	return q.color
}

func (q *Queen) getName() string {
	return "Queen"
}

func (q *Queen) move(toX, toY int, board *Board) bool {
	return math.Abs(float64(q.x-toX)) == math.Abs(float64(q.y-toY)) || q.x == toX || q.y == toY // Diagonal or straight
}

func (q *Queen) validMoves(board *Board) []Move {
	var moves []Move
	for i := 1; i < 8; i++ {
		// Diagonal moves
		if q.x+i < 8 && q.y+i < 8 {
			moves = append(moves, Move{q.x, q.y, q.x + i, q.y + i})
		}
		if q.x-i >= 0 && q.y-i >= 0 {
			moves = append(moves, Move{q.x, q.y, q.x - i, q.y - i})
		}
		// Straight moves
		if q.x+i < 8 {
			moves = append(moves, Move{q.x, q.y, q.x + i, q.y})
		}
		if q.y+i < 8 {
			moves = append(moves, Move{q.x, q.y, q.x, q.y + i})
		}
	}
	return moves
}
