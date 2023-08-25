package hand

import "github.com/aaron-jencks/poker/card"

type Hand struct {
	C1 card.Card
	C2 card.Card
}

func CreateHand(c1, c2 card.Card) Hand {
	return Hand{
		c1, c2,
	}
}
