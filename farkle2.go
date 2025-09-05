package main

import (
	"github.com/zeb33n/farkle2/cli"
	"github.com/zeb33n/farkle2/client"
	"github.com/zeb33n/farkle2/core"
	"github.com/zeb33n/farkle2/local"
	"github.com/zeb33n/farkle2/server"
	"github.com/zeb33n/farkle2/tournament"
)

func main() {
	// load the config
	core.CONFIG.LoadConfig(core.CONFIGFILE)

	// flags
	var botFlag cli.FlagBool
	botFlag.Name = "-B"
	botFlag.Help = "Whether or not to play with bots"
	botFlag.Value = &local.LOCALOPTIONS.Bots

	var botNamesFlag cli.FlagStringArray
	botNamesFlag.Name = "-b"
	botNamesFlag.Help = "names of bots to play with"
	botNamesFlag.Value = &core.CONFIG.BotNames

	var finalScoreFlag cli.FlagInt
	finalScoreFlag.Name = "-s"
	finalScoreFlag.Help = "the score required to win"
	finalScoreFlag.Value = &core.CONFIG.FinalScore

	var firstToFlag cli.FlagInt
	firstToFlag.Name = "-t"
	firstToFlag.Help = "how many games a bot has to win to win a tournament set"
	firstToFlag.Value = &core.CONFIG.FirstTo

	// commands
	var localCmd cli.Command
	localCmd.Name = "local"
	localCmd.Help = "play a game localy against friends."
	localCmd.Run = local.LocalRun
	localCmd.Flags = &[]cli.Flag{&botFlag, &finalScoreFlag, &botNamesFlag}

	var serverCmd cli.Command
	serverCmd.Name = "server"
	serverCmd.Help = "start a game server"
	serverCmd.Run = server.ServerRun

	var tournamentCmd cli.Command
	tournamentCmd.Name = "tournament"
	tournamentCmd.Help = "start a game tournament"
	tournamentCmd.Run = tournament.TournamentRun
	tournamentCmd.Flags = &[]cli.Flag{&finalScoreFlag, &firstToFlag, &botNamesFlag}

	var clientCmd cli.Command
	clientCmd.Name = "client"
	clientCmd.Help = "connect to the server over a unix socket"
	clientCmd.Run = client.ClientRun

	// app
	var app cli.App
	app.Name = "Farkle [::]"
	app.Help = `A multiplayer Dice game!
USAGE: farkle2 [command] [options]`
	app.Commands = &[]cli.Command{localCmd, serverCmd, clientCmd, tournamentCmd}

	cli.CliRun(&app)
}
