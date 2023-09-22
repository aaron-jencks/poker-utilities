package hand

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/aaron-jencks/poker/card"
	"github.com/stretchr/testify/assert"
)

func TestHandParsing(t *testing.T) {
	tcs := []struct {
		name  string
		shand string
		hand  Hand
	}{
		{
			name:  "high card",
			shand: "2d5s6sjhac",
			hand: Hand{
				Hand:    HIGH_CARD,
				Kicker0: card.ACE,
				Kicker1: card.JACK,
				Contents: []card.Card{
					card.ParsePokerCardString("2d"),
					card.ParsePokerCardString("5s"),
					card.ParsePokerCardString("6s"),
					card.ParsePokerCardString("jh"),
					card.ParsePokerCardString("ac"),
				},
			},
		},
		{
			name:  "pair",
			shand: "2c5h5sjcad",
			hand: Hand{
				Hand:    PAIR,
				Kicker0: card.FIVE,
				Kicker1: card.ACE,
				Contents: []card.Card{
					card.ParsePokerCardString("2c"),
					card.ParsePokerCardString("5h"),
					card.ParsePokerCardString("5s"),
					card.ParsePokerCardString("jc"),
					card.ParsePokerCardString("ad"),
				},
			},
		},
		{
			name:  "pair highest card",
			shand: "2c3h5sacad",
			hand: Hand{
				Hand:    PAIR,
				Kicker0: card.ACE,
				Kicker1: card.FIVE,
				Contents: []card.Card{
					card.ParsePokerCardString("2c"),
					card.ParsePokerCardString("3h"),
					card.ParsePokerCardString("5s"),
					card.ParsePokerCardString("ac"),
					card.ParsePokerCardString("ad"),
				},
			},
		},
		{
			name:  "flush",
			shand: "2c4c7ctcqc",
			hand: Hand{
				Hand:    FLUSH,
				Kicker0: card.QUEEN,
				Kicker1: card.TEN,
				Contents: []card.Card{
					card.ParsePokerCardString("2c"),
					card.ParsePokerCardString("4c"),
					card.ParsePokerCardString("7c"),
					card.ParsePokerCardString("tc"),
					card.ParsePokerCardString("qc"),
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			h := ParsePokerHandString(tc.shand)
			assert.Equal(tt, tc.hand, h, "parsed hands should be equal")
		})
	}
}

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

func TestLessThan(t *testing.T) {
	tcs := []struct {
		name string
		h1   string
		h2   string
		h1lt bool
		eq   bool
	}{
		{
			name: "flush last card",
			h1:   "2c4c7ctcqc",
			h2:   "3h4h7hthqh",
			h1lt: true,
		},
		{
			name: "flush first card",
			h1:   "2c4c7ctcqc",
			h2:   "2h4h7hthkh",
			h1lt: true,
		},
		{
			name: "flush equal",
			h1:   "2c4c7ctcqc",
			h2:   "2h4h7hthqh",
			eq:   true,
		},
		{
			name: "flush greater",
			h2:   "2c4c7ctcqc",
			h1:   "3h4h7hthqh",
		},
		{
			name: "straight equal",
			h1:   "7c8h9dthjs",
			h2:   "7s8s9dthjs",
			eq:   true,
		},
		{
			name: "straight less",
			h1:   "7c8h9dthjs",
			h2:   "8s9dthjsqs",
			h1lt: true,
		},
		{
			name: "straight greater",
			h2:   "7c8h9dthjs",
			h1:   "8s9dthjsqs",
		},
		{
			name: "straight ace low less",
			h1:   "ac2h3d4h5s",
			h2:   "3d4h5s6s7s",
			h1lt: true,
		},
		{
			name: "straight ace low equal",
			h1:   "ac2h3d4h5s",
			h2:   "as2s3d4h5s",
			eq:   true,
		},
		{
			name: "straight ace low greater",
			h2:   "ac2h3d4h5s",
			h1:   "3d4h5s6s7s",
		},
		{
			name: "high card less",
			h1:   "2d5s6sjhkc",
			h2:   "2s5s6sjhas",
			h1lt: true,
		},
		{
			name: "high card less kicker 1",
			h1:   "2d5s6sjhkc",
			h2:   "2s5s6sqskc",
			h1lt: true,
		},
		{
			name: "high card less kicker 3",
			h1:   "2d5s7sjhkc",
			h2:   "2s6s7sjhkc",
			h1lt: true,
		},
		{
			name: "high card equal",
			h1:   "2d5s6sjhkc",
			h2:   "2s5s6sjhks",
			eq:   true,
		},
		{
			name: "high card less kicker 3",
			h2:   "2d5s7sjhkc",
			h1:   "2s6s7sjhkc",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			h1 := ParsePokerHandString(tc.h1)
			h2 := ParsePokerHandString(tc.h2)
			if tc.h1lt {
				assert.True(tt, h1.LessThan(h2), "h1 should be less than h2")
				assert.False(tt, h1.Equals(h2), "h1 should not be equal to h2")
			} else if tc.eq {
				assert.True(tt, h1.Equals(h2), "h1 should be equal to h2")
				assert.False(tt, h1.LessThan(h2), "h1 should not be less than h2")
			} else {
				assert.False(tt, h1.LessThan(h2), "h1 should not be less than h2")
				assert.False(tt, h1.Equals(h2), "h1 should not be equal to h2")
			}
		})
	}
}
