// Package cli
package cli

import (
	"fmt"
	"os"
)

type Mode struct {
	Name string
	Help string
	Run  func()
}

var MODES []Mode

var helpMode = Mode{
	Name: "help",
	Help: "Display This Message.",
	Run:  help,
}

func help() {
	fmt.Println("Farkle! a multiplayer dice game [::]\nUSAGE:\n    farkle2 [mode]\nMODES:")
	for _, mode := range MODES {
		fmt.Printf("    %s: %s\n", mode.Name, mode.Help)
	}
}

func CliRun(modes *[]Mode) {
	MODES = append(MODES, helpMode)
	MODES = append(MODES, *modes...)
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("TOO MANY ARGUMENTS PROVIDED.")
		help()
		return
	}
	if len(args) < 1 {
		fmt.Println("NO ARGUMENTS PROVIDED.")
		help()
		return
	}
	modeName := args[0]
	for _, mode := range MODES {
		if mode.Name == modeName {
			mode.Run()
			return
		}
	}
	fmt.Printf("UNRECOGNISED MODE: %s\n", modeName)
	help()
}
