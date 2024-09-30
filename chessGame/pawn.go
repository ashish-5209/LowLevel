package main

// Pawn Implementation
type Pawn struct {
	color bool
	x, y  int
}

func (p *Pawn) getPosition() (int, int) {
	return p.x, p.y
}

func (p *Pawn) getColor() bool {
	return p.color
}

func (p *Pawn) getName() string {
	return "Pawn"
}

func (p *Pawn) move(toX, toY int, board *Board) bool {
	direction := 1
	if !p.color {
		direction = -1
	}
	// Regular move
	if toX == p.x+direction && toY == p.y {
		return true
	}
	// First move can jump 2 squares
	if (p.x == 1 && p.color) || (p.x == 6 && !p.color) {
		if toX == p.x+2*direction && toY == p.y {
			return true
		}
	}
	// Diagonal capture
	if toX == p.x+direction && (toY == p.y+1 || toY == p.y-1) {
		return true
	}
	return false
}

func (p *Pawn) validMoves(board *Board) []Move {
	var moves []Move
	direction := 1
	if !p.color {
		direction = -1
	}
	if p.x+direction >= 0 && p.x+direction < 8 {
		moves = append(moves, Move{p.x, p.y, p.x + direction, p.y})
	}
	return moves
}
