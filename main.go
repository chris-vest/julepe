package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type suit int

// 3 suits for 5 people
// should be 30 cards total
const (
	bastos = suit(iota)
	oros
	copas
	// espadas
)

type cardValue int

// we _never_ play with 8 and 9
const (
	two = cardValue(iota)
	four
	five
	six
	seven
	jack
	knight
	king
	three
	one
)

type card struct {
	suit   suit
	number cardValue
	trump  bool
}

// String method for cards returns the number-suit of the card
func (c card) String() string {
	return fmt.Sprintf("%s-%s", CardValue(c.number), Suit(c.suit))
}

// Suit returns the suit as a string
func Suit(suit suit) string {
	suits := []string{"bastos", "oros", "copas", "espadas"}

	return suits[suit]
}

// CardValue returns the number as a string, i.e. the actual
// value of the card
func CardValue(number cardValue) string {
	numbers := []string{"two", "four", "five", "six", "seven", "ten", "eleven", "twelve", "three", "one"}

	return numbers[number]
}

type deck []card

type hand []card

type player struct {
	hand     hand
	isDealer bool
	wallet   int
}

type players []player

type table struct {
	pot      int
	discards []card
}

func createDeck(playerCount int) deck {
	var deck []card

	// Loop over each type and suit appending to the deck
	for i := 0; i < 10; i++ {
		// 3 suits
		for n := 0; n < 3; n++ {
			card := card{
				suit:   suit(n),
				number: cardValue(i),
			}
			deck = append(deck, card)
		}
	}

	fmt.Println(len(deck))

	return deck
}

func (d deck) shuffle() deck {
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

// func (d deck) deal(players int) {
func (d deck) deal() (deckAfterDeal deck, players players) {

	players = make([]player, 5)

	for i := range d {
		switch {
		case i < 5:
			players[0].hand = append(players[0].hand, d[i])
		case i < 10:
			players[1].hand = append(players[1].hand, d[i])
		case i < 15:
			players[2].hand = append(players[2].hand, d[i])
		case i < 20:
			players[3].hand = append(players[3].hand, d[i])
		case i < 25:
			players[4].hand = append(players[4].hand, d[i])
		case i == 25:
			log.Println("Breaking...")
			break
		}
	}

	return d[25:], players
}

func (d deck) trump(players players) (deck, suit) {
	// Player 0 will always be dealer for now...
	players[0].hand = append(players[0].hand, d[0])

	log.Println("trump card:", d[0].String())

	return d[1:], d[0].suit
}

func (p players) trump(trumpSuit suit) players {
	for i := range p {
		for j := range p[i].hand {
			if p[i].hand[j].suit == trumpSuit {
				p[i].hand[j].trump = true
			}
		}
	}

	return p
}

func main() {
	fmt.Println("Julepe!")

	playerCount := 5

	deck := createDeck(playerCount)
	// fmt.Println(deck)

	shuffledDeck := deck.shuffle()
	fmt.Printf("shuffled deck: %s", shuffledDeck)

	remainingDeck, players := deck.deal()
	fmt.Println(players)

	fmt.Printf("there are: %d cards remaining\n", len(remainingDeck))

	// Get the trump card, update remainingDeck
	remainingDeck, trumpSuit := remainingDeck.trump(players)

	fmt.Printf("there are: %d cards remaining\n", len(remainingDeck))

	// Update hands with trump type
	players = players.trump(trumpSuit)

	for i := range players {
		log.Printf("Player %d: %s", i, players[i].hand)
	}
}
