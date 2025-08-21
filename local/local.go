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

type LocalOptions struct {
	Bots   bool
	Config string
}

var LOCALOPTIONS = LocalOptions{
	Bots:   false,
	Config: "config.json",
}

func LocalRun() {
	var config core.Config
	config.LoadConfig(LOCALOPTIONS.Config)

	ioHandler := ioLocal{bots: []string{}}
	splayers := map[string]bool{}
	if LOCALOPTIONS.Bots {
		ioHandler = ioLocal{bots: config.Bots}
		for _, botName := range config.Bots {
			splayers[botName] = true
		}
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
	game.RunGame(&splayers, config.FinalScore)
	core.TuiClose()
}
