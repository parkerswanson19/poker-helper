package pokerHelper

import "math/rand"

// suit: 1 = clubs, 2 = spades, 3 = hearts, 3 = diamonds

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var suits = []int{1, 2, 3, 4}

type card struct {
	rank string
	suit int
}

type Deck struct {
	deck []card
}

func (d *Deck) GetDeck() []card {
	return d.deck
}

func (d *Deck) RemoveCardFromDeck(index int) {
	if index < len(d.deck) {
		d.deck = append(d.deck[:index], d.deck[index+1:]...)
	}
}

func NewCard(suit int, rank string) card {
	return card{
		suit: suit,
		rank: rank,
	}
}

func (d *Deck) DealCard() card {
	deck_length := len(d.deck)
	rand_card := rand.Intn(deck_length)

	toReturn := d.deck[rand_card]
	d.deck = append(d.deck[:rand_card], d.deck[rand_card+1:]...)

	return toReturn

}

func NewDeck() Deck {

	d := []card{}

	for _, r := range ranks {
		for _, s := range suits {
			d = append(d, card{rank: r, suit: s})
		}
	}

	return Deck{
		deck: d,
	}

}

func (d *Deck) ShuffleDeck() {
	rand.Shuffle(len(d.deck), func(i, j int) {
		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
	})

}
