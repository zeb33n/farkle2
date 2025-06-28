package main

import (
	"github.com/zeb33n/farkle2/game"
	"github.com/zeb33n/farkle2/tui"
)

func main() {
	tui.TuiInit()
	game.RunGame([]string{"Zeb", "Will"}, 10000)
	tui.TuiClose()
}
