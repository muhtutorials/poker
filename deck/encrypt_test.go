package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncryptCard(t *testing.T) {
	card := Card{
		Value: Ace,
		Suit:  Spades,
	}

	key := []byte("foobar")

	output, err := EncryptCard(card, key)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(output, ":encrypted")

	decryptedCard, err := DecryptCard(output, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decryptedCard, ":decrypted")

	if !reflect.DeepEqual(card, decryptedCard) {
		t.Errorf("want %v got %v\n", card, decryptedCard)
	}
}
