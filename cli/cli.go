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
	getOptions() []Option
}

type Info struct {
	Name string
	Help string
}

func (i *Info) getName() string { return i.Name }
func (i *Info) getHelp() string { return i.Help }

type App struct {
	Info
	Commands *[]Command
}

func (a *App) getOptions() []Option {
	out := []Option{}
	for _, c := range *a.Commands {
		out = append(out, &c)
	}
	return out
}

type Command struct {
	Info
	Run   func()
	Flags *[]Flag
}

func (c *Command) getOptions() []Option {
	out := []Option{}
	if c.Flags == nil {
		return nil
	}
	for _, f := range *c.Flags {
		out = append(out, f)
	}
	return out
}

type flagKind int

const (
	BOOL flagKind = iota
	STRING
)

type Flag interface {
	Option
	parse(string)
	getKind() flagKind
}

type FlagBool struct {
	Info
	Value *bool
}

func (f *FlagBool) parse(s string)       { *f.Value = true }
func (f *FlagBool) getKind() flagKind    { return BOOL }
func (f *FlagBool) getOptions() []Option { return []Option{f} }

type FlagString struct {
	Info
	Value *string
}

func (f *FlagString) parse(s string)       { *f.Value = s }
func (f *FlagString) getKind() flagKind    { return STRING }
func (f *FlagString) getOptions() []Option { return []Option{f} }

func help(o Option) {
	fmt.Printf("NAME: %s\nINFO: %s\n", o.getName(), o.getHelp())
	fmt.Print("OPTIONS:")
	if o.getOptions() == nil {
		fmt.Println(" none")
	} else {
		for _, f := range o.getOptions() {
			fmt.Printf("\n    %s: %s", f.getName(), f.getHelp())
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkOption(name string, opts []Option) Option {
	if opts == nil {
		return nil
	}
	for _, o := range opts {
		if o.getName() == name {
			return o
		}
	}
	return nil
}

func CliRun(app *App) {
	var helpCmd Command
	helpCmd.Name = "help"
	helpCmd.Help = "display help message"
	helpCmd.Run = func() { help(app) }
	*app.Commands = append(*app.Commands, helpCmd)
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
	i := 1
	for i <= len(args[1:]) {
		flag, ok := checkOption(args[i], cmd.getOptions()).(Flag)
		if flag == nil || !ok {
			help(cmd)
			log.Fatalf("UNKNOWN FLAG %s FOR %s", args[i], cmd.getName())
		}
		switch flag.getKind() {
		case BOOL:
			flag.parse("")
		case STRING:
			i++
			flag.parse(args[i])
		}
		i++
	}
	cmd.Run()
}
