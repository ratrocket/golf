package golf

import(
	"testing"
	"fmt"
)

func TestString(t *testing.T) {
	h := egHand()
	fmt.Printf("%v", h)
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
