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
	buf := make([]byte, 1024)
	for {
		gs := core.Output{}
		n, err := r.Read(buf[:])
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(buf[:n], &gs)
		switch gs.Msg {
		case core.WELCOME:
			players, ok := gs.Content.(map[string]bool)
			if !ok {
				log.Fatal("couldnt unmarshall welcome")
			}
			core.TuiRenderWelcomeLocal()
		case core.GAMESTATE:
			// yadder
		case core.TURNCHANGE:
			// yaddre
		}

	}
}

func ClientRun() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	go reader(c)
}
