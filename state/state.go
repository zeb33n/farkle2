package state

type Player struct {
	Name  string
	Score int
}

type GameState struct {
	Dice          []int
	RoundScore    int
	CurrentScore  int
	CurrentPlayer int
	Players       []Player
}
