package snakeLadderGame

func App() {
	snakes := map[int]int{
		16: 6,
		47: 26,
		49: 11,
		56: 53,
		62: 19,
		64: 60,
		87: 24,
		93: 73,
		95: 75,
		98: 78,
	}

	ladders := map[int]int{
		1:  38,
		4:  14,
		9:  31,
		21: 42,
		28: 84,
		36: 44,
		51: 67,
		71: 91,
		80: 100,
	}

	game := NewGame(100, snakes, ladders)
	game.AddPlayer(1, "Player 1", 0)
	game.AddPlayer(2, "Player 2", 0)
	game.AddPlayer(3, "Player 3", 0)

	game.PlayGame()
}
