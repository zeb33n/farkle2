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
	for true {
		tui.TuiRenderWelcomeLocal(splayers)
		var c string
		for true {
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
	game.RunGame(splayers, 10000)
	tui.TuiClose()
}
