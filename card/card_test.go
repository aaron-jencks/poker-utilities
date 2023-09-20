package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCardString(t *testing.T) {
	tcs := []struct {
		s string
		c Card
	}{
		{"2c", CreateCard(TWO, CLUBS)},
		{"7h", CreateCard(SEVEN, HEARTS)},
		{"as", CreateCard(ACE, SPADES)},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.c, ParsePokerCardString(tc.s))
	}
}
