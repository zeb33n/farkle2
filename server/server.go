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
	out []chan string
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

func (io *ioServer) AwaitInputPlayer(player string) core.MsgType {
	for {
		input := io.AwaitInput()
		if input.PlayerName != player {
			continue
		}
		return input.Msg
	}
}

func (*ioServer) OutputGamestate(gs *core.GameState) {
	// send output down the channels
	// might need some json serialisarion
}

func (*ioServer) OutputTurnChange(name string) {
	// send output down the channels
}

func (io *ioServer) OutputWelcome(names []string) {
	println("Awaiting Game Start")
	for _, name := range names {
		println(name)
	}
	// send output down the channels
	// might need some json serialisarion
}

func (io *ioServer) handleConnection(c net.Conn) {
	outChannel := make(chan string)
	io.out = append(io.out, outChannel)
	go func() {
		for {
			buf := make([]byte, 512)
			n, err := c.Read(buf)
			if err != nil {
				log.Fatal("ERROR: reading into buffer", err)
			}
			io.in <- buf[:n]
		}
	}()
	go func() {
		for s := range outChannel {
			_, err := c.Write([]byte(s))
			if err != nil {
				fmt.Println("ERROR: writing to socket", err)
			}
		}
	}()
}

func (io *ioServer) ServerWelcome() {
	players := []string{}
	playersIndex := map[string]int{}
	readys := []bool{}
	i := 0
	for {
		input := io.AwaitInput()
		switch input.Msg {
		default:
			continue
		case core.READY:
			readys[playersIndex[input.PlayerName]] = true
		case core.NAME:
			players = append(players, input.PlayerName)
			readys = append(readys, false)
			playersIndex[input.PlayerName] = i
			i++
		}
		for j, playerName := range players {
			fmt.Printf("%s: %v\n", playerName, readys[j])
		}
		if allTrue(readys) {
			break
		}
	}
	game := core.Game{IO: io}
	game.RunGame(players, 10000)
}

func allTrue(s []bool) bool {
	for _, e := range s {
		if !e {
			return false
		}
	}
	return true
}

func ServerRun() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	// listen for connections
	io := ioServer{in: make(chan []byte), out: []chan string{}}
	go io.ServerWelcome()
	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go io.handleConnection(fd)
	}

	// wait for game start
}
