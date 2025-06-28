package game

import (
	"math/rand"

	"github.com/zeb33n/farkle2/score"
	"github.com/zeb33n/farkle2/state"
	"github.com/zeb33n/farkle2/tui"
	"github.com/zeb33n/farkle2/utils"
)

func passTurn(gamestate *state.GameState) {
	gamestate.CurrentPlayer++
	if gamestate.CurrentPlayer == len(gamestate.Players) {
		gamestate.CurrentPlayer = 0
	}
	gamestate.CurrentScore = 0
	gamestate.Dice = make([]int, 6)
}

func takeTurn(gamestate *state.GameState) {
	for i := range gamestate.Dice {
		gamestate.Dice[i] = rand.Intn(6) + 1
	}
	round_score, num_dice := score.Score(gamestate.Dice)
	gamestate.RoundScore = round_score
	gamestate.CurrentScore += round_score
	tui.TuiRenderGamestate(gamestate)
	gamestate.Dice = make([]int, num_dice)
	if round_score == 0 {
		gamestate.CurrentScore = 0
		passTurn(gamestate)
	}
	if num_dice == 0 {
		gamestate.Dice = make([]int, 6)
	}
}

func RunGame(players []string, finalscore int) {
	gamestate := &state.GameState{make([]int, 6), 0, 0, 0, []state.Player{{Name: "zeb", Score: 0}, {Name: "will", Score: 0}}}
	for true {
		x := utils.WaitForKeypress()
		if x == "r" {
			takeTurn(gamestate)
		} else {
			gamestate.Players[gamestate.CurrentPlayer].Score += gamestate.CurrentScore
			passTurn(gamestate)
			takeTurn(gamestate)
		}
	}
}
