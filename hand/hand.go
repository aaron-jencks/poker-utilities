package hand

import (
	"sort"

	"github.com/aaron-jencks/poker/card"
)

// PokerHands represents a poker hand ranking
// and label
type PokerHands byte

const (
	HIGH_CARD PokerHands = iota
	PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
	ROYAL_FLUSH
)

// Hand represents a hand in poker
type Hand struct {
	Hand     PokerHands    // the determined hand of the given cards
	Kicker0  card.CardFace // the kicker used for determining ties
	Kicker1  card.CardFace // the second kicker used for determining ties
	Contents []card.Card   // all of the cards in the hand, used for breaking ties in the case of a flush
}

// Contains returns true if any of the cards in the hand contain the given face
func (h Hand) Contains(f card.CardFace) bool {
	for _, hc := range h.Contents {
		if hc.Face() == f {
			return true
		}
	}
	return false
}

// Equals returns true if the other hand is equivalent in ranking to this hand
func (h Hand) Equals(other Hand) bool {
	if h.Hand == other.Hand {
		if h.Hand == FLUSH {
			// we need to compare full hands for this one
			for ci := range h.Contents {
				if h.Contents[ci] != other.Contents[ci] {
					return false
				}
			}
			return true
		}

		sal := (h.Hand == STRAIGHT || h.Hand == STRAIGHT_FLUSH) && h.Contains(card.ACE) && h.Contains(card.FIVE)
		oal := (other.Hand == STRAIGHT || other.Hand == STRAIGHT_FLUSH) && other.Contains(card.ACE) && other.Contains(card.FIVE)
		if sal {
			return sal && oal
		}
		if h.Kicker0 == other.Kicker0 && h.Kicker1 == other.Kicker1 {
			for ci := range h.Contents {
				if h.Contents[ci] != other.Contents[ci] {
					return false
				}
			}
			return true
		}
	}
	return false
}

// LessThan returns true if this hand's ranking is less than the other
func (h Hand) LessThan(other Hand) bool {
	if h.Hand == other.Hand {
		if h.Hand == FLUSH {
			// we need to compare full hands for this one
			for ci := range h.Contents {
				if h.Contents[ci] < other.Contents[ci] {
					return true
				} else if h.Contents[ci] > other.Contents[ci] {
					return false
				}
			}
			return false
		}

		sal := (h.Hand == STRAIGHT || h.Hand == STRAIGHT_FLUSH) && h.Contains(card.ACE) && h.Contains(card.FIVE)

		if !sal && h.Kicker0 == other.Kicker0 && h.Kicker1 == other.Kicker1 {
			// not a straight and kickers match, compare full hand
			for ci := range h.Contents {
				if h.Contents[ci] < other.Contents[ci] {
					return true
				} else if h.Contents[ci] > other.Contents[ci] {
					return false
				}
			}
			return false
		}

		return !sal && (h.Kicker0 == other.Kicker0 || h.Kicker1 == other.Kicker1)
	}
	return h.Hand < other.Hand
}

func (h Hand) String() string {
	r := ""
	for _, c := range h.Contents {
		r += c.String()
	}
	return r
}

// is_straight returns whether the cards array contains a straight or not
func is_straight(cards []card.Card) bool {
	sort.Slice(cards, func(i, j int) bool { return cards[i].LessThan(cards[j]) })

	for ci := 1; ci < len(cards); ci++ {
		// edge case for ace low straight
		if ci == len(cards)-1 && cards[ci].Face() == card.ACE && cards[0].Face() == card.TWO {
			return true
		}

		// all other cases, including ace high straight
		if cards[ci].Face()-cards[ci-1].Face() != 1 {
			return false
		}
	}

	return true
}

func FindHand(cards []card.Card) Hand {
	sort.Slice(cards, func(i, j int) bool { return cards[i].LessThan(cards[j]) })
	fcounts := map[card.CardFace]int{}
	scounts := map[card.CardSuit]int{}
	for _, c := range cards {
		if _, fok := fcounts[c.Face()]; !fok {
			fcounts[c.Face()] = 1
		} else {
			fcounts[c.Face()] += 1
		}

		if _, sok := scounts[c.Suit()]; !sok {
			scounts[c.Suit()] = 1
		} else {
			scounts[c.Suit()] += 1
		}
	}

	straight := is_straight(cards)
	if straight {
		// check for edge case of ace low
		al := cards[0].Face() == card.TWO && cards[len(cards)-1].Face() == card.ACE
		k0 := cards[len(cards)-1].Face()
		k1 := cards[len(cards)-2].Face()
		if al {
			k0 = cards[len(cards)-2].Face()
			k1 = cards[len(cards)-3].Face()
		}

		if len(scounts) == 1 {
			if cards[0].Face() == card.TEN {
				// royal flush
				return Hand{
					Hand:     ROYAL_FLUSH,
					Kicker0:  k0,
					Kicker1:  k1,
					Contents: cards,
				}
			}

			return Hand{
				Hand:     STRAIGHT_FLUSH,
				Kicker0:  k0,
				Kicker1:  k1,
				Contents: cards,
			}
		}

		return Hand{
			Hand:     STRAIGHT,
			Kicker0:  k0,
			Kicker1:  k1,
			Contents: cards,
		}
	}

	tak := false
	pr := false
	tpr := false
	for k, v := range fcounts {
		if v == 4 {
			// highest possible ranking hand at this point
			// two players cannot have the same 4 of a kind at once
			// so only one kicker is needed, the face of the 4
			return Hand{
				Hand:     FOUR_OF_A_KIND,
				Kicker0:  k,
				Kicker1:  k,
				Contents: cards,
			}
		} else if v == 3 {
			tak = true
		} else if v == 2 {
			if pr {
				tpr = true
			} else {
				pr = true
			}
		}
	}

	if tak && pr {
		// full house
		// determine the kickers
		k0 := card.ACE
		k1 := card.ACE
		for k, v := range fcounts {
			if v == 3 {
				k0 = k
			} else {
				k1 = k
			}
		}

		return Hand{
			Hand:     FULL_HOUSE,
			Kicker0:  k0,
			Kicker1:  k1,
			Contents: cards,
		}
	}

	if tak {
		k0 := card.TWO
		k1 := card.TWO
		for k, v := range fcounts {
			if v == 3 {
				k0 = k
			} else if k > k1 {
				k1 = k
			}
		}

		return Hand{
			Hand:     THREE_OF_A_KIND,
			Kicker0:  k0,
			Kicker1:  k1,
			Contents: cards,
		}
	}

	if tpr {
		k0 := card.TWO
		k1 := card.TWO
		for k, v := range fcounts {
			if v == 2 {
				if k > k0 {
					k1 = k0
					k0 = k
				}
			}
		}

		return Hand{
			Hand:     TWO_PAIR,
			Kicker0:  k0,
			Kicker1:  k1,
			Contents: cards,
		}
	}

	if pr {
		k0 := card.TWO
		for k, v := range fcounts {
			if v == 2 {
				k0 = k
			}
		}

		k1 := cards[len(cards)-1].Face()
		if cards[len(cards)-2].Face() == k1 {
			// pair is the highest card
			k1 = cards[len(cards)-3].Face()
		}

		return Hand{
			Hand:     TWO_PAIR,
			Kicker0:  k0,
			Kicker1:  k1,
			Contents: cards,
		}
	}

	return Hand{
		Hand:     HIGH_CARD,
		Kicker0:  cards[len(cards)-1].Face(),
		Kicker1:  cards[len(cards)-2].Face(),
		Contents: cards,
	}
}

func ParsePokerHandString(s string) Hand {
	cards := make([]card.Card, 0, len(s)>>1)
	for si := 1; si < len(s); si += 2 {
		cards = append(cards, card.ParsePokerCardString(s[si-1:si+1]))
	}

	return FindHand(cards)
}
