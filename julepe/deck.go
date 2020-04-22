package julepe

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Deck []Card

type Suit int

// 3 suits for 5 people
// should be 30 cards total
const (
	Bastos = Suit(iota)
	Oros
	Copas
	// Espadas
)

// we _never_ play with 8 and 9
const (
	Two = iota
	Four
	Five
	Six
	Seven
	Jack
	Knight
	King
	Three
	One
)

// Card struct contains suit, number and
//   if suit is of type trump for the round
type Card struct {
	Suit   Suit
	Number int
	Trump  bool
}

// String method for cards returns the number-suit of the card
func (c Card) String() string {
	return fmt.Sprintf("%s-%s", c.StringCardValue(), c.Suit.String())
}

// SuitString returns the suit as a string
func (s Suit) String() string {
	Suits := []string{"Bastos", "Oros", "Copas"}

	return Suits[s]
}

// CardValueString returns the number as a string, i.e. the actual
// value of the card
func (c Card) StringCardValue() string {
	Numbers := []string{"Two", "Four", "Five", "Six", "Seven", "Ten", "Eleven", "Twelve", "Three", "One"}

	return Numbers[c.Number]
}

// CreateDeck returns a new deck of cards
func CreateDeck(playerCount int) (deck Deck) {
	// Loop over each type and suit appending to the deck
	for i := 0; i < 10; i++ {
		// 3 suits
		for n := 0; n < 3; n++ {
			card := Card{
				Suit:   Suit(n),
				Number: i,
			}
			deck = append(deck, card)
		}
	}

	return deck
}

// Shuffle shuffles a Deck
func (d Deck) Shuffle() Deck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(d); n > 0; n-- {
		randIndex := r.Intn(n)
		// We swap the value at index n-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at n-1 into our
		// unshuffled portion of the slice.
		d[n-1], d[randIndex] = d[randIndex], d[n-1]
	}

	return d
}

// Deal deals a deck to a number of players
func (d Deck) Deal() (deckAfterDeal Deck, players Players) {

	players = make([]Player, 5)

	for i := range d {
		switch {
		case i < 5:
			players[0].Hand = append(players[0].Hand, d[i])
		case i < 10:
			players[1].Hand = append(players[1].Hand, d[i])
		case i < 15:
			players[2].Hand = append(players[2].Hand, d[i])
		case i < 20:
			players[3].Hand = append(players[3].Hand, d[i])
		case i < 25:
			players[4].Hand = append(players[4].Hand, d[i])
		case i == 25:
			log.Println("Breaking...")
			break
		}
	}

	return d[25:], players
}

// Trump returns the trump card for the round
func (d Deck) Trump(players Players) (Deck, Suit) {
	// Player 0 will always be dealer for now...
	players[0].Hand = append(players[0].Hand, d[0])

	log.Println("trump card:", d[0].String())

	return d[1:], d[0].Suit
}
