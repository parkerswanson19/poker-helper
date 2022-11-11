package pokerHelper

import (
	"fmt"
)

func NewGame(numPlayers, numTotalChips int, initPlayers bool) ([]Player, Deck) {
	deck := NewDeck()
	fmt.Println("Unshuffled: ")
	fmt.Println(deck, "\n")

	deck.ShuffleDeck()

	fmt.Println("Shuffled: ")
	fmt.Println(deck, "\n")

	chipsPerPlayer := numTotalChips / numPlayers

	players := []Player{}
	for i := 0; i < numPlayers; i++ {
		players = append(players, newPlayer(chipsPerPlayer))
		if initPlayers {
			for j := 0; j < 2; j++ {
				players[i].Hand = append(players[i].Hand, deck.DealCard())
			}
		}
	}

	fmt.Println("After dealing: ")
	fmt.Println(deck, "\n")

	return players, deck
}
