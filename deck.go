package pokerhelper

// suit: 1 = clubs, 2 = spades, 3 = hearts, 3 = diamonds

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var suits = []int{1, 2, 3, 4}

type Card struct {
	rank string
	suit int
}

type Deck struct {
	deck []Card
}

func newCard() Card {
	return Card{
		rank: "A",
		suit: 1,
	}
}

func newDeck() Deck {

	d := []Card{}

	for _, r := range ranks {
		for _, s := range suits {
			d = append(d, Card{rank: r, suit: s})
		}
	}

	return Deck{
		deck: d,
	}

}
