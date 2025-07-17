package core

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

type WelcomeFromServer struct {
	MessageType string
	Players     []string
}
