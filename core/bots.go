package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func BotGetResponse(name string, gs *GameState) MsgTypeC {
	gsb, err := json.Marshal(gs)
	cmd := fmt.Sprintf("echo '%s' | docker run -i %s", string(gsb), name)
	if err != nil {
		log.Fatal("Couldnt send gamestate to bot", err)
	}
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal("Couldnt send gamestate to bot", err)
	}
	switch string(out) {
	case "b":
		return BANK
	case "r":
		return ROLL
	default:
		return BANK
	}
}
