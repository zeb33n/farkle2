// Package cli
package cli

import (
	"fmt"
	"log"
	"os"
	"slices"
)

type Mode struct {
	Name string
	Help string
	Run  func(*map[string]bool)
	Opts []Mode
}

func help(m *Mode) {
	fmt.Printf("NAME: %s\nINFO: %s\n", m.Name, m.Help)
	fmt.Print("OPTIONS:")
	if m.Opts == nil {
		fmt.Println(" none")
	} else {
		for _, f := range m.Opts {
			fmt.Printf("\n    %s: %s", f.Name, f.Help)
		}
		fmt.Println()
	}
	fmt.Println()
}

func CliRun(modes *[]Mode) {
	cliMode := Mode{
		Name: "Farkle [::]",
		Help: `A multiplayer Dice game!
USAGE: farkle2 [command] [options]`,
		Opts: *modes,
	}

	args := os.Args[1:]
	if len(args) < 1 {
		help(&cliMode)
		log.Fatal("NO COMMAND PROVIDED.")
	}
	modeName := args[0]
	for _, mode := range *modes {
		if mode.Name == modeName {
			if slices.Contains(args, "--help") {
				help(&mode)
				return
			}
			if len(args) > 1 {
				names := []string{}
				for _, opt := range mode.Opts {
					names = append(names, opt.Name)
				}
				for _, arg := range args[1:] {
					if !slices.Contains(names, arg) {
						help(&mode)
						log.Fatalf("UNKNOWN FLAG %q", arg)
					}
				}
			}
			if len(args)-1 > len(mode.Opts) {
				help(&mode)
				log.Fatal("TOO MANY ARGUMENTS PROVIDED.")
			}
			flags := map[string]bool{}
			for _, opt := range mode.Opts {
				flags[opt.Name] = slices.Contains(args, opt.Name)
			}
			mode.Run(&flags)
			return
		}
	}
	help(&cliMode)
	log.Fatalf("UNRECOGNISED MODE: %s\n", modeName)
}
