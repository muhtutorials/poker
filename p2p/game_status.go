package p2p

const (
	TexasHoldem GameVariant = iota
	FiveCardStud
	FiveOPoker
	Guts
	Countdown
)

type GameVariant uint32

func (v GameVariant) String() string {
	switch v {
	case TexasHoldem:
		return "Texas Hold'em"
	case FiveCardStud:
		return "Five Card Stud"
	case FiveOPoker:
		return "Five-O Poker"
	case Guts:
		return "Guts"
	case Countdown:
		return "Countdown"
	default:
		return "Unknown game variant"
	}
}

const (
	Connected GameStatus = iota
	AtTable
	ShuffleAndEncrypt
	PreFlop
	Flop
	Turn
	River
)

type GameStatus int32

func (s GameStatus) String() string {
	switch s {
	case Connected:
		return "Connected"
	case AtTable:
		return "At table"
	case ShuffleAndEncrypt:
		return "Shuffle and encrypt"
	case PreFlop:
		return "Preflop"
	case Flop:
		return "Flop"
	case Turn:
		return "Turn"
	case River:
		return "River"
	default:
		return "Unknown game status"
	}
}

const (
	PlayerActionNone PlayerAction = iota
	PlayerActionFold
	PlayerActionCheck
	PlayerActionBet
)

type PlayerAction int32

func (a PlayerAction) String() string {
	switch a {
	case PlayerActionNone:
		return "None"
	case PlayerActionFold:
		return "Fold"
	case PlayerActionCheck:
		return "Check"
	case PlayerActionBet:
		return "Bet"
	default:
		return "Invalid player action"
	}
}
