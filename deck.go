package golf

func Deck() []Card {
	var cards []Card
	// Omit joker (rank & suit) by exploiting that they are last in
	// their respective slices.
	for _, s := range suits[:len(suits)-1] {
		for _, r := range ranks[:len(ranks)-1] {
			cards = append(cards, s|r)
		}
	}
	return append(cards, Jok|Joker, Jok|Joker)
}
