package simulation

import (
	"sort"

	"github.com/aaron-jencks/poker/card"
	"github.com/aaron-jencks/poker/deck"
	"github.com/aaron-jencks/poker/hand"
)

func FindBestHand(hole []card.Card, table []card.Card) hand.Hand {
	possibleHands := make([]hand.Hand, 10) // 5C3 is 10
	for ti0 := 0; ti0 < len(table)-2; ti0++ {
		for ti1 := ti0 + 1; ti1 < len(table)-1; ti1++ {
			for ti2 := ti1 + 1; ti2 < len(table); ti2++ {
				possibleHands = append(possibleHands, hand.FindHand([]card.Card{
					hole[0],
					hole[1],
					table[ti0],
					table[ti1],
					table[ti2],
				}))
			}
		}
	}
	sort.Slice(possibleHands, func(i, j int) bool { return possibleHands[i].LessThan(possibleHands[j]) })
	return possibleHands[len(possibleHands)-1]
}

func findTableWinner(handMap map[int]hand.Hand) []int {
	// find the winners
	var bestHands []int
	var bh hand.Hand
	first := true
	for seat, h := range handMap {
		if first || bh.LessThan(h) {
			bh = h
			bestHands = nil
			if first {
				first = false
			}
		}
		if bh.Equals(h) {
			bestHands = append(bestHands, seat)
		}
	}
	return bestHands
}

func SimulateTableHand(nplayers int, fixed_hands map[int][]card.Card, folds map[int]bool) []int {
	deck := deck.CreateStandardDeck()

	hcardMap := map[int][]card.Card{}
	for seat, hcards := range fixed_hands {
		hcardMap[seat] = hcards
		for _, c := range hcards {
			deck.DrawCard(c)
		}
	}

	// deal cards
	for si := 0; si < nplayers; si++ {
		if hcardMap[si] != nil {
			continue
		}
		hcardMap[si] = deck.Draw(2)
	}

	table := deck.Draw(5)

	// find hands
	handMap := map[int]hand.Hand{}
	for seat := 0; seat < nplayers; seat++ {
		if folds[seat] {
			continue
		}
		handMap[seat] = FindBestHand(hcardMap[seat], table)
	}

	return findTableWinner(handMap)
}
