// Package client
package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"

	"github.com/zeb33n/farkle2/core"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		output := core.Output{}
		n, err := r.Read(buf[:])
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(buf[:n], &output)
		switch output.Msg {
		case core.WELCOME:
			fmt.Printf("%v\n", output.Content)
			fmt.Println(reflect.TypeOf(output.Content))
			players, ok := output.Content.(map[string]any)
			if !ok {
				log.Fatal("couldnt unmarshall welcome")
			}
			splayers := []string{}
			for k := range players {
				splayers = append(splayers, k)
			}
			core.TuiRenderWelcomeLocal(splayers)
		case core.GAMESTATE:
			gs, ok := output.Content.(core.GameState)
			if !ok {
				log.Fatal("couldnt unmarshall gamestate")
			}
			core.TuiRenderGamestate(&gs)
		case core.TURNCHANGE:
			gs, ok := output.Content.(core.GameState)
			if !ok {
				log.Fatal("couldnt unmarshall turnchange")
			}
			core.TuiRenderTurnChange(&gs)
		}
	}
}

func waitForName(c net.Conn) {
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
	go waitForName(c)
	for {
	}
}
