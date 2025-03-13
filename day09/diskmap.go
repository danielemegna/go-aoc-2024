package day09

import (
	"github.com/samber/lo"
)

type DenseDiskMap struct {
	data []int
}

type DenseDiskMapEvo struct {
	data []any
}

type FileBlock struct {
	size      int
	fileIndex int
}

type EmptyBlock struct {
	size int
}

type ExpandedDiskMap = DenseDiskMap

func ParseDenseDiskMapEvo(diskMapString string) DenseDiskMapEvo {
	var diskMapRunes = []rune(diskMapString)
	var diskMapData = lo.Map(diskMapRunes, func(digit rune, index int) any {
		var blockSize = int(digit - '0')
		var isFile bool = index%2 == 0
		if isFile {
			return FileBlock{size: blockSize, fileIndex: index / 2}
		}
		return EmptyBlock{size: blockSize}
	})
	return DenseDiskMapEvo{data: diskMapData}
}

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
	for digitIndex, digit := range this.data {
		var isFile bool = digitIndex%2 == 0
		var valueToWrite = -1
		if isFile {
			valueToWrite = fileIndex
			fileIndex++
		}
		expandedDiskMapData = append(expandedDiskMapData, repeatInSlice(valueToWrite, digit)...)
	}
	return ExpandedDiskMap{data: expandedDiskMapData}
}

func (this ExpandedDiskMap) Checksum() int {
	var sliceOfDigits = lo.DropRightWhile(this.data, func(digit int) bool {
		return digit == -1
	})
	return lo.Sum(lo.Map(sliceOfDigits, func(digitIndex int, index int) int {
		if digitIndex == -1 {
			return 0
		}
		return index * digitIndex
	}))
}

func repeatInSlice(value int, length int) []int {
	var diskMapData = []int{}
	for range length {
		diskMapData = append(diskMapData, value)
	}
	return diskMapData
}
