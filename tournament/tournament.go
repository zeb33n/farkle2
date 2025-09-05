// Package tournament
package tournament

import (
	// "encoding/json"
	"fmt"
	"log"
	// "log"
	// "os"

	"github.com/zeb33n/farkle2/core"
)

type LocalOptions struct {
	Bots   bool
	BestOf int
	Score  int
}

type botStats struct {
	LatestScore int
	TotalScore  int
	TotalBanks  int
	TotalGames  int
}

type ioTournament struct {
	handlers map[string]*core.BotHandler
	stats    map[string]*botStats
}

func (io *ioTournament) AwaitInputPlayer(name string, gs *core.GameState) core.MsgTypeC {
	return io.handlers[name].GetResponse(gs)
}

func (io *ioTournament) updateStats(gs *core.GameState) {
	for _, player := range gs.Players {
		if io.stats[player.Name].LatestScore < player.Score {
			io.stats[player.Name].LatestScore = player.Score
			io.stats[player.Name].TotalBanks += 1
		} else if io.stats[player.Name].LatestScore > player.Score {
			io.stats[player.Name].TotalScore += io.stats[player.Name].LatestScore
			io.stats[player.Name].LatestScore = player.Score
			io.stats[player.Name].TotalGames += 1
		}
	}
}

// func (io *ioTournament) WriteGame(gs *core.GameState) {
// 	gss, err := json.Marshal(gs)
// 	if err != nil {
// 		log.Fatal("couldnt encode gs")
// 	}
// 	err = os.WriteFile(name, gss, 0o666)
// 	if err != nil {
// 		log.Fatal("couldnt write to file")
// 	}
// }

func (io *ioTournament) OutputGamestate(gs *core.GameState) {
	io.updateStats(gs)
	// TODO log to file for replays
}

func (io *ioTournament) OutputTurnChange(gs *core.GameState) {
	io.updateStats(gs)
	// TODO log to file for replays
}

func getStatsString(s *botStats) string {
	return fmt.Sprintf(
		"avg score: %d\navg bank: %d\ngames played: %d\n",
		s.TotalScore/s.TotalGames,
		s.TotalScore/s.TotalBanks,
		s.TotalGames,
	)
}

// TODO bot stats struct and msg

func TournamentRun() {
	bots := core.CONFIG.BotNames
	fmt.Printf("%v\n", bots)
	if len(bots) <= 1 {
		log.Fatal("Need at least 2 bots please add some in your config file")
	}
	handlers := map[string]*core.BotHandler{}
	stats := map[string]*botStats{}
	for _, name := range bots {
		handlers[name] = &core.BotHandler{Name: name}
		handlers[name].Start()
		defer handlers[name].Stop()
		stats[name] = &botStats{0, 0, 0, 0}
	}

	io := ioTournament{handlers: handlers, stats: stats}
	core.TuiRenderTournament(bots, 0, "")
	var roundNum int
	for roundNum = 1; len(bots) > 1; roundNum++ {
		resultsRound := map[string]int{}
		for _, botName := range bots {
			resultsRound[botName] = 0
		}

		winners := make([]string, len(bots)/2)
		for i := 0; i < len(bots); i += 2 {
			if i+1 == len(bots) {
				winners = append(winners, bots[i])
				continue
			}
			var winner string
			for resultsRound[bots[i]] < core.CONFIG.FirstTo &&
				resultsRound[bots[i+1]] < core.CONFIG.FirstTo {
				game := core.Game{IO: &io}
				winner = game.RunGame(
					&map[string]bool{bots[i]: true, bots[i+1]: true},
					core.CONFIG.FinalScore,
				)
				resultsRound[winner] += 1
				core.TuiRenderTournament(
					bots,
					roundNum,
					fmt.Sprintf(
						"CurrentMatch\n%s wins: %d\n%s wins: %d\n",
						bots[i],
						resultsRound[bots[i]],
						bots[i+1],
						resultsRound[bots[i+1]],
					),
				)
			}
			winners[i/2] = winner
		}
		bots = winners
	}
	statsString := ""
	for name, stats := range io.stats {
		statsString += fmt.Sprintf("%s:\n%s\n", name, getStatsString(stats))
	}
	core.TuiRenderTournament(
		bots,
		roundNum,
		fmt.Sprintf("WINNER: %s\n\n%s", bots[0], statsString),
	)
}
