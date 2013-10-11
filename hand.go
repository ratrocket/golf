package golf

import "fmt"

type Hand [8]Card

func (h Hand) String() string {
	// Really we can only show h[n] if h[n].Faceup()==true
	// Deal with that later though.
	return fmt.Sprintf("%v %v %v %v\n%v %v %v %v\n",
		h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7])
}

func (h Hand) Score() int {
	return 0
}

// TODO something like this???
// Keep track of what's been accounted for with a uint8?
//
// NB when scoring, check that the cards under consideration are face
// up.  This will allow for showing "running scores" during the game.
func boxes(h Hand, memo *uint8) { }
func columnses(h Hand, memo *uint8) { } // future
func columns(h Hand, memo *uint8) { }
func pips(h Hand, memo *uint8) { }
