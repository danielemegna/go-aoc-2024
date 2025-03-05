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

func (this DenseDiskMap) ToExpandedDiskMap() ExpandedDiskMap {
	var expandedDiskMapData = []int{}

	var fileIndex = 0
	for index, digit := range this.data {
		var isFile bool = index%2 == 0
		var valueToWrite int
		if isFile {
			valueToWrite = fileIndex
			fileIndex++
		} else {
			valueToWrite = -1
		}
		expandedDiskMapData = append(expandedDiskMapData, repeatInSlice(valueToWrite, digit)...)
	}
	return ExpandedDiskMap{data: expandedDiskMapData}
}

func repeatInSlice(value int, length int) []int {
	var diskMapData = []int{}
	for range length {
		diskMapData = append(diskMapData, value)
	}
	return diskMapData
}
