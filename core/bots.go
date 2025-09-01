package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type BotHandler struct {
	Name string
	proc *exec.Cmd
	in   io.WriteCloser
	out  io.ReadCloser
}

func (b *BotHandler) Start() {
	cmd := fmt.Sprintf("docker run -i --rm %s", b.Name)
	b.proc = exec.Command("bash", "-c", cmd)
	var err error
	b.in, err = b.proc.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	b.out, err = b.proc.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := b.proc.Start(); err != nil {
		log.Fatal(err)
	}
}

func (b *BotHandler) GetResponse(gs *GameState) MsgTypeC {
	gsb, err := json.Marshal(gs)
	if err != nil {
		log.Fatal(err)
	}
	gsb = append(gsb, '\n')
	b.in.Write(gsb)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1)
	n, err := b.out.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	switch string(buf[:n]) {
	case "b":
		return BANK
	case "r":
		return ROLL
	default:
		return BANK
	}
}

func (b *BotHandler) Stop() {
	err := b.in.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = b.out.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = b.proc.Process.Signal(os.Interrupt)
	if err != nil {
		err = b.proc.Process.Kill()
		if err != nil {
			fmt.Printf("COULDNT STOP BOT %s", b.Name)
		}
	}
}
