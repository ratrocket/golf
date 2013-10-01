package golf

import "fmt"

type Rank int

const (
	A Rank = 1
	_2 Rank = 2
	_3 Rank = 3
	_4 Rank = 4
	_5 Rank = 5
	_6 Rank = 6
	_7 Rank = 7
	_8 Rank = 8
	_9 Rank = 9
	T Rank = 10
	J Rank = 10
	Q Rank = 10
	K Rank = 10
	R Rank = -5
)

type Suit int

const (
	C Suit = 1 //iota
	D Suit
	H Suit
	S Suit
	R Suit
)

/*type Card struct {
	rank Rank
	suit Suit
	value int
	faceup bool
}*/
