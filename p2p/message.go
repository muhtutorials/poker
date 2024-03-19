package p2p

type Message struct {
	From    string
	Payload any
}

type Broadcast struct {
	To      []string
	Payload any
}

type Handshake struct {
	GameVariant
	Version string
	GameStatus
	ListenAddr string
}

type PeerList []string

type EncryptedDeck [][]byte

type MessagePlayerAction struct {
	GameStatus
	PlayerAction
	Value int
}
