// Package core
package core

import (
	"math/rand"
)

type Player struct {
	Name  string
	Score int
}

type GameState struct {
	Dice          []int
	ScoringDice   []int
	RoundScore    int
	CurrentScore  int
	CurrentPlayer int
	Players       []Player
}

type Game struct {
	IO inputOutput
}

func (g *Game) passTurn(gamestate *GameState) {
	gamestate.CurrentPlayer++
	if gamestate.CurrentPlayer == len(gamestate.Players) {
		gamestate.CurrentPlayer = 0
	}
	gamestate.CurrentScore = 0
	gamestate.Dice = make([]int, 6)
	g.IO.OutputTurnChange(gamestate)
}

func (g *Game) takeTurn(gamestate *GameState) {
	for i := range gamestate.Dice {
		gamestate.Dice[i] = rand.Intn(6) + 1
	}
	roundScore, numDice, positions := Score(gamestate.Dice)
	gamestate.ScoringDice = positions
	gamestate.RoundScore = roundScore
	gamestate.CurrentScore += roundScore
	g.IO.OutputGamestate(gamestate)
	gamestate.Dice = make([]int, numDice)
	if roundScore == 0 {
		gamestate.CurrentScore = 0
		g.passTurn(gamestate)
	}
	if numDice == 0 {
		gamestate.Dice = make([]int, 6)
	}
}

func checkForWinner(players []Player, finalscore int) bool {
	for _, player := range players {
		if player.Score >= finalscore {
			return false
		}
	}
	return true
}

func (g *Game) RunGame(splayers *map[string]bool, finalscore int) {
	players := []Player{}
	for k := range *splayers {
		players = append(players, Player{Name: k, Score: 0})
	}
	gamestate := &GameState{
		Dice:          make([]int, 6),
		ScoringDice:   []int{},
		RoundScore:    0,
		CurrentScore:  0,
		CurrentPlayer: 0,
		Players:       players,
	}
	g.IO.OutputTurnChange(gamestate)
	for checkForWinner(gamestate.Players, finalscore) {
		msg := g.IO.AwaitInputPlayer(
			gamestate.Players[gamestate.CurrentPlayer].Name,
			gamestate,
		)
		switch msg {
		case ROLL:
			g.takeTurn(gamestate)
		case BANK:
			gamestate.Players[gamestate.CurrentPlayer].Score += gamestate.CurrentScore
			g.passTurn(gamestate)
			g.takeTurn(gamestate)
		}
	}
	println("WINNER!")
}
