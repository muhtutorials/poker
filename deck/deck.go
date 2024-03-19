package deck

import (
	"fmt"
	"math/rand"
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Spades Suit = iota
	Clubs
	Diamonds
	Hearts
)

type Rank int

func (r Rank) String() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		panic("invalid card rank")
	}
}

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	default:
		panic("invalid card suit")
	}
}

func (s Suit) UnicodeSymbol() string {
	switch s {
	case Spades:
		return "♠"
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	default:
		panic("invalid card suit")
	}
}

type Card struct {
	Value Rank
	Suit  Suit
}

func NewCard(v Rank, s Suit) Card {
	if v > 13 {
		panic("the rank of the card cannot be higher than 13")
	}

	if s > 3 {
		panic("the suit of the card cannot be higher than 3")
	}

	return Card{
		Value: v,
		Suit:  s,
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s %s", c.Value, c.Suit, c.Suit.UnicodeSymbol())
}

type Deck [52]Card

func New() Deck {
	var (
		deck   Deck
		nRanks = 13
		nSuits = 4
	)

	x := 0
	for i := 0; i < nRanks; i++ {
		for j := 0; j < nSuits; j++ {
			deck[x] = NewCard(Rank(i+1), Suit(j))
			x++
		}
	}

	deck.shuffle()

	return deck
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
