// Package client
package client

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/zeb33n/farkle2/core"
)

func reader(r io.Reader) {
	buf := make([]byte, 256)
	for {
		// TODO this works fine for small games
		// need a bigger bufferfor larger games
		// or smaller game state
		_, err := r.Read(buf[:1])
		if err != nil {
			log.Fatal(err)
		}
		n := int(buf[0])
		output := core.Output{}
		n, err = r.Read(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(buf[:n], &output)
		if err != nil {
			log.Fatal(err)
		}
		switch output.Msg {
		case core.WELCOME:
			displayWelcome(output.Content)
		case core.GAMESTATE:
			displayGameState(output.Content)
		case core.TURNCHANGE:
			displayTurnChange(output.Content)
		}
	}
}

func map2GameState(m map[string]any) core.GameState {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	gs := core.GameState{}
	if err := json.Unmarshal(bytes, &gs); err != nil {
		log.Fatal(err)
	}
	return gs
}

func displayTurnChange(content any) {
	m, ok := content.(map[string]any)
	if !ok {
		log.Fatal("couldnt unmarshall turn change")
	}
	gs := map2GameState(m)
	core.TuiRenderTurnChange(&gs)
}

func displayGameState(content any) {
	m, ok := content.(map[string]any)
	if !ok {
		log.Fatal("couldnt unmarshall game state")
	}
	gs := map2GameState(m)
	core.TuiRenderGamestate(&gs)
}

func displayWelcome(content any) {
	players, ok := content.(map[string]any)
	if !ok {
		log.Fatal("couldnt unmarshall welcome")
	}
	splayers := map[string]bool{}
	for k, v := range players {
		splayers[k], ok = v.(bool)
		if !ok {
			log.Fatal("couldnt unmarshall welcome bytes")
		}
	}
	core.TuiRenderWelcomeServer(splayers)
}

func waitForName(c net.Conn) string {
	var name string
	for {
		char := core.WaitForKeyPress(true)
		if char == "\n" {
			break
		}
		name += char
	}
	b, err := json.Marshal(core.Input{PlayerName: name, Msg: core.NAME})
	if err != nil {
		log.Fatal(err)
	}
	c.Write(b)
	return name
}

func waitForReady(c net.Conn, name string) {
	for {
		char := core.WaitForKeyPress(false)
		if char == "." {
			break
		}
	}
	b, err := json.Marshal(core.Input{PlayerName: name, Msg: core.READY})
	if err != nil {
		log.Fatal(err)
	}
	c.Write(b)
}

func playGame(c net.Conn, name string) {
	for {
		char := core.WaitForKeyPress(false)
		var b []byte
		var err error
		switch char {
		case "r":
			b, err = json.Marshal(core.Input{PlayerName: name, Msg: core.ROLL})
		case "b":
			b, err = json.Marshal(core.Input{PlayerName: name, Msg: core.BANK})
		default:
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		c.Write(b)
	}
}

func ClientRun() {
	core.TuiInit()
	defer core.TuiClose()
	// TODO need to get list of players from server on start
	core.TuiRenderWelcomeLocal([]string{})
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	go reader(c)
	name := waitForName(c)
	waitForReady(c, name)
	playGame(c, name)
}
