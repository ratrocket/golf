// The layout of the hand in its array is:
//   0 2 4 6
//   1 3 5 7
// That simplifies scoring.
package golf

import "fmt"

type Hand [8]Card

// NB when scoring, check that the cards under consideration are face
// up.  This will allow for showing "running scores" during the game.
//
// In golang, how do you decide when to make something a function and
// when to make something a method?

func (h Hand) String() string {
	// Really we can only show h[n] if h[n].Faceup()==true
	// Deal with that later though.
	return fmt.Sprintf("%v %v %v %v\n%v %v %v %v\n",
		h[0], h[2], h[4], h[6],
		h[1], h[3], h[5], h[7])
}

// Could we do the "memoization" of what cards we've already scored with
// a closure around a "memo" variable?
// Check out that part of the gotour (around page 50ish?)
func (h Hand) Score() int {
	score := 0

	if h.box0() {
		score += h.boxScore(0)
		// Here till scoring colOrPip is "goto consider pos 4"
		if h.box4() {
			score += h.boxScore(4)
			return score
		}
		score += h.colOrPip(4)
		score += h.colOrPip(6)
		return score
	}
	if h.box2() {
		score += h.boxScore(2)
		if h.colsB() {
			score += h.colsScore('B')
			return score
		}
		score += h.colOrPip(0)
		score += h.colOrPip(6)
		return score
	}
	if h.box4() { // This is "consider pos 4" (I think)
		score += h.boxScore(4)
		// don't have box0 and no columns possible
		score += h.colOrPip(0)
		score += h.colOrPip(2)
		return score
	}

	// now completely done considering box scenarios
	// if you make it here there are NO boxes
	if h.colsA() {
		score += h.colsScore('A')
		if h.colsC() {
			score += h.colsScore('C')
			return score
		}
		score += h.colOrPip(2)
		score += h.colOrPip(6)
		return score
	}
	if h.colsB() {
		score += h.colsScore('B')
		score += h.colOrPip(2) // remember: no boxes possible here
		score += h.colOrPip(4)
		return score
	}
	if h.colsC() {
		score += h.colsScore('C')
		score += h.colOrPip(0)
		score += h.colOrPip(4)
		return score
	}

	// here there are no boxes and no columnses
	score += h.colOrPip(0)
	score += h.colOrPip(2)
	score += h.colOrPip(4)
	score += h.colOrPip(6)
	return score
}

func (h Hand) box0() bool {
	return h[0].Equal(h[1]) && h[1].Equal(h[2]) && h[2].Equal(h[3])
}
func (h Hand) box2() bool {
	return h[2].Equal(h[3]) && h[3].Equal(h[4]) && h[4].Equal(h[5])
}
func (h Hand) box4() bool {
	return h[4].Equal(h[5]) && h[5].Equal(h[6]) && h[6].Equal(h[7])
}
func (h Hand) colsA() bool {
	return h[0].Equal(h[1]) && h[1].Equal(h[4]) && h[4].Equal(h[5])
}
func (h Hand) colsB() bool {
	return h[0].Equal(h[1]) && h[1].Equal(h[6]) && h[6].Equal(h[7])
}
func (h Hand) colsC() bool {
	return h[2].Equal(h[3]) && h[3].Equal(h[6]) && h[6].Equal(h[7])
}

// i has to be 0, 2, 4, 6
func (h Hand) colOrPip(i int) int {
	if h[i].Equal(h[i+1]) {
		return h.colScore(i)
	}
	return h[i].Value() + h[i+1].Value()
}

// Purpose of these is mostly to account for jokers.
func (h Hand) colScore(i int) int {
	if h[i].Joker() {
		return -10
	}
	return 0
}
func (h Hand) boxScore(i int) int {
	if h[i].Joker() {
		return -50 // joker box!
	}
	return -10
}
func (h Hand) colsScore(r rune) int {
	var i int
	switch r {
	case 'A', 'B':
		i = 0
	case 'C':
		i = 2
	}
	if h[i].Joker() {
		return -20 // joker columns?
	}
	return -5
}
