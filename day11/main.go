package day11

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func StonesCountAfterTwentyfiveBlinks(fileLine string) int {
	var stones = parseStones(fileLine)

	for range 25 {
		stones = StonesOnBlink(stones)
	}

	return len(stones)
}

func StonesCountAfterSeventyfiveBlinks(fileLine string) int {
	var stones = parseStones(fileLine)

	for range 75 {
		stones = StonesOnBlink(stones)
	}

	return len(stones)
}

func parseStones(fileLine string) []Stone {
	return lo.Map(
		strings.Split(fileLine, " "),
		func(numberString string, index int) Stone {
			var engravedNumber, _ = strconv.ParseInt(numberString, 10, 64)
			return Stone{engravedNumber: engravedNumber}
		},
	)
}
