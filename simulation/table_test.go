package simulation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aaron-jencks/poker/card"
	"github.com/aaron-jencks/poker/hand"
	"github.com/stretchr/testify/assert"
)

func TestFindBestHand(t *testing.T) {
	tcs := []struct {
		hand  string
		table string
		bh    hand.PokerHands
	}{
		{
			hand:  "4h3d",
			table: "2sts4c2hac",
			bh:    hand.TWO_PAIR,
		},
		{
			hand:  "qc9d",
			table: "6h5ckh8c3h",
			bh:    hand.HIGH_CARD,
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s %s", tc.hand, tc.table), func(tt *testing.T) {
			h := card.ParseMultiPokerCardString(tc.hand)
			tbl := card.ParseMultiPokerCardString(tc.table)
			ah := FindBestHand(h, tbl)
			assert.Equal(tt, tc.bh, ah.Hand, "expected parsed hands to be equal")
		})
	}
}

func TestFindWinner(t *testing.T) {
	tcs := []struct {
		hs     []string
		table  string
		winner []int
	}{
		{
			hs: []string{
				"qc9d", "7s5d", "as4h", "qh3c",
				"th6s", "8h3d", "kd8d", "ah7c",
			},
			table:  "6h5ckh8c3h",
			winner: []int{6},
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s %s", strings.Join(tc.hs, ","), tc.table), func(tt *testing.T) {
			tbcs := card.ParseMultiPokerCardString(tc.table)

			tm := map[int]hand.Hand{}
			for hi, h := range tc.hs {
				hob := FindBestHand(card.ParseMultiPokerCardString(h), tbcs)
				tm[hi] = hob
			}

			assert.Equal(tt, tc.winner, findTableWinner(tm), "expected winners to be equal")
		})
	}
}
