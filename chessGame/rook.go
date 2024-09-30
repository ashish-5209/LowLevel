package main

// Rook Implementation
type Rook struct {
	color bool
	x, y  int
}

func (r *Rook) getPosition() (int, int) {
	return r.x, r.y
}

func (r *Rook) getColor() bool {
	return r.color
}

func (r *Rook) getName() string {
	return "Rook"
}

func (r *Rook) move(toX, toY int, board *Board) bool {
	if r.x == toX || r.y == toY { // Rook moves in straight lines
		return true
	}
	return false
}

func (r *Rook) validMoves(board *Board) []Move {
	var moves []Move
	for i := 0; i < 8; i++ {
		if i != r.y {
			moves = append(moves, Move{r.x, r.y, r.x, i})
		}
		if i != r.x {
			moves = append(moves, Move{r.x, r.y, i, r.y})
		}
	}
	return moves
}
