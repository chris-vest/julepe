package julepe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	card := Card{
		Suit:   Suit(0),
		Number: 1,
		Trump:  false,
	}

	cardString := card.String()

	require.Equal(t, "Four-Bastos", cardString)
}
