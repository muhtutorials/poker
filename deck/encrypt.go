package deck

import (
	"bytes"
	"encoding/gob"
)

func EncryptCard(card Card, key []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(card); err != nil {
		return nil, err
	}

	return EnDeCrypt(buf.Bytes(), key)
}

func DecryptCard(payload, key []byte) (Card, error) {
	var card Card

	b, err := EnDeCrypt(payload, key)
	if err != nil {
		return card, err
	}

	if err = gob.NewDecoder(bytes.NewReader(b)).Decode(&card); err != nil {
		return card, err
	}

	return card, nil
}

// EnDeCrypt is used for both encryption and decryption
func EnDeCrypt(payload, key []byte) ([]byte, error) {
	output := make([]byte, len(payload))
	for i := 0; i < len(payload); i++ {
		output[i] = payload[i] ^ key[i%len(key)]
	}

	return output, nil
}
