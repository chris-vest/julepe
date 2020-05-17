package julepe

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

// Player struct describes player attributes
type Player struct {
	Hand     []Card
	IsDealer bool
	Playing  bool
	Wallet   float32
}

// Players is a slice of Player
type Players []Player

// SelectDealer selects the new dealer at the beginning of the game
func (p Players) SelectDealer() Players {
	p[0].IsDealer = true
	return p
}

// IncrementDealer selects the new dealer at the beginning of each round
// func (p Players) IncrementDealer() Players {
// 	for i := range p {
// 		if p[i].IsDealer == true {
// 			// increment player number who is dealer
// 			p[i].IsDealer = false

// 			break
// 		}
// 	}
// }

// Trump updates player hands with trump suit
func (p Players) Trump(TrumpSuit Suit) Players {
	for i := range p {
		for j := range p[i].Hand {
			if p[i].Hand[j].Suit == TrumpSuit {
				p[i].Hand[j].Trump = true
			}
		}
	}

	return p
}

// PlayRound asks players if they want to play
func (p Players) PlayRound() Players {
	for i := range p {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Player %v. Your hand: %s\nDo you want to play? ", i, p[i].Hand)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Go away, you're too dumb to play - and your computer sucks")
			p[i].Playing = false
			continue
		}
		answer = strings.ToLower(answer)
		// definitely check error your scrub
		if strings.HasPrefix(answer, "y") || strings.HasPrefix(answer, "s") {
			p[i].Playing = true
		} else if strings.HasPrefix(answer, "n") {
			p[i].Playing = false
		} else {
			fmt.Println("Go away, you're too dumb to play - you can't even type")
			p[i].Playing = false
		}
	}
	return p
}

// AddToPot adds 50 cents to pot
func (p Player) AddToPot() Player {
	// make sure the player can afford it somewhere
	p.Wallet = p.Wallet - 0.5
	return p
}

// ExchangeCards
func (p Player) ExchangeCards(playerNumber int, table Table) (Player, Table) {
	// ask player which cards to discard
	fmt.Printf("Player %v. Your hand: %s\nWhich cards do you want to discard?", playerNumber, p.Hand)

	var items []string

	for i := range p.Hand {
		items = append(items, p.Hand[i].String())
	}

	var index int
	var indexes []int
	var result string
	var results []string
	var err error

	// Need to add option to say player doesn't wish to
	// exchange more cards
	for i := 0; i < len(p.Hand); i++ {
		prompt := promptui.SelectWithAdd{
			Label:    "Which cards do you wish to discard?",
			Items:    items,
			AddLabel: "Done",
		}

		index, result, err = prompt.Run()

		// If player chooses Done, break
		if index == -1 {
			break
			// Else append to slices
		} else {
			results = append(results, result)
			indexes = append(indexes, index)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Printf("You choose to exchange %s\n", results)

	// fmt.Println("Indexes: ", indexes)

	// discard cards to _separate_ discard pile
	for i := range indexes {
		// Add discarded cards to table
		table.ExchangeDiscards = append(table.ExchangeDiscards, p.Hand[indexes[i]])
		// Remove cards from hand
		// Perhaps need to find a way to move directly from
		// hand --> table.ExchangeDiscards
		p.Hand = append(p.Hand[:indexes[i]], p.Hand[indexes[i]+1:]...)
		// give player two cards from top of remaining deck
		if len(table.Leftovers) != 0 {
			// Add top card from Leftovers to table
			p.Hand = append(p.Hand, table.Leftovers[0])
			// Remove card from Leftovers
			table.Leftovers = table.Leftovers[1:]
			// if no remaining cards, use discard pile
		} else if len(table.Leftovers) == 0 {
			// Add top card from Discards to table
			p.Hand = append(p.Hand, table.Discards[0])
			// Remove card from Discards
			table.Leftovers = table.Discards[1:]
			// if no discards, use the _separate_ discard pile (should never happen...)
		} else {
			// Add top card from ExchangeDiscards to table
			// This should never realistically happen...
			p.Hand = append(p.Hand, table.ExchangeDiscards.Shuffle()[0])
			// Remove card from Leftovers
			table.Leftovers = table.ExchangeDiscards[1:]
		}
	}

	fmt.Printf("Player %v. Your new hand: %s\n", playerNumber, p.Hand)

	return p, table
}

// DealerDiscard discards the extra dealer card
func (p Player) DealerDiscard(playerNumber int, table Table) (Player, Table) {
	// ask player which cards to discard
	fmt.Printf("Player %v, you are the dealer. Your hand: %s\nWhich card do you want to discard?", playerNumber, p.Hand)

	var items []string

	for i := range p.Hand {
		items = append(items, p.Hand[i].String())
	}

	prompt := promptui.SelectWithAdd{
		Label:    "Which card do you wish to discard?",
		Items:    items,
		AddLabel: "Done",
	}

	index, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Printf("You choose to exchange %s\n", result)

	// add card to discard pile
	table.ExchangeDiscards = append(table.ExchangeDiscards, p.Hand[index])

	// discard the card from the hand
	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)

	return p, table
}
