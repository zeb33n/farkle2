// Package tui
package tui

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"slices"
	"strings"

	"github.com/zeb33n/farkle2/state"
)

var lenLastRender int = 0

type colourKind int

const (
	RESET colourKind = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	GRAY
	WHITE
)

var COLOURS = map[colourKind]string{
	RESET:   "\033[0m",
	RED:     "\033[31m",
	GREEN:   "\033[32m",
	YELLOW:  "\033[33m",
	BLUE:    "\033[34m",
	MAGENTA: "\033[35m",
	CYAN:    "\033[36m",
	GRAY:    "\033[37m",
	WHITE:   "\033[97m",
}

func setStringColour(s string, colour colourKind) string {
	return COLOURS[colour] + s + COLOURS[RESET]
}

func TuiRenderGamestate(gamestate *state.GameState) {
	diceSides := []string{"[.]", "[:]", "[.:]", "[::]", "[:.:]", "[:::]"}
	roll := ""
	for i, e := range gamestate.Dice {
		if slices.Contains(gamestate.ScoringDice, i) {
			roll += setStringColour(diceSides[e-1], RED)
		} else {
			roll += diceSides[e-1]
		}
	}
	players := ""
	for _, e := range gamestate.Players {
		players += fmt.Sprintf("%s: %d\n", e.Name, e.Score)
	}
	gameString := fmt.Sprintf(`
%s
Player: %s
RollScore: %d CurrentScore: %d
%s
controls: [r] roll [b] bank
`,
		roll,
		gamestate.Players[gamestate.CurrentPlayer].Name,
		gamestate.RoundScore,
		gamestate.CurrentScore,
		players,
	)
	renderString(gameString)
}

func TuiRenderTurnChange(player string) {
	gameString := fmt.Sprintf(`
%s's turn! 
controls: [r] roll [b] bank
`,
		player,
	)
	renderString(gameString)
}

func TuiRenderWelcomeLocal(splayers []string) {
	players := ""
	for _, e := range splayers {
		players += fmt.Sprintf("%s\n", e)
	}
	welcomeString := fmt.Sprintf(`
%s
Press [.] to start
EnterName: 
`,
		players,
	)
	renderString(welcomeString)
	// exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func renderString(s string) {
	for i := 0; i < lenLastRender; i++ {
		fmt.Printf("\033[1A\033[2K")
	}
	lenLastRender = strings.Count(s, "\n")
	fmt.Print(s)
}

func handleSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			os.Exit(1)
		}
	}()
}

func TuiInit() {
	handleSigInt()
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Do not display entered characters
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func TuiClose() {
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}
