// Package core
package core

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"slices"
	"strings"
	"unicode/utf8"
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
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+utf8.RuneCountInString(s))/2, s))
}

func nextSquare(i int) int {
	ui := uint32(i)
	// bit twiddling fun
	ui--
	ui |= ui >> 1
	ui |= ui >> 2
	ui |= ui >> 4
	ui |= ui >> 8
	ui |= ui >> 16
	ui++
	return int(ui)
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

// TOD) better but still looks janky for 6 players e.g maybe include byes in
// tornament logic

func writeBotNames(round []string, width int) string {
	line := ""
	lenRoundsquare := nextSquare(len(round))
	for i := 0; i < lenRoundsquare; i += 2 {
		if len(round) == 1 {
			line += centerString(round[0], width/lenRoundsquare)
		} else if len(round) > i+1 {
			line += centerString(round[i], width/lenRoundsquare)
			line += centerString(round[i+1], width/lenRoundsquare)
		} else {
			line += strings.Repeat(" ", (width/lenRoundsquare)*2)
		}
		// TODO colour the strings
	}
	return centerString(line, width) + "\n"
}

func writePipes(round []string, width int) string {
	pipe := ""
	lenRoundsquare := nextSquare(len(round))
	for i := 0; i < lenRoundsquare; i += 2 {
		p := ""
		if len(round) == 1 {
			break
		} else if len(round) > i+1 {
			p = fmt.Sprintf(
				"┗%s┳%s┛",
				strings.Repeat("━", width/(lenRoundsquare*2)-1),
				strings.Repeat("━", width/(lenRoundsquare*2)-1),
			)
		} else {
			p = strings.Repeat(" ", (width/lenRoundsquare)*2)
		}
		pipe += centerString(p, (width/lenRoundsquare)*2)
	}
	return centerString(pipe, width) + "\n"
}

func TuiRenderTournament(players []string) {
	tournamentbracket = append(tournamentbracket, players)
	lenMax := 0
	for _, player := range tournamentbracket[0] {
		if l := utf8.RuneCountInString(player); l > lenMax {
			lenMax = l
		}
	}
	width := nextSquare(len(tournamentbracket[0])) * lenMax
	out := ""
	for _, round := range tournamentbracket {
		out += writeBotNames(round, width) + writePipes(round, width)
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
