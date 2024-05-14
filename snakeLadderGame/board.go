package snakeLadderGame

// Board represents the game board.
type Board struct {
	Size       int
	Snakes     map[int]int
	Ladders    map[int]int
	Players    []*Player
	CurrentPos map[int]int
}
