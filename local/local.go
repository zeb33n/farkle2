// Package local
package local

import (
	"github.com/zeb33n/farkle2/core"
)

type ioLocal struct{}

func (*ioLocal) AwaitInput() core.Input {
	for {
		switch core.WaitForKeyPress(false) {
		case "r":
			return core.Input{PlayerName: "", Msg: core.ROLL}
		case "b":
			return core.Input{PlayerName: "", Msg: core.BANK}
		default:
			continue
		}
	}
}

func (io *ioLocal) AwaitInputPlayer(_ string) core.MsgTypeC {
	return io.AwaitInput().Msg
}

func (*ioLocal) OutputGamestate(gs *core.GameState) {
	core.TuiRenderGamestate(gs)
}

func (*ioLocal) OutputTurnChange(gs *core.GameState) {
	core.TuiRenderTurnChange(gs)
}

func (*ioLocal) OutputWelcome(names *map[string]bool) {
	players := []string{}
	for k := range *names {
		players = append(players, k)
	}
	core.TuiRenderWelcomeLocal(players)
}

func LocalRun() {
	ioHandler := ioLocal{}
	core.TuiInit()

	splayers := map[string]bool{}
	name := ""
	for {
		ioHandler.OutputWelcome(&splayers)
		var c string
		for {
			c = core.WaitForKeyPress(true)
			if c == "\n" || c == "." {
				break
			}
			name += c
		}
		if c == "." {
			break
		}
		splayers[name] = true
		name = ""
	}
	game := core.Game{IO: &ioLocal{}}
	game.RunGame(&splayers, 10000)
	core.TuiClose()
}
