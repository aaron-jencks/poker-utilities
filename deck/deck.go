package deck

import (
	"math/rand"

	"github.com/aaron-jencks/poker/card"
)

// Deck represents a deck of cards
type Deck struct {
	Cards []card.Card
}

// Count returns the number of cards left in the deck
func (d Deck) Count() int {
	return len(d.Cards)
}

// Shuffle shuffles the remaining cards in the deck
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

// Draw returns the top n cards of the deck and removes them from the deck
func (d *Deck) Draw(n int) []card.Card {
	if n == 0 || n >= d.Count() {
		return []card.Card{}
	}

	result := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return result
}

// DrawCard returns a specific card from the deck and removes it from the deck
func (d *Deck) DrawCard(c card.Card) card.Card {
	for ci, dc := range d.Cards {
		if dc == c {
			d.Cards = append(d.Cards[:ci], d.Cards[ci+1:]...)
			return dc
		}
	}
	return card.EMPTY
}

// CardFaceProbability returns the probability that f is the next card
func (d *Deck) CardFaceProbability(f card.CardFace) float64 {
	var count float64 = 0
	for _, dc := range d.Cards {
		if dc.Face() == f {
			count++
		}
	}
	return count / float64(d.Count())
}

// CardSuitProbability returns the probability that the next card has the suit s
func (d *Deck) CardSuitProbability(s card.CardSuit) float64 {
	var count float64 = 0
	for _, dc := range d.Cards {
		if dc.Suit() == s {
			count++
		}
	}
	return count / float64(d.Count())
}

// CreateEmptyDeck returns a deck with no cards in it
func CreateEmptyDeck() Deck {
	return Deck{
		make([]card.Card, 0),
	}
}

// CreateStandardDeck generates a deck with the standard 52 cards in it, plus a specified number of jokers
func CreateStandardDeck() Deck {
	d := Deck{
		make([]card.Card, 0, 52),
	}
	for s := byte(0); s < byte(card.SUITS); s++ {
		for f := card.TWO; f < card.JOKER; f++ {
			d.Cards = append(d.Cards, card.CreateCard(byte(f), s))
		}
	}
	d.Shuffle()
	return d
}

// CreateStackedDeck generates a deck with a specific set of cards in it
func CreateStackedDeck(cards []card.Card) Deck {
	return Deck{
		cards,
	}
}
