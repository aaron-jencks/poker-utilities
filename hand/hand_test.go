package hand

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandRanking(t *testing.T) {
	hands := []string{
		"2d5s6sjhac",
		"5h5s2cjcad",
		"qhqcqd5sad",
		"7c8h9dthjs",
		"tc4cqc7c2c",
		"8h8d8ckhks",
		"ahadacas4d",
		"4h5h6h7h8h",
		"tsjsqsksas",
	}

	shands := make([]Hand, len(hands))
	for hi, h := range hands {
		shands[hi] = ParsePokerHandString(h)
	}

	rand.Shuffle(len(shands), func(i, j int) { shands[i], shands[j] = shands[j], shands[i] })

	sort.Slice(shands, func(i, j int) bool { return shands[i].LessThan(shands[j]) })

	for ai := range hands {
		assert.Equal(t, hands[ai], shands[ai].String())
	}
}
