package main

// Board represents the chessboard (8x8 grid)
type Board struct {
	squares [8][8]Piece
}

// Initialize the chessboard
func (b *Board) initBoard() {
	// Add Pawns
	for i := 0; i < 8; i++ {
		b.squares[1][i] = &Pawn{true, 1, i}
		b.squares[6][i] = &Pawn{false, 6, i}
	}

	// Add Rooks
	b.squares[0][0], b.squares[0][7] = &Rook{true, 0, 0}, &Rook{true, 0, 7}
	b.squares[7][0], b.squares[7][7] = &Rook{false, 7, 0}, &Rook{false, 7, 7}

	// Add Knights
	b.squares[0][1], b.squares[0][6] = &Knight{true, 0, 1}, &Knight{true, 0, 6}
	b.squares[7][1], b.squares[7][6] = &Knight{false, 7, 1}, &Knight{false, 7, 6}

	// Add Bishops
	b.squares[0][2], b.squares[0][5] = &Bishop{true, 0, 2}, &Bishop{true, 0, 5}
	b.squares[7][2], b.squares[7][5] = &Bishop{false, 7, 2}, &Bishop{false, 7, 5}

	// Add Queens
	b.squares[0][3] = &Queen{true, 0, 3}
	b.squares[7][3] = &Queen{false, 7, 3}

	// Add Kings
	b.squares[0][4] = &King{true, 0, 4}
	b.squares[7][4] = &King{false, 7, 4}
}
