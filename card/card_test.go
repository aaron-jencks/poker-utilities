package card

import (
	"math/rand"
	"sort"
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
		assert.Equal(t, tc.s, ParsePokerCardString(tc.s).String())
	}
}

func TestSortingCards(t *testing.T) {
	cards := []string{"2c", "7h", "as"}
	ocards := make([]Card, len(cards))
	for ci, c := range cards {
		ocards[ci] = ParsePokerCardString(c)
	}
	rand.Shuffle(len(ocards), func(i, j int) { ocards[i], ocards[j] = ocards[j], ocards[i] })
	sort.Slice(ocards, func(i, j int) bool { return ocards[i].LessThan(ocards[j]) })
	for ci := range cards {
		assert.Equal(t, cards[ci], ocards[ci].String())
	}
}
