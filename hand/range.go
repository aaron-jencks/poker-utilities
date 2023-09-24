package hand

import "github.com/aaron-jencks/poker/card"

type PokerRange struct {
	F0     card.CardFace
	F1     card.CardFace
	Suited bool
}
