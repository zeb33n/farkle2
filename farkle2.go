package main

import (
	"github.com/zeb33n/farkle2/cli"
	"github.com/zeb33n/farkle2/client"
	"github.com/zeb33n/farkle2/local"
	"github.com/zeb33n/farkle2/server"
)

func main() {
	// flags
	var botFlag cli.FlagBool
	botFlag.Name = "-b"
	botFlag.Help = "Play against bots from the `bots` directory"
	botFlag.Value = &local.LOCALOPTIONS.Bots

	var confFlag cli.FlagString
	confFlag.Name = "-c"
	confFlag.Help = "load a config file"
	confFlag.Value = &local.LOCALOPTIONS.Config

	// commands
	var localCmd cli.Command
	localCmd.Name = "local"
	localCmd.Help = "play a game localy against friends."
	localCmd.Run = local.LocalRun
	localCmd.Flags = &[]cli.Flag{&botFlag, &confFlag}

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
	app.Commands = &[]cli.Command{localCmd, serverCmd, clientCmd}

	cli.CliRun(&app)
}
