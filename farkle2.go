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
	num_dice       int
	current_score  int
	current_player int
	players        []Player
}

func take_turn(gamestate GameState) {
	dice := make([]int, gamestate.num_dice)
	for i := 0; i < gamestate.num_dice; i++ {
		dice[i] = rand.Intn(6) + 1
	}
	println("Rolled: ")
	fmt.Printf("%v\n", dice)
	round_score, num_dice := score.Score(dice)
	println("Score: ")
	println(round_score)
	if round_score == 0 {
		pass_turn(gamestate)
	}
	gamestate.current_score += round_score
	gamestate.num_dice = num_dice
}

func pass_turn(gamestate GameState) {
	gamestate.current_player++
	if gamestate.current_player == len(gamestate.players) {
		gamestate.current_player = 0
	}
	gamestate.num_dice = 6
}

func main() {
	gamestate := GameState{6, 0, 0, []Player{{name: "bob", score: 0}}}
	for true {
		if true {
			take_turn(gamestate)
		} else {
			gamestate.players[gamestate.current_player].score += gamestate.current_score
			pass_turn(gamestate)
		}
	}
	take_turn(gamestate)
}
