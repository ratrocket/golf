package golf

/*
A card is a bitset.  The lower 16 bits indicate the rank (there are 14
ranks including the joker) and the upper 16 bits indicate the suit
(jokers have a "no suit" suit) and whether or not the card is face up.
*/
type Card uint32

const (
	// Ranks
	Ace Card = (1 << iota) // 2^0
	Two
	Tre
	For
	Fiv
	Six
	Sev
	Eit
	Nin // 2^8
	Ten
	Jak
	Qen
	Kng
	Jok // 2^13
	_   // 2^14 discarded
	_   // 2^15 discarded
	// Suits
	Diamond // 2^16
	Heart
	Club
	Spade
	Joker // 2^20
	// "boolean" to indicate face up/down; highest bit pos
	Faceup Card = 1 << 31
)

const (
	AllRanks Card = 0x00003FFF // All ranks OR'd
	AllSuits Card = 0x001F0000 // All suits OR'd
	InvalidBits Card = 0x7FE0C000 // All must be 0
)

var ranks = [14]Card{
	Ace, Two, Tre, For, Fiv, Six, Sev,
	Eit, Nin, Ten, Jak, Qen, Kng, Jok,
}
var suits = [5]Card{Diamond, Heart, Club, Spade, Joker}

func Ranks() [14]Card { return ranks }
func Suits() [5]Card  { return suits }

func (c Card) Prank() string { // Print rank
	switch {
	case c&Ace > 0: return "A"
	case c&Two > 0: return "2"
	case c&Tre > 0: return "3"
	case c&For > 0: return "4"
	case c&Fiv > 0: return "5"
	case c&Six > 0: return "6"
	case c&Sev > 0: return "7"
	case c&Eit > 0: return "8"
	case c&Nin > 0: return "9"
	case c&Ten > 0: return "T"
	case c&Jak > 0: return "J"
	case c&Qen > 0: return "Q"
	case c&Kng > 0: return "K"
	case c&Jok > 0: return "R"
	}
	return "ERR"
}

func (c Card) Psuit() string {
	switch {
	case c&Diamond > 0: return "d"
	case c&Heart > 0:   return "h"
	case c&Club > 0:    return "c"
	case c&Spade > 0:   return "s"
	case c&Joker > 0:   return "*"
	}
	return "ERS"
}

/*
validness
exactly one rank has to be set
card | [0 repeated 16] 0011 1111 1111 1111 > 0
that mask is 0x00003FFF
the mask is also all Ranks() OR'd together

exactly one suit has to be set
card | [0 repeated 8] 0001 1111 [0 repeated 16] > 0
that mask is 0x001F0000
the mask is also all Suits() OR'd together

they have to be >0 AND ensure exactly one bit is set via:
card & (card - 1) == 0
(Have to do this for rank AND suit AFTER applying the valid{rank,suit}
mask.  Otherwise the suit/rank (resp.) will mess it up.)

ALSO there are illegal combinations of rank & suit, all having to do
with jokers.
To account for that we could test against masks WITHOUT the joker
rank/suit included (this validates all "normal" cards) and if that
validation fails explicitly test if it's a joker.

How do you test that exactly TWO bits are set...?
*/

func (c Card) Valid() bool {
	if c & InvalidBits > 0 {
		return false
	}

	// Exclude the joker rank/suit. Check explicitly later.
	rMask := AllRanks &^ Jok
	sMask := AllSuits &^ Joker

	rankOnly := rMask & c
	suitOnly := sMask & c

	// A {rank,suit} is set && exactly ONE {rank,suit} is set
	rankValid := rankOnly > 0 && (rankOnly&(rankOnly-1) == 0)
	suitValid := suitOnly > 0 && (suitOnly&(suitOnly-1) == 0)

	// valid joker if == Jok_rank|Jok_suit after Faceup bit cleared
	isJoker := c&^Faceup == Jok|Joker

	return (rankValid && suitValid) || isJoker
}

func (c Card) Value() int {
	if !c.Valid() {
		return 100
	}

	// strip out Faceup bit and suit
	// TODO rearrange this into its own function (it's also used in
	// pip(), and Equal()).
	c = c &^ (Faceup | AllSuits)

	switch {
	case Ace <= c && c <= Ten:
		return c.pip()
	case c == Jak || c == Qen:
		return 10
	case c == Kng:
		return 0
	case c == Jok:
		return -5
	}
	return 101
}

// Ignore faceupness, ignore suit.
// TODO should we check that they're valid?
func (c Card) Equal(o Card) bool {
	return (c &^ (Faceup | AllSuits)) == (o &^ (Faceup | AllSuits))
}

// Face value for Ace through Ten, ie, index+1 of set bit.
// There are smarter ways to find the set bit.
func (c Card) pip() int {
	c = c &^ (Faceup | AllSuits) // strip out Faceup bit and suit

	for i, r := range ranks[0:10] {
		if c == r {
			return i + 1
		}
	}
	return 102
}

func (c Card) String() string {
	return c.Prank() + c.Psuit()
}
