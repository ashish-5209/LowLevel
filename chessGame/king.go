package main

import "math"

// King Implementation
type King struct {
	color bool
	x, y  int
}

func (k *King) getPosition() (int, int) {
	return k.x, k.y
}

func (k *King) getColor() bool {
	return k.color
}

func (k *King) getName() string {
	return "King"
}

func (k *King) move(toX, toY int, board *Board) bool {
	return math.Abs(float64(k.x-toX)) <= 1 && math.Abs(float64(k.y-toY)) <= 1 // 1-square move in any direction
}

func (k *King) validMoves(board *Board) []Move {
	var moves []Move
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			newX, newY := k.x+dx, k.y+dy
			if newX >= 0 && newX < 8 && newY >= 0 && newY < 8 {
				moves = append(moves, Move{k.x, k.y, newX, newY})
			}
		}
	}
	return moves
}
