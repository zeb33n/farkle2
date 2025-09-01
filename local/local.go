// Package local
package local

import (
	"time"

	"github.com/zeb33n/farkle2/core"
)

type ioLocal struct {
	bots map[string]*core.BotHandler
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
	if _, ok := io.bots[name]; ok {
		time.Sleep(time.Second)
		return io.bots[name].GetResponse(gs)
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

func (io *ioLocal) AwaitPlayers(names *map[string]bool) {
	name := ""
	for {
		io.OutputWelcome(names)
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
		(*names)[name] = true
		name = ""
	}
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
	ioHandler := ioLocal{bots: map[string]*core.BotHandler{}}
	splayers := map[string]bool{}
	if LOCALOPTIONS.Bots {
		for _, botName := range config.Bots {
			ioHandler.bots[botName] = &core.BotHandler{Name: botName}
			ioHandler.bots[botName].Start()
			defer ioHandler.bots[botName].Stop()
			splayers[botName] = true
		}
	}
	core.TuiInit()
	ioHandler.AwaitPlayers(&splayers)
	game := core.Game{IO: &ioHandler}
	game.RunGame(&splayers, config.FinalScore)
	core.TuiClose()
}
