package golf

import "fmt"

type Rank int

const (
	_      = iota // throw away 0
	A Rank = iota
	_2
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	T
	J
	Q
	K
	R
)

func (r Rank) String() string {
	switch r {
		case A: return "A"
		case _2: return "2"
		case _3: return "3"
		case _4: return "4"
		case _5: return "5"
		case _6: return "6"
		case _7: return "7"
		case _8: return "8"
		case _9: return "9"
		case T: return "T"
		case J: return "J"
		case Q: return "Q"
		case K: return "K"
		case R: return "R"
	}
	return "ERR"
}

func (r Rank) Value() int {
	switch x := r; {
	case A <= x && x <= T: return x
	case J <= x && x <= K: return 10
	case x == R: return -5
	}
	return 100
}

type Suit int

const (
	C Suit = iota
	D
	H
	S
	R
)

func (s Suit) String() string {
	switch r {
		case C: return "c"
		case D: return "d"
		case H: return "h"
		case S: return "s"
		case R: return "R"
	}
	return "ERS"
}

/*type Card struct {
	rank Rank
	suit Suit
	value int
	faceup bool
}*/
