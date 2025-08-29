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
	bracketString     string
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
	RESET:   "\x1b[0m",
	RED:     "\x1b[31m",
	GREEN:   "\x1b[32m",
	YELLOW:  "\x1b[33m",
	BLUE:    "\x1b[34m",
	MAGENTA: "\x1b[35m",
	CYAN:    "\x1b[36m",
	GRAY:    "\x1b[37m",
	WHITE:   "\x1b[97m",
}

func setStringColour(s string, colour colourKind) string {
	return COLOURS[colour] + s + COLOURS[RESET]
}

func centerString(s string, w int) string {
	sLen := strings.Clone(s)
	for _, colour := range COLOURS {
		sLen = strings.ReplaceAll(sLen, colour, "")
	}
	paddingl := strings.Repeat(" ", (w-utf8.RuneCountInString(sLen))/2)
	paddingr := strings.Repeat(" ", (w-utf8.RuneCountInString(sLen)+1)/2)
	return paddingl + s + paddingr
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

// TODO refactor this janky af function

func makeBracketString() string {
	lenMax := 0
	for _, player := range tournamentbracket[0] {
		if l := utf8.RuneCountInString(player); l > lenMax {
			lenMax = l
		}
	}
	width := nextSquare(len(tournamentbracket[0])) * lenMax
	out := ""
	for i, round := range tournamentbracket {
		line := ""
		// print the player names
		for _, player := range round {
			// set colours
			if player == "BYE" {
				player = setStringColour(player, BLUE)
			} else if len(round) == 1 {
				player = setStringColour(player, GREEN)
			} else if i < len(tournamentbracket)-1 {
				if slices.Contains(tournamentbracket[i+1], player) {
					player = setStringColour(player, GREEN)
				} else {
					player = setStringColour(player, RED)
				}
			}
			// center the strings
			line += centerString(player, width/len(round))
		}
		out += line + "\n"
		line = ""
		// print the pipes
		for j := 0; j < len(round)-1; j += 2 {
			pipe := fmt.Sprintf(
				"┗%s┳%s┛",
				strings.Repeat("━", width/(len(round)*2)-1),
				strings.Repeat("━", width/(len(round)*2)-1),
			)
			line += centerString(pipe, (width / len(round) * 2))
		}
		out += line + "\n"
	}
	return out
}

func TuiRenderTournament(players []string, roundNum int, msg string) {
	if roundNum > len(tournamentbracket) {
		byes := slices.Repeat([]string{"BYE"}, nextSquare(len(players))-len(players))
		players = append(players, byes...)
		tournamentbracket = append(tournamentbracket, players)
		bracketString = makeBracketString()
	}
	renderString(bracketString + "\n" + msg)
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
