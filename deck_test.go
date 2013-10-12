package golf

import "testing"

// Ummm, this *is* a deck...
var strCards = []string{
	"Ad", "2d", "3d", "4d", "5d", "6d", "7d", "8d", "9d",
	"Td", "Jd", "Qd", "Kd", "Ah", "2h", "3h", "4h", "5h",
	"6h", "7h", "8h", "9h", "Th", "Jh", "Qh", "Kh", "Ac",
	"2c", "3c", "4c", "5c", "6c", "7c", "8c", "9c", "Tc",
	"Jc", "Qc", "Kc", "As", "2s", "3s", "4s", "5s", "6s",
	"7s", "8s", "9s", "Ts", "Js", "Qs", "Ks", "R*", "R*",
}

func TestDeck(t *testing.T) {
	for i, c := range Deck() {
		if c.String() != strCards[i] {
			t.Errorf("%v: card unexpected in deck\n", c.String())
		}
	}
}

