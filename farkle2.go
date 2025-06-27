package main

import (
	"fmt"
	"math/rand"

	score "github.com/zeb33n/farkle2/score"
)

type Player struct {
	name  string
	score int
}

type GameState struct {
	dice           []int
	round_score    int
	current_score  int
	current_player int
	players        []Player
}

func display_score(gamestate *GameState) {
	for i := 0; i < 5; i++ {
		fmt.Printf("\033[1A\033[2K")
	}
	dice_sides := []string{"[.]", "[:]", "[.:]", "[::]", "[:.:]", "[:::]"}
	for _, e := range gamestate.dice {
		fmt.Printf("%s ", dice_sides[e-1])
	}
	fmt.Printf(
		"\nRoll: %d\nScore: %d\nPlayer: %s\n",
		gamestate.round_score,
		gamestate.current_score,
		gamestate.players[gamestate.current_player].name,
	)
}

func take_turn(gamestate *GameState) {
	var x int
	fmt.Scan(&x)
	for i := range gamestate.dice {
		gamestate.dice[i] = rand.Intn(6) + 1
	}
	round_score, num_dice := score.Score(gamestate.dice)
	gamestate.round_score = round_score
	gamestate.current_score += round_score
	display_score(gamestate)
	gamestate.dice = make([]int, num_dice)
	if round_score == 0 {
		gamestate.current_score = 0
		pass_turn(gamestate)
	}
	if num_dice == 0 {
		gamestate.dice = make([]int, 6)
	}
}

func pass_turn(gamestate *GameState) {
	gamestate.current_player++
	if gamestate.current_player == len(gamestate.players) {
		gamestate.current_player = 0
	}
	gamestate.dice = make([]int, 6)
}

func main() {
	fmt.Print("\n\n\n\n")
	gamestate := &GameState{make([]int, 6), 0, 0, 0, []Player{{name: "bob", score: 0}}}
	for true {
		if true {
			take_turn(gamestate)
		} else {
			gamestate.players[gamestate.current_player].score += gamestate.current_score
			pass_turn(gamestate)
		}
	}
}
