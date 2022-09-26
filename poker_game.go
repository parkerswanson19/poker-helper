package pokerhelper

import (
	"fmt"
)

func newGame() {
	deck := newDeck()

	fmt.Println(deck)
}
