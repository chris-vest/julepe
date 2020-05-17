package main

import (
	"fmt"
	"log"

	"github.com/chris-vest/julepe/julepe"
)

func main() {
	fmt.Println("Julepe!")

	playerCount := 5

	table := julepe.Table{
		Pot: 0,
	}

	deck := julepe.CreateDeck(playerCount)
	// fmt.Println(deck)

	shuffledDeck := deck.Shuffle()
	fmt.Printf("shuffled deck: %s", shuffledDeck)

	// Deal cards, create players
	remainingDeck, players := deck.Deal()
	fmt.Println(players)

	// Player 1 is dealer
	players.SelectDealer()

	// All players add 0.50
	table = table.PotAddAll(players)

	fmt.Printf("there are: %d cards remaining\n", len(remainingDeck))

	// Get the trump card, update remainingDeck
	remainingDeck, trumpSuit := remainingDeck.Trump(players)

	fmt.Printf("there are: %d cards remaining\n", len(remainingDeck))

	// Update player hands with trump type
	players = players.Trump(trumpSuit)

	// for i := range players {
	// 	log.Printf("player %d: %s", i, players[i].Hand)
	// }

	// Place the remaining deck on the table
	table.Leftovers = remainingDeck
	// Empty remainingDeck
	remainingDeck = nil

	// players decide to play the round, or not
	// ask players if they want to play
	players = players.PlayRound()

	for i := range players {
		if players[i].Playing == true {
			log.Printf("player %d: %s", i, players[i].Hand)
		} else {
			// create discard pile
			// of hands of players not playing the round
			table.Discards = append(table.Discards, players[i].Hand...)
		}
	}

	fmt.Println(players)

	// all players who are playing can discard cards and draw
	// firstly from the leftovers, then the discard pile (shuffled)
	for i := range players {
		if players[i].Playing == true {
			players[i], table = players[i].ExchangeCards(i, table)
		}
		fmt.Println("Player: ", i, "; Table: ", table)
		// dealer discards 6th card so he has 5 cards
		if players[i].IsDealer == true {
			players[i], table = players[i].DealerDiscard(i, table)
			fmt.Println("Player: ", i, "; Table: ", table)
		}
	}

	// dealer player number + 1 starts by playing a card
	// the remaining players play a card

	// continue until no cards left in hands

	// figure out who won - depending on type of win, assess what to do with pot value
}
