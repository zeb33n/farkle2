package core

type inputOutput interface {
	AwaitInput() Input
	AwaitInputPlayer(string) MsgType
	OutputGamestate(*GameState)
	OutputTurnChange(string)
	OutputWelcome([]string)
}

type MsgType int

const (
	NAME MsgType = iota
	READY
	BANK
	ROLL
	UNREADY
)

type Input struct {
	PlayerName string
	Msg        MsgType
}

type TurnChange struct {
	Name string
}

type WelcomeFromServer struct {
	MessageType string
	Players     []string
}
