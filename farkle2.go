package main

import (
	"github.com/zeb33n/farkle2/game"
	"github.com/zeb33n/farkle2/tui"
	"github.com/zeb33n/farkle2/utils"
)

func main() {
	tui.TuiInit()
	var splayers []string
	name := ""
	for {
		tui.TuiRenderWelcomeLocal(splayers)
		var c string
		for {
			c = utils.WaitForKeypress(true)
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
	tui.TuiRenderTurnChange(splayers[0])
	game.RunGame(splayers, 10000)
	tui.TuiClose()
}
