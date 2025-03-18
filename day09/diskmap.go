package day09

import (
	"github.com/samber/lo"
)

type DenseDiskMap struct {
	data []any
}

type FileBlock struct {
	size      int
	fileIndex int
}

type EmptyBlock struct {
	size int
}

type ExpandedDiskMap struct {
	data []int
}

func ParseDenseDiskMap(diskMapString string) DenseDiskMap {
	var diskMapRunes = []rune(diskMapString)
	var diskMapData = lo.Map(diskMapRunes, func(digit rune, index int) any {
		var blockSize = int(digit - '0')
		var isFile bool = index%2 == 0
		if isFile {
			return FileBlock{size: blockSize, fileIndex: index / 2}
		}
		return EmptyBlock{size: blockSize}
	})
	return DenseDiskMap{data: diskMapData}
}

func (this DenseDiskMap) ToExpandedDiskMap() ExpandedDiskMap {
	var expandedDiskMapData = []int{}

	for _, block := range this.data {
		var fileBlock, isFile = block.(FileBlock)
		var valueToWrite int
		var blockSize int
		if isFile {
			valueToWrite = fileBlock.fileIndex
			blockSize = fileBlock.size
		} else {
			var emptyBlock, _ = block.(EmptyBlock)
			valueToWrite = -1
			blockSize = emptyBlock.size
		}
		expandedDiskMapData = append(expandedDiskMapData, repeatInSlice(valueToWrite, blockSize)...)
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

func (this DenseDiskMap) LastFileBlockIndex() int {
	// optmiziation opportunity starting from something < len(this.data)
	for i := len(this.data) - 1; i >= 0; i-- {
		var block = this.data[i]
		var _, found = block.(FileBlock)
		if !found {
			continue
		}
		return i
	}

	return -1
}

func (this DenseDiskMap) FirstEmptyBlockIndexWith(minimumSize int) int {
	// optmiziation opportunity starting from something > 0
	for i := 0; i < len(this.data); i++ {
		var block = this.data[i]
		var emptyBlock, found = block.(EmptyBlock)
		if !found {
			continue
		}
		if emptyBlock.size >= minimumSize {
			return i
		}
	}

	return -1
}

func repeatInSlice(value int, length int) []int {
	var diskMapData = []int{}
	for range length {
		diskMapData = append(diskMapData, value)
	}
	return diskMapData
}
