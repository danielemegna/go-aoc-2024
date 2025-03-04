package day09

import (
	"github.com/samber/lo"
)

type DenseDiskMap struct {
	data []int
}

func ParseDenseDiskMap(diskMapString string) DenseDiskMap {
	var diskMapRunes = []rune(diskMapString)
	var diskMapData = lo.Map(diskMapRunes, func(value rune, _ int) int {
		return int(value - '0')
	})
	return DenseDiskMap{data: diskMapData}
}
