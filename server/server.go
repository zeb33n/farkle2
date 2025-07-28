// Package server
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/zeb33n/farkle2/core"
)

type ioServer struct {
	in  chan []byte
	out []chan []byte
}

func (io *ioServer) AwaitInput() core.Input {
	// 90% sure this loop is blocking
	for bytes := range io.in {
		var data core.Input
		err := json.Unmarshal(bytes, &data)
		if err != nil {
			fmt.Printf("WARNING: bad json received %v\n", err)
			continue
		}
		return data
	}
	panic("Channel Closed before input received")
}

func (io *ioServer) AwaitInputPlayer(player string) core.MsgTypeC {
	for {
		input := io.AwaitInput()
		if input.PlayerName != player {
			continue
		}
		return input.Msg
	}
}

func (io *ioServer) OutputGamestate(gs *core.GameState) {
	gameBytes := marshallOuput(core.GAMESTATE, gs)
	for _, ch := range io.out {
		ch <- gameBytes
	}
}

func (io *ioServer) OutputTurnChange(gs *core.GameState) {
	turnBytes := marshallOuput(core.TURNCHANGE, gs)
	for _, ch := range io.out {
		ch <- turnBytes
	}
}

func (io *ioServer) OutputWelcome(names *map[string]bool) {
	welcomeBytes := marshallOuput(core.WELCOME, names)
	for _, ch := range io.out {
		ch <- welcomeBytes
	}
}

func (io *ioServer) handleConnection(c net.Conn) {
	outChannel := make(chan []byte)
	io.out = append(io.out, outChannel)
	// Handle messages from clients
	go func() {
		for {
			buf := make([]byte, 256)
			n, err := c.Read(buf)
			if err != nil {
				log.Fatal("ERROR: reading into buffer", err)
			}
			io.in <- buf[:n]
		}
	}()
	// Send messages to clients
	go func() {
		for s := range outChannel {
			s = append([]byte{byte(len(s))}, s...)
			_, err := c.Write(s)
			println(len(s))
			if err != nil {
				fmt.Println("ERROR: writing to socket", err)
			}
		}
	}()
}

func (io *ioServer) serverWelcome() {
	players := map[string]bool{}
	for {
		input := io.AwaitInput()
		switch input.Msg {
		default:
			continue
		case core.READY:
			players[input.PlayerName] = true
		case core.NAME:
			players[input.PlayerName] = false
		}
		for k, v := range players {
			fmt.Printf("%s: %v\n", k, v)
		}
		if allTrue(&players) {
			break
		}
		io.OutputWelcome(&players)
	}
	game := core.Game{IO: io}
	game.RunGame(&players, 10000)
	println("starting game")
}

func allTrue(s *map[string]bool) bool {
	for _, e := range *s {
		if !e {
			return false
		}
	}
	return true
}

func marshallOuput(msg core.MsgTypeS, content any) []byte {
	out := core.Output{Msg: msg, Content: content}
	bytes, err := json.Marshal(out)
	if err != nil {
		log.Fatal("Could not Marshal the gamestate")
	}
	return bytes
}

func ServerRun() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	// listen for connections
	io := ioServer{in: make(chan []byte), out: []chan []byte{}}
	go io.serverWelcome()
	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		io.handleConnection(fd)
	}
}
