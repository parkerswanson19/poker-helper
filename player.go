package pokerhelper

type player struct {
	name     string
	hand     []Card
	analysis Analysis
}

func newPlayer() player {
	dealtHand := []Card{}
	dealtAnalysis := newAnalysis()

	for i := 0; i < 2; i++ {
		dealtHand = append(dealtHand, newCard())
	}

	return player{
		name:     "Parker",
		hand:     dealtHand,
		analysis: dealtAnalysis,
	}
}
