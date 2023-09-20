package hand

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandRanking(t *testing.T) {
	hands := []string{
		"2d5s6sjhac",
		"2c5h5sjcad",
		"3h3d6c6hks",
		"5sqhqcqdad",
		"7c8h9dthjs",
		"2c4c7ctcqc",
		"8h8d8ckhks",
		"4dahadacas",
		"4h5h6h7h8h",
		"tsjsqsksas",
	}

	ehands := []PokerHands{
		HIGH_CARD,
		PAIR,
		TWO_PAIR,
		THREE_OF_A_KIND,
		STRAIGHT,
		FLUSH,
		FULL_HOUSE,
		FOUR_OF_A_KIND,
		STRAIGHT_FLUSH,
		ROYAL_FLUSH,
	}

	rankings := []string{
		"high card",
		"pair",
		"two pair",
		"three of a kind",
		"straight",
		"flush",
		"full house",
		"four of a kind",
		"straight flush",
		"royal flush",
	}

	shands := make([]Hand, len(hands))
	for hi, h := range hands {
		shands[hi] = ParsePokerHandString(h)
		assert.Equal(t, ehands[hi], shands[hi].Hand,
			"hand: %s, expected %s, found %s",
			h, rankings[hi], rankings[shands[hi].Hand])
	}

	rand.Shuffle(len(shands), func(i, j int) { shands[i], shands[j] = shands[j], shands[i] })

	sort.Slice(shands, func(i, j int) bool { return shands[i].LessThan(shands[j]) })

	for ai := range hands {
		fmt.Println(shands[ai].String())
		assert.Equal(t, hands[ai], shands[ai].String(),
			fmt.Sprintf("expected %s, found %s", rankings[ai], rankings[shands[ai].Hand]))
	}
}
