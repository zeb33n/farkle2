package core

type inputOutput interface {
	AwaitInput() Input
	AwaitInputPlayer(string) MsgTypeC
	OutputGamestate(*GameState)
	OutputTurnChange(string)
	OutputWelcome([]string)
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
	MsgType MsgTypeS
	Msg     any
}

type Input struct {
	PlayerName string
	Msg        MsgTypeC
}

type WelcomeFromServer struct {
	MessageType string
	Players     []string
}
