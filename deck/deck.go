package deck

import (
	"math/rand"

	"github.com/aaron-jencks/poker/card"
)

type Deck struct {
	Cards []card.Card
}

func (d Deck) Count() int {
	return len(d.Cards)
}

func (d *Deck) Shuffle() {
	if d.Count() < 2 {
		return
	}
	rand.Shuffle(d.Count(), func(i, j int) {
		temp := d.Cards[i]
		d.Cards[i] = d.Cards[j]
		d.Cards[i] = temp
	})
}

func (d *Deck) Draw(n int) []card.Card {
	if n == 0 || n >= d.Count() {
		return []card.Card{}
	}

	result := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return result
}

func (d *Deck) DrawCard(c card.Card) card.Card {
	for ci, dc := range d.Cards {
		if dc == c {
			d.Cards = append(d.Cards[:ci], d.Cards[ci+1:]...)
			return dc
		}
	}
	return card.EMPTY
}

func CreateEmptyDeck() Deck {
	return Deck{
		make([]card.Card, 0),
	}
}

func CreateStandardDeck(jokers int) Deck {
	d := Deck{
		make([]card.Card, 0, 52),
	}
	for s := byte(0); s < card.SUITS; s++ {
		for f := card.TWO; f < card.JOKER; f++ {
			d.Cards = append(d.Cards, card.CreateCard(f, s))
		}
	}
	for j := 0; j < jokers; j++ {
		d.Cards = append(d.Cards, card.CreateJoker())
	}
	d.Shuffle()
	return d
}

func CreateStackedDeck(cards []card.Card) Deck {
	return Deck{
		cards,
	}
}
