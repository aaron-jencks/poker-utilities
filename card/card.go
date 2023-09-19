package card

// CardSuit represents an enum of card suits
type CardSuit byte

const (
	CLUBS CardSuit = iota
	DIAMONDS
	HEARTS
	SPADES
	SUITS
)

// CardFace represents an enum of card faces
type CardFace byte

const (
	TWO CardFace = iota + 2
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
	JOKER
	FACES
)

const EMPTY = Card(0) // represents an empty card with no face or suit

// represents a single playing card
// lowest 2 bits are the suit
// the next 4 bits are the face
// the highest 2 bits are unused
// uu-ffff-ss
type Card byte

// Face returns the face of the card
func (c Card) Face() CardFace {
	return CardFace(byte(c&0x3c) >> 2)
}

// Suit returns the suit of the card
func (c Card) Suit() CardSuit {
	return CardSuit(byte(c) & 0x03)
}

// CreateCard creates a card using the given suit and face
func CreateCard(face, suit byte) Card {
	return Card((face << 2) + suit)
}

// Equal returns true if the face of this card equals the face of the other
func (c Card) Equal(other Card) bool {
	return c.Face() == other.Face() // Texas Hold'em doesn't compare suits
}

// LessThan returns true if the face of this card is less than the face of the other
func (c Card) LessThan(other Card) bool {
	return c.Face() < other.Face() // Texas Hold'em doesn't compare suits
}
