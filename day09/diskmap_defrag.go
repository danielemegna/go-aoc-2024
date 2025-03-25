package day09

func Defrag(diskMap ExpandedDiskMap) ExpandedDiskMap {
	var newData = diskMap.data

	var topCursor int = 0
	var bottomCursor int = len(diskMap.data) - 1

	for {
		var firstEmptySpaceIndex int
		var lastNotEmptyDigitIndex int
		var found bool

		firstEmptySpaceIndex, found = findIndexOf(newData, func(digit int) bool { return digit == -1 }, topCursor, false)
		if !found {
			break
		}

		lastNotEmptyDigitIndex, found = findIndexOf(newData, func(digit int) bool { return digit != -1 }, bottomCursor, true)
		if !found {
			break
		}

		if firstEmptySpaceIndex >= lastNotEmptyDigitIndex {
			break
		}

		newData[firstEmptySpaceIndex] = newData[lastNotEmptyDigitIndex]
		newData[lastNotEmptyDigitIndex] = -1
		topCursor = firstEmptySpaceIndex
		bottomCursor = lastNotEmptyDigitIndex
	}

	return ExpandedDiskMap{data: newData}
}

func DefragWholeFiles(diskMap DenseDiskMap) DenseDiskMap {
	var diskData = diskMap.data

	var selectedFileBlockIndex = diskMap.LastFileBlockIndex()
	var selectedFileBlock = diskData[selectedFileBlockIndex].(FileBlock)

	var selectedEmptyBlockIndex = diskMap.FirstEmptyBlockIndexWith(selectedFileBlock.size)
	if selectedEmptyBlockIndex > selectedFileBlockIndex {
		return diskMap
	}

	var selectedEmptyBlock = diskData[selectedEmptyBlockIndex].(EmptyBlock)
	var movedFileBlockAndCloseSpaces = []any{
		EmptyBlock{size: 0},
		selectedFileBlock,
		EmptyBlock{size: selectedEmptyBlock.size - selectedFileBlock.size},
	}

	var middleRest = diskData[selectedEmptyBlockIndex+1 : selectedFileBlockIndex-1]

	var newLastEmptyBlockSize = diskData[selectedFileBlockIndex-1].(EmptyBlock).size + selectedFileBlock.size
	if selectedFileBlockIndex+1 < len(diskData) {
		newLastEmptyBlockSize += diskData[selectedFileBlockIndex+1].(EmptyBlock).size
	}

	var newLastEmptyBlock = EmptyBlock{size: newLastEmptyBlockSize}

	diskData = append(diskData[:selectedEmptyBlockIndex],
		append(movedFileBlockAndCloseSpaces,
			append(middleRest, newLastEmptyBlock)...,
		)...,
	)

	return DefragWholeFiles(DenseDiskMap{data: diskData})
}

func findIndexOf(collection []int, predicate func(item int) bool, startIndex int, reverse bool) (int, bool) {
	if reverse {
		for i := startIndex; i >= 0; i-- {
			if predicate(collection[i]) {
				return i, true
			}
		}

		return -1, false
	}

	for i := startIndex; i < len(collection); i++ {
		if predicate(collection[i]) {
			return i, true
		}
	}

	return -1, false
}
