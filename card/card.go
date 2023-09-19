package card

type CardSuit byte

const (
	CLUBS CardSuit = iota
	DIAMONDS
	HEARTS
	SPADES
	SUITS
)

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

const EMPTY = Card(0)

// lowest 2 bits are the suit
// the next 4 bits are the face
// the highest 2 bits are unused
// uu-ffff-ss
type Card byte

func (c Card) Face() CardFace {
	return CardFace(byte(c&0x3c) >> 2)
}

func (c Card) Suit() CardSuit {
	return CardSuit(byte(c) & 0x03)
}

func CreateCard(face, suit byte) Card {
	return Card((face << 2) + suit)
}

func CreateJoker() Card {
	return Card(JOKER << 2)
}
