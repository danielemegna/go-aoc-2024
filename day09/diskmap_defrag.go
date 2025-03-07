package day09

import (
	"github.com/samber/lo"
)

func Defrag(diskMap ExpandedDiskMap) ExpandedDiskMap {
	var newData = diskMap.data

	for {
		var found bool
		var firstEmptySpaceIndex int
		var lastNotEmptyDigitIndex int

		_, firstEmptySpaceIndex, found = lo.FindIndexOf(newData, func(digit int) bool { return digit == -1 })
		if !found {
			break
		}

		_, lastNotEmptyDigitIndex, found = lo.FindLastIndexOf(newData, func(digit int) bool { return digit != -1 })
		if !found {
			break
		}

		if firstEmptySpaceIndex >= lastNotEmptyDigitIndex {
			break
		}

		newData[firstEmptySpaceIndex] = newData[lastNotEmptyDigitIndex]
		newData[lastNotEmptyDigitIndex] = -1
	}

	return ExpandedDiskMap{data: newData}
}
