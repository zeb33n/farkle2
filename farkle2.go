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
	current_score  int
	current_player int
	players        []Player
}

func take_turn(gamestate *GameState) {
	var x int
	fmt.Scan(&x)
	for i := range gamestate.dice {
		gamestate.dice[i] = rand.Intn(6) + 1
	}
	round_score, num_dice := score.Score(gamestate.dice)
	gamestate.current_score += round_score
	fmt.Printf("%v\n", gamestate)
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
	gamestate := &GameState{make([]int, 6), 0, 0, []Player{{name: "bob", score: 0}}}
	for true {
		if true {
			take_turn(gamestate)
		} else {
			gamestate.players[gamestate.current_player].score += gamestate.current_score
			pass_turn(gamestate)
		}
	}
}
