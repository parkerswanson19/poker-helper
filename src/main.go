package main

import (
	"fmt"
	"log"
	ph "pokerhelper/src/poker-helper"
	"reflect"
)

func main() {

	players, curr_deck := ph.NewGame(5, 1000, false)

	fmt.Println("Players hand: ", players[0].Hand)

	// read in both of the user's cards
	fmt.Println("Input your first card (Suit, Rank): ")
	readInCards(players, curr_deck)
	fmt.Println("Input your second card (Suit, Rank): ")
	readInCards(players, curr_deck)

	fmt.Println("Players hand after input: ", players[0].Hand)
	fmt.Println("Deck: \n", curr_deck.GetDeck())

	// Start looking into the analysis

}

func readInCards(players []ph.Player, curr_deck ph.Deck) {
	var suit int
	var rank string
	_, err := fmt.Scanln(&suit, &rank)
	if err != nil {
		log.Fatal(err)
	}
	newCard := ph.NewCard(suit, rank)
	players[0].Hand = append(players[0].Hand, newCard)
	for i := 0; i < len(curr_deck.GetDeck()); i++ {
		if reflect.DeepEqual(curr_deck.GetDeck()[i], newCard) {
			curr_deck.RemoveCardFromDeck(i)
			i--
		}
	}
}
