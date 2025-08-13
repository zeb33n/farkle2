// Package local
package local

import (
	"slices"
	"time"

	"github.com/zeb33n/farkle2/core"
)

type ioLocal struct {
	bots []string
}

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

func (io *ioLocal) AwaitInputPlayer(name string, gs *core.GameState) core.MsgTypeC {
	if slices.Contains(io.bots, name) {
		time.Sleep(time.Second / 2)
		return core.BotGetResponse(name, gs)
	} else {
		return io.AwaitInput().Msg
	}
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

// TODO if the bot flag is set load bots

func LocalRun(flags *map[string]bool) {
	ioHandler := ioLocal{bots: []string{}}
	splayers := map[string]bool{}
	if (*flags)["-b"] {
		ioHandler = ioLocal{bots: []string{"python_example"}}
		splayers = map[string]bool{"python_example": true}
	}
	core.TuiInit()
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
	game := core.Game{IO: &ioHandler}
	game.RunGame(&splayers, 10000)
	core.TuiClose()
}
