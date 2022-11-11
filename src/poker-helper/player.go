package pokerHelper

type Player struct {
	name     string
	Hand     []card
	chips    int
	analysis Analysis
}

func newPlayer(chips int) Player {
	dealtHand := []card{}

	return Player{
		name:  "Parker",
		Hand:  dealtHand,
		chips: chips,
	}
}
