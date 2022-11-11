package pokerHelper

type Analysis struct {
	odds      int
	curr_hand []card
	curr_deck Deck
}

func newAnalysis(d Deck, h []card) Analysis {
	return Analysis{
		odds:      9,
		curr_hand: h,
		curr_deck: d,
	}
}

func PerformAnalysis() {

}
