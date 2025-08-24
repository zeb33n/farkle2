// Package core
package core

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"slices"
	"strings"
)

var (
	lenLastRender     int        = 0
	tournamentbracket [][]string = [][]string{}
)

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

func centerString(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

func TuiRenderGamestate(gamestate *GameState) {
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

func TuiRenderTurnChange(gs *GameState) {
	players := ""
	for _, e := range gs.Players {
		players += fmt.Sprintf("%s: %d\n", e.Name, e.Score)
	}
	gameString := fmt.Sprintf(`

%s's turn!


%s
controls: [r] roll [b] bank
`,
		gs.Players[gs.CurrentPlayer].Name,
		players,
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
}

func TuiRenderWelcomeServer(splayers map[string]bool) {
	players := ""
	for k, v := range splayers {
		readyness := setStringColour("waiting", RED)
		if v {
			readyness = setStringColour("ready", GREEN)
		}
		players += fmt.Sprintf("%s: %s\n", k, readyness)
	}
	welcomeString := fmt.Sprintf(`
%s
Press [.] to start
EnterName: 
`,
		players,
	)
	renderString(welcomeString)
}

func TuiRenderTournament(players []string) {
	tournamentbracket = append(tournamentbracket, players)
	widthName := 20
	widthLine := len(tournamentbracket[0]) * widthName
	out := ""
	for j, round := range tournamentbracket {
		line := ""
		for _, player := range round {
			player = centerString(player, widthName)
			// Need to ignore colours!
			// if j+1 < len(tournamentbracket) {
			// 	if slices.Contains(tournamentbracket[j+1], player) {
			// 		player = setStringColour(player, GREEN)
			// 	} else {
			// 		player = setStringColour(player, RED)
			// 	}
			// }
			line += player
		}
		out += centerString(line, widthLine) + "\n"
		pipe := ""
		for i := 0; i < len(round)-1; i += 2 {
			pipe += fmt.Sprintf(
				"%s┗%s┳%s┛",
				strings.Repeat(" ", widthName/2),
				strings.Repeat("━", widthName/2-1),
				strings.Repeat("━", widthName/2),
			)
		}
		pipe += "\n"
		pipe = strings.Repeat(" ", widthName*j/2) + pipe
		out += pipe
	}
	renderString(out)
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
			TuiClose()
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
