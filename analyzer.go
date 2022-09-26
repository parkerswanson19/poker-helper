package pokerhelper

type Analysis struct {
	odds int
}

func newAnalysis() Analysis {
	return Analysis{
		odds: 9,
	}
}
