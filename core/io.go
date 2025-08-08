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

// frontmatterLen number of bytes in the frontmatter
const frontmatterLen uint = 4

func bytes2Uint(bs []byte) uint {
	var out uint
	for _, b := range bs {
		if b > 9 {
			log.Fatal("ERROR: cant encode byte as uint")
		}
		out = out*10 + uint(b)
	}
	return out
}

func uint2Bytes(ui uint) []byte {
	if ui >= frontmatterLen*256 {
		log.Fatal("ERROR: message is too long")
	}
	out := make([]byte, frontmatterLen)
	i := len(out) - 1
	for ui != 0 {
		out[i] = byte(ui % 10)
		ui /= 10
		i--
	}
	return out
}

func SockRead(c net.Conn) []byte {
	buf := make([]byte, frontmatterLen*256)
	bs := []byte{}
	for range frontmatterLen {
		_, err := c.Read(buf[:1])
		if err != nil {
			log.Fatal(err)
		}
		bs = append(bs, buf[0])
	}
	n := bytes2Uint(bs)
	_, err := c.Read(buf[:n])
	if err != nil {
		log.Fatal(err)
	}
	return buf[:n]
}

func SockWrite(msg []byte, c net.Conn) {
	msg = append(uint2Bytes(uint(len(msg))), msg...)
	_, err := c.Write(msg)
	if err != nil {
		log.Fatal("ERROR: writing to socket", err)
	}
}
