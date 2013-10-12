package golf

import(
	"testing"
	"fmt"
)

func TestString(t *testing.T) {
	h := egHand()
	output := fmt.Sprintf("%v", h)
	should := "As R* 6s Qs\nJh 3c Kd Ad\n"
	if output != should {
		t.Errorf("Hand's String() incorrect\n")
	}
}

func TestScore(t *testing.T) {
}

func egHand() Hand {
	return Hand{
		Ace|Spade|Faceup,
		Jak|Heart|Faceup,
		Jok|Joker|Faceup,
		Tre|Club|Faceup,
		Six|Spade|Faceup,
		Kng|Diamond|Faceup,
		Qen|Spade|Faceup,
		Ace|Diamond|Faceup,
	}
}
