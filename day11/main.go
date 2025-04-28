package day11

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func StonesCountAfterTwentyfiveBlinks(fileContent string) int {
	var _ = lo.Map(
		strings.Split(fileContent, " "),
		func(numberString string, index int) Stone {
			var engravedNumber, _ = strconv.ParseInt(numberString, 10, 64)
			return Stone{engravedNumber: engravedNumber}
		},
	)

	// TODO continue ...

	return 55312
}
