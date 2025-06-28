package game

import (
	"fmt"
	"math/rand"

	"github.com/zeb33n/farkle2/score"
	"github.com/zeb33n/farkle2/state"
	"github.com/zeb33n/farkle2/tui"
)

func takeTurn(gamestate *state.GameState) {
	var x int
	fmt.Scan(&x)
	for i := range gamestate.Dice {
		gamestate.Dice[i] = rand.Intn(6) + 1
	}
	round_score, num_dice := score.Score(gamestate.Dice)
	gamestate.RoundScore = round_score
	gamestate.CurrentScore += round_score
	tui.TuiRender(gamestate)
	gamestate.Dice = make([]int, num_dice)
	if round_score == 0 {
		gamestate.CurrentScore = 0
		passTurn(gamestate)
	}
	if num_dice == 0 {
		gamestate.Dice = make([]int, 6)
	}
}

func passTurn(gamestate *state.GameState) {
	gamestate.CurrentPlayer++
	if gamestate.CurrentPlayer == len(gamestate.Players) {
		gamestate.CurrentPlayer = 0
	}
	gamestate.Dice = make([]int, 6)
}

func RunGame(players []string, finalscore int) {
	gamestate := &state.GameState{make([]int, 6), 0, 0, 0, []state.Player{{Name: "bob", Score: 0}}}
	for true {
		if true {
			takeTurn(gamestate)
		} else {
			gamestate.Players[gamestate.CurrentPlayer].Score += gamestate.CurrentScore
			passTurn(gamestate)
		}
	}
}
