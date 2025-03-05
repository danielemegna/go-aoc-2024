package day09

import (
	"github.com/samber/lo"
)

type DenseDiskMap struct {
	data []int
}

type ExpandedDiskMap = DenseDiskMap

func ParseDenseDiskMap(diskMapString string) DenseDiskMap {
	var diskMapRunes = []rune(diskMapString)
	var diskMapData = lo.Map(diskMapRunes, func(value rune, _ int) int {
		return int(value - '0')
	})
	return DenseDiskMap{data: diskMapData}
}

func ToExpandedDiskMap(denseDiskMap DenseDiskMap) ExpandedDiskMap {
	return ExpandedDiskMap{data: repeatInSlice(0, denseDiskMap.data[0])}
}

func repeatInSlice(value int, length int) []int {
	var diskMapData = []int{}
	for i := 0; i < length; i++ {
		diskMapData = append(diskMapData, value)
	}
	return diskMapData
}
