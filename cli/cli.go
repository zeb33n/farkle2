// Package cli
package cli

import (
	"fmt"
	"log"
	"os"
)

type Option interface {
	getName() string
	getHelp() string
	getOptions() *[]Option
}

type Info struct {
	Name    string
	Help    string
	Options *[]Option
}

func (i *Info) getName() string       { return i.Name }
func (i *Info) getHelp() string       { return i.Help }
func (i *Info) getOptions() *[]Option { return i.Options }

type App struct {
	Info
}

// TODO handle flags better than map string any

type Command struct {
	Info
	Run func(*map[string]any)
}

type Flag struct {
	Info
	Kind    FlagKind
	Default any
}

type FlagKind int

const (
	STRING FlagKind = iota
	BOOL
)

func help(o Option) {
	fmt.Printf("NAME: %s\nINFO: %s\n", o.getName(), o.getHelp())
	fmt.Print("OPTIONS:")
	if o.getOptions() == nil {
		fmt.Println(" none")
	} else {
		for _, f := range *o.getOptions() {
			fmt.Printf("\n    %s: %s", f.getName(), f.getHelp())
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkOption(name string, opts *[]Option) Option {
	if opts == nil {
		return nil
	}
	for _, o := range *opts {
		if o.getName() == name {
			return o
		}
	}
	return nil
}

func getFlagMap(opts *[]Option) map[string]any {
	flags := map[string]any{}
	if opts == nil {
		return nil
	}
	for _, o := range *opts {
		f, ok := o.(*Flag)
		if !ok {
			log.Fatal("THATS NOT A FLAG")
		}
		flags[f.Name] = f.Default
	}
	return flags
}

func CliRun(app *App) {
	var helpCmd Command
	helpCmd.Name = "help"
	helpCmd.Help = "display help message"
	helpCmd.Run = func(_ *map[string]any) { help(app) }
	fullOptions := append(*app.getOptions(), &helpCmd)
	app.Options = &fullOptions
	if len(os.Args) == 1 {
		help(app)
		log.Fatalf("NO COMMAND PROVIDED")
	}
	args := os.Args[1:]
	cmd, ok := checkOption(args[0], app.getOptions()).(*Command)
	if cmd == nil || !ok {
		help(app)
		log.Fatalf("UNKNOWN COMMAND %s", args[0])
	}
	flags := getFlagMap(cmd.getOptions())
	i := 1
	for i <= len(args[1:]) {
		flag, ok := checkOption(args[i], cmd.getOptions()).(*Flag)
		if flag == nil || !ok {
			help(cmd)
			log.Fatalf("UNKNOWN FLAG %s FOR %s", args[i], cmd.getName())
		}
		switch flag.Kind {
		case BOOL:
			flags[flag.Name] = true
		case STRING:
			i++
			flags[flag.Name] = args[i]
		}
		i++
	}
	cmd.Run(&flags)
}
