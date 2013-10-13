package golf

import(
	"testing"
	"fmt"
	"math"
)

// Show representation of cards in different bases.
// name: "rank", "suit", or "fcup" (for faceup)
// value: the string representation of the suit or rank (made something
// 	up for Faceup)
// c: the card to show
func repr(name, value string, c Card) {
	str := "%s %v: %032b %08x (%d)\n"
	fmt.Printf(str, name, value, c, uint32(c), c)
}

func TestRanks(t *testing.T) {
	for i, rank := range Ranks() {
		repr("rank", rank.Prank(), rank)

		expected := Card(math.Pow(2, float64(i)))
		if rank != expected {
			t.Errorf("rank %v not as expected", i + 1)
		}
	}
}

func TestSuits(t *testing.T) {
	for i, suit := range Suits() {
		repr("suit", suit.Psuit(), suit)

		expected := Card(math.Pow(2, float64(i + 16)))
		if suit != expected {
			t.Errorf("suit %v not as expected", i + 1)
		}
	}
}

func TestFaceup(t *testing.T) {
	repr("fcup", "t", Faceup)
}

/*
func TestOne(t *testing.T) {
	//var one uint32 = 1
	fmt.Printf("one: \n  bin: %0*b\n  hex: %#0 x\n  dec: %d\n",
		//32, one, one, one)
		//32, Ace, Ace, Ace)
		32, uint32(Ace), uint32(Ace), uint32(Ace))

	if 1 == uint32(Ace) {
		fmt.Println("Ace is one")
	} else {
		fmt.Println("Ace is not one")
	}
}
*/

func TestValidRankMask(t *testing.T) {
	var ranksOrd Card // "OR"d
	var expected Card = 0x00003FFF

	for _, r := range Ranks() {
		ranksOrd |= r
	}
	if ranksOrd != expected {
		t.Errorf("valid-rank mask not as expected")
	}
}

func TestValidSuitMask(t *testing.T) {
	var suitsOrd Card
	var expected Card = 0x001F0000

	for _, s := range Suits() {
		suitsOrd |= s
	}
	if suitsOrd != expected {
		t.Errorf("valid-suit mask not as expected")
	}
}

func TestValid(t *testing.T) {
	for _, c := range Deck() {
		if !c.Valid() { // face down (c&Faceup == 0)
			t.Errorf("%v not valid (face down)\n", c)
		}
		if !(c|Faceup).Valid() { // face up (c&Faceup > 0)
			t.Errorf("%v not valid (face up)\n", c)
		}
	}

	for _, c := range Ranks() { // Ranks alone not valid
		if c.Valid() {
			t.Errorf("%v is valid, shouldn't be\n", c.Prank())
		}
	}

	for _, c := range Suits() { // Suits alone not valid
		if c.Valid() {
			t.Errorf("%v is valid, shouldn't be\n", c.Psuit())
		}
	}
	// check output of invalids() (all face up AND down)
	for _, c := range invalids() {
		if c.Valid() {
			t.Errorf("%0*b valid (face down)\n", 32, c)
		}
		if (c|Faceup).Valid() {
			t.Errorf("%0*b valid (face up)\n", 32, c)
		}
	}
}

func TestCountValidsInDeck(t *testing.T) {
	var valids int
	for _, c := range Deck() {
		if c.Valid() {
			//fmt.Printf("%x\n", c)
			valids++
		}
	}
	if valids != 54 {
		t.Errorf("Not all deck is valid")
	}
}

func TestValue(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 0}
	for i, c := range Deck() {
		switch {
		case i < 52:
			if c.Value() != values[i % 13] {
				t.Errorf("%v score wrong", c)
			}
		case i >= 52:
			if c.Value() != -5 {
				t.Errorf("%v score wrong (joker)", c)
			}
		}
	}
}

func TestJoker(t *testing.T) {
	jok1 := Jok|Joker|Faceup
	jok2 := Jok|Joker
	notJok1 := Ace|Spade|Faceup
	notJok2 := For|Heart|Faceup

	if !jok1.Joker() || !jok2.Joker() {
		t.Errorf("A joker is not a joker")
	}
	if notJok1.Joker() || notJok2.Joker() {
		t.Errorf("A non-joker is a joker")
	}
}

func TestEqual(t *testing.T) {
	ace := Ace|Spade|Faceup
	ac2 := Ace|Heart
	two := Two|Diamond
	tw2 := Two|Club|Faceup
	kng := Kng|Heart
	kn2 := Kng|Club|Faceup
	qen := Qen|Spade

	if !ace.Equal(ac2) { t.Errorf("Aces not equal") }
	if !tw2.Equal(two) { t.Errorf("Twos not equal") }
	if !kng.Equal(kn2) { t.Errorf("Kngs not equal") }

	if qen.Equal(ace) { t.Errorf("Qen == Ace") }
	if ac2.Equal(kn2) { t.Errorf("Ace == Kng") }
	if two.Equal(qen) { t.Errorf("Two == Qen") }
}

func takesTooLongTestCount106Valids(t *testing.T) {
	var valids int
	var c Card = 0
	for ; c <= 0xFFFFFFFF; c++ {
		if c.Valid() {
			valids++
		}
	}
	fmt.Println("valids counted: ", valids)
	if valids != 106 {
		t.Errorf("Not 106 valids")
	}
}

func invalids() []Card {
	ranks := Ranks()
	suits := Suits()

	invs := []Card{
		Joker | Kng, // Joker suit and any non-Jok rank
		Spade | Jok, // any non-Joker suit and Jok rank
		// anything with bits 2^14 or 2^15 set
		1 << 14 | Tre,     // 2^14 w/ valid rank
		1 << 15 | For,     // 2^15 w/ valid rank
		1 << 14 | Heart,   // 2^14 w/ valid suit
		1 << 15 | Diamond, // 2^15 w/ valid suit
	}
	// anything with bits 2^21 to 2^30 set
	//   say: all of those w/ valid rank <= fail b/c no suit tho...
	//        all of those w/ valid rank and suit
	for i := 21; i < 31; i++ {
		rank := ranks[i % len(ranks)]
		suit := suits[i % len(suits)]

		r  := 1 << uint(i) | rank
		rs := 1 << uint(i) | rank | suit

		invs = append(invs, r, rs)
	}
	return invs
}

func TestUnaryComplement(t *testing.T) {
	var jok uint16 = 0x0020 // 0000 0000 0010 0000
	var msk uint16 = 0x003F // 0000 0000 0011 1111
	// we WANT 0x001F, iow, msk with jok bit unset

	// msk - jok == 0x001F
	if msk - jok != 0x001F {
		t.Errorf("complement fail 1")
	}
	mskSansJok := msk - jok
	if mskSansJok != (msk & ^jok) { // msk & ^jok ==? 0x001F // YES
		t.Errorf("complement fail 2")
	}
	// golang's bit clear (AND NOT) operator
	if mskSansJok != (msk &^ jok) {
		t.Errorf("complement fail 3")
	}
}
