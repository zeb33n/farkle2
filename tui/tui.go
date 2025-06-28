package tui

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/zeb33n/farkle2/state"
)

var lenLastRender int = 0

func TuiRenderGamestate(gamestate *state.GameState) {
	dice_sides := []string{"[.]", "[:]", "[.:]", "[::]", "[:.:]", "[:::]"}
	roll := ""
	for _, e := range gamestate.Dice {
		roll += dice_sides[e-1]
	}
	players := ""
	for _, e := range gamestate.Players {
		players += fmt.Sprintf("%s: %d\n", e.Name, e.Score)
	}
	game_string := fmt.Sprintf(
		`
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
	renderString(game_string)
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
	exec.Command("clear").Run()
	// tty setup
	renderString("\nPress AnyKey to start.\n")
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Do not display entered characters
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func TuiClose() {
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}
