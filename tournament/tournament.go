// Package tournament
package tournament

import (
	"fmt"

	"github.com/zeb33n/farkle2/core"
)

type ioTournament struct{}

func (io *ioTournament) AwaitInputPlayer(name string, gs *core.GameState) core.MsgTypeC {
	return core.BotGetResponse(name, gs)
}

func (io *ioTournament) OutputGamestate(gs *core.GameState) {
	// fmt.Printf("%v\n", gs)
	// log to file for replays
}

func (io *ioTournament) OutputTurnChange(*core.GameState) {
	// log to file for replays
}

func TournamentRun() {
	var conf core.Config
	conf.LoadConfig("config.json")
	io := ioTournament{}
	bots := conf.Bots
	fmt.Printf("%v\n", bots)
	for len(bots) > 1 {
		winners := make([]string, len(bots)/2)
		for i := 0; i < len(bots); i += 2 {
			if i+1 == len(bots) {
				winners = append(winners, bots[i])
				continue
			}
			game := core.Game{IO: &io}
			winner := game.RunGame(
				&map[string]bool{bots[i]: true, bots[i+1]: true},
				conf.FinalScore,
			)
			winners[i/2] = winner
		}
		bots = winners
		fmt.Printf("%v\n", bots)
	}
}
