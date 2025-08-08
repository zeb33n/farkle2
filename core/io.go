package core

import (
	"log"
	"net"
)

type inputOutput interface {
	AwaitInput() Input
	AwaitInputPlayer(string) MsgTypeC
	OutputGamestate(*GameState)
	OutputTurnChange(*GameState)
	OutputWelcome(*map[string]bool)
}

type MsgTypeC int

const (
	NAME MsgTypeC = iota
	READY
	BANK
	ROLL
	UNREADY
)

type MsgTypeS int

const (
	WELCOME MsgTypeS = iota
	TURNCHANGE
	GAMESTATE
)

type Output struct {
	Msg     MsgTypeS
	Content any
}

type Input struct {
	PlayerName string
	Msg        MsgTypeC
}

func SockRead(c net.Conn) []byte {
	// msgs max 1kb long
	buf := make([]byte, 1024)
	_, err := c.Read(buf[:1])
	if err != nil {
		log.Fatal(err)
	}
	n := int(buf[0])
	n, err = c.Read(buf[:n])
	if err != nil {
		log.Fatal(err)
	}
	return buf[:n]
}

func SockWrite(msg []byte, c net.Conn) {
	msg = append([]byte{byte(len(msg))}, msg...)
	_, err := c.Write(msg)
	if err != nil {
		log.Fatal("ERROR: writing to socket", err)
	}
}
