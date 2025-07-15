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

func (io *ioLocal) AwaitInputPlayer(_ string) core.MsgType {
	return io.AwaitInput().Msg
}

func (*ioLocal) OutputGamestate(gs *core.GameState) {
	core.TuiRenderGamestate(gs)
}

func (*ioLocal) OutputTurnChange(name string) {
	core.TuiRenderTurnChange(name)
}

func (*ioLocal) OutputWelcome(names []string) {
	core.TuiRenderWelcomeLocal(names)
}

func LocalRun() {
	ioHandler := ioLocal{}
	core.TuiInit()

	var splayers []string
	name := ""
	for {
		ioHandler.OutputWelcome(splayers)
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
		splayers = append(splayers, name)
		name = ""
	}
	ioHandler.OutputTurnChange(splayers[0])
	game := core.Game{IO: &ioLocal{}}
	game.RunGame(splayers, 10000)
	core.TuiClose()
}
