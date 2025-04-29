package day11

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func StonesCountAfterTwentyfiveBlinks(fileLine string) int {
	var stones = lo.Map(
		strings.Split(fileLine, " "),
		func(numberString string, index int) Stone {
			var engravedNumber, _ = strconv.ParseInt(numberString, 10, 64)
			return Stone{engravedNumber: engravedNumber}
		},
	)

	for range 25 {
		stones = StonesOnBlink(stones)
	}

	return len(stones)
}
