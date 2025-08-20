package main

import (
	"github.com/zeb33n/farkle2/cli"
	"github.com/zeb33n/farkle2/client"
	"github.com/zeb33n/farkle2/local"
	"github.com/zeb33n/farkle2/server"
)

func main() {
	// flags
	var botFlag cli.Flag
	botFlag.Name = "-b"
	botFlag.Help = "Play against bots from the `bots` directory"
	botFlag.Kind = cli.BOOL
	botFlag.Default = false

	var confFlag cli.Flag
	confFlag.Name = "-c"
	confFlag.Help = "load a config file"
	confFlag.Kind = cli.STRING
	confFlag.Default = ""

	// commands
	var localCmd cli.Command
	localCmd.Name = "local"
	localCmd.Help = "play a game localy against friends."
	localCmd.Run = local.LocalRun
	localCmd.Options = &[]cli.Option{&botFlag, &confFlag}

	var serverCmd cli.Command
	serverCmd.Name = "server"
	serverCmd.Help = "start a game server"
	serverCmd.Run = server.ServerRun

	var clientCmd cli.Command
	clientCmd.Name = "client"
	clientCmd.Help = "connect to the server over a unix socket"
	clientCmd.Run = client.ClientRun

	// app
	var app cli.App
	app.Name = "Farkle [::]"
	app.Help = `A multiplayer Dice game!
USAGE: farkle2 [command] [options]`
	app.Options = &[]cli.Option{&localCmd, &serverCmd, &clientCmd}

	cli.CliRun(&app)
}
