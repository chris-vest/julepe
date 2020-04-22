package julepe

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Hand     []Card
	IsDealer bool
	Playing  bool
	Wallet   float32
}

type Players []Player

type Hand []Card

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

func (p Players) PlayRound() Players {
	for i := range p {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Player %v. Your hand: %s", i, p[i].Hand)
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

func (p Player) AddToPot() Player {
	// make sure the player can afford it somewhere
	p.Wallet = p.Wallet - 0.5
	return p
}
