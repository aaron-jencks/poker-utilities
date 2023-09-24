package hand

import "github.com/aaron-jencks/poker/card"

type PokerRange struct {
	F0     card.CardFace
	F1     card.CardFace
	Suited bool
}

func (r PokerRange) Pairs() [][]card.Card {
	results := [][]card.Card{}
	for s := card.CLUBS; s < card.SUITS; s++ {
		c0 := card.CreateCard(r.F0, s)

		if r.Suited {
			c1 := card.CreateCard(r.F1, s)
			results = append(results, []card.Card{c0, c1})
			continue
		}

		for s1 := card.CLUBS; s1 < card.SUITS; s1++ {
			if s == s1 {
				continue
			}

			c1 := card.CreateCard(r.F1, s1)
			results = append(results, []card.Card{c0, c1})
		}
	}
	return results
}
