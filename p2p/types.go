package p2p

type Broadcast struct {
	To      []string
	Payload any
}

type EncryptedDeck [][]byte
