package day11

import (
	"strconv"
	"strings"
)

func StonesCountAfterTwentyfiveBlinks(fileLine string) int {
	var stones = parseStones(fileLine)

	for range 25 {
		stones = stones.OnBlink()
	}

	return stones.Size()
}

func StonesCountAfterSeventyfiveBlinks(fileLine string) int {
	var stones = parseStones(fileLine)

	for range 75 {
		stones = stones.OnBlink()
	}

	return stones.Size()
}

func parseStones(fileLine string) LineOfStones {
	var lineOfStones = LineOfStones{}
	for numberString := range strings.SplitSeq(fileLine, " ") {
		var engravedNumber, _ = strconv.Atoi(numberString)
		lineOfStones.Add(engravedNumber, 1)
	}
	return lineOfStones
}
