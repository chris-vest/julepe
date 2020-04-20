package julepe

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Hand     Hand
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
		text, _ := reader.ReadString('\n')
		if strings.Contains("yes", text) {
			p[i].Playing = true
		} else {
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
