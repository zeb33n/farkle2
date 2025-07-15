package main

import (
	"github.com/zeb33n/farkle2/cli"
	"github.com/zeb33n/farkle2/local"
	"github.com/zeb33n/farkle2/server"
)

var modes = []cli.Mode{
	{
		Name: "local",
		Help: "play a game localy against friends.",
		Run:  local.LocalRun,
	},
	{
		Name: "server",
		Help: "start a game server",
		Run:  server.ServerRun,
	},
}

func main() {
	cli.CliRun(&modes)
}
