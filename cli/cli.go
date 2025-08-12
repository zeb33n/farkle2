// Package cli
package cli

import (
	"fmt"
	"log"
	"os"
)

type Mode struct {
	Name string
	Help string
	Run  func(*[]Mode)
	Opts *[]Mode
}

var MODES []Mode

var CLI Mode = Mode{
	Name: "Farkle [::]\n",
	Help: `A multiplayer Dice game!

USAGE: farkle2 [command] [options]

COMMANDS:`,
	Opts: &MODES,
}

// can tidy this up with an interface
func help(m *Mode) {
	fmt.Printf("%s\n%s\n", m.Name, m.Help)
	if m.Opts == nil {
		return
	}
	for _, f := range *m.Opts {
		fmt.Printf("    %s: %s\n", f.Name, f.Help)
	}
	fmt.Println()
}

func CliRun(modes *[]Mode) {
	MODES = append(MODES, *modes...)
	args := os.Args[1:]
	if len(args) < 1 {
		help(&CLI)
		log.Fatal("NO COMMAND PROVIDED.")
	}
	modeName := args[0]
	for _, mode := range MODES {
		if mode.Name == modeName {
			if args[1] == "--help" {
				help(&mode)
				return
			}
			if len(args) > len(*mode.Opts)-1 {
				help(&mode)
				log.Fatal("TOO MANY ARGUMENTS PROVIDED.")
			}
			mode.Run(mode.Opts)
			return
		}
	}
	help(&CLI)
	log.Fatalf("UNRECOGNISED MODE: %s\n", modeName)
}
