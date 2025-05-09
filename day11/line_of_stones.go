package day11

import "github.com/samber/lo"

type LineOfStones map[Stone]int

func (this LineOfStones) Add(stone Stone, newOccurences int) {
	var currentOccurrences, _ = this[stone]
	this[stone] = currentOccurrences + newOccurences
}

func (this LineOfStones) Size() int {
	return lo.Sum(lo.Values(this))
}

func (this LineOfStones) OnBlink() LineOfStones {
	var newLine = LineOfStones{}
	for stone, occurrences := range this {
		var left, right = stone.OnBlink()
		newLine.Add(left, occurrences)
		if right != nil {
			newLine.Add(*right, occurrences)
		}
	}
	return newLine
}
