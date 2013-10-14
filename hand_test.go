package golf

import(
	"testing"
	"fmt"
)

// need to load this from a file
/*
var hands = []Hand{
	Hand{
	},
	Hand{
	},
	Hand{
	},
}
*/

func TestString(t *testing.T) {
	h := egHand()
	output := fmt.Sprintf("%v", h)
	should := "As R* 6s Qs\nJh 3c Kd Ad\n"
	if output != should {
		t.Errorf("Hand's String() incorrect\n")
	}
}

func TestOneScore(t *testing.T) {
	if egHand().Score() != 26 {
		t.Errorf("Score not 26 on egHand()\n")
	}
}

func TestScore(t *testing.T) {
}

func egHand() Hand {
	return Hand{
		Ace|Spade|Faceup,
		Jak|Heart|Faceup, // 11
		Jok|Joker|Faceup,
		Tre|Club|Faceup,  // -2
		Six|Spade|Faceup,
		Kng|Diamond|Faceup, // 6
		Qen|Spade|Faceup,
		Ace|Diamond|Faceup, // 11; total 26
	}
}
