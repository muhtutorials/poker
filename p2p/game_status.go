package p2p

const (
	Connected GameStatus = iota
	Ready
	ShuffleAndDeal
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
	case Ready:
		return "Ready"
	case ShuffleAndDeal:
		return "Shuffle and deal"
	case PreFlop:
		return "Pre flop"
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
