package main

import (
	"github.com/zeb33n/farkle2/cli"
	"github.com/zeb33n/farkle2/client"
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
	{
		Name: "client",
		Help: "connect to the server over a unix socket",
		Run:  client.ClientRun,
	},
}

func main() {
	cli.CliRun(&modes)
}
