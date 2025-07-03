// Package local
package local

import (
	"github.com/zeb33n/farkle2/core"
	"github.com/zeb33n/farkle2/utils"
)

func LocalRun() {
	core.TuiInit()
	var splayers []string
	name := ""
	for {
		core.TuiRenderWelcomeLocal(splayers)
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
	core.TuiRenderTurnChange(splayers[0])
	core.RunGame(splayers, 10000)
	core.TuiClose()
}
