package day11

import "github.com/samber/lo"

type LineOfStones map[Stone]int

func (this LineOfStones) Add(engravedNumber Stone, occurences int) {
	var currentOccurrencesCount, _ = this[engravedNumber]
	this[engravedNumber] = currentOccurrencesCount + occurences
}

func (this LineOfStones) Size() int {
	return lo.Sum(lo.Values(this))
}

func (this LineOfStones) OnBlink() LineOfStones {
	var newLine = LineOfStones{}
	for engravedNumber, occurrences := range this {
		var left, right = engravedNumber.OnBlink()
		newLine.Add(left, occurrences)
		if right != nil {
			newLine.Add(*right, occurrences)
		}
	}
	return newLine
}
