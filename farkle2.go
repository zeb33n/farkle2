package main

import (
	"github.com/zeb33n/farkle2/cli"
	"github.com/zeb33n/farkle2/local"
)

var modes = []cli.Mode{
	{
		Name: "local",
		Help: "play a game localy against friends.",
		Run:  local.LocalRun,
	},
}

func main() {
	cli.CliRun(&modes)
}
