package main

import "math"

// Knight Implementation
type Knight struct {
	color bool
	x, y  int
}

func (k *Knight) getPosition() (int, int) {
	return k.x, k.y
}

func (k *Knight) getColor() bool {
	return k.color
}

func (k *Knight) getName() string {
	return "Knight"
}

func (k *Knight) move(toX, toY int, board *Board) bool {
	dx, dy := math.Abs(float64(k.x-toX)), math.Abs(float64(k.y-toY))
	return (dx == 2 && dy == 1) || (dx == 1 && dy == 2) // L-shaped movement
}

func (k *Knight) validMoves(board *Board) []Move {
	directions := [8][2]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
	var moves []Move
	for _, d := range directions {
		newX, newY := k.x+d[0], k.y+d[1]
		if newX >= 0 && newX < 8 && newY >= 0 && newY < 8 {
			moves = append(moves, Move{k.x, k.y, newX, newY})
		}
	}
	return moves
}
