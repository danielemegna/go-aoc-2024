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

	var lastFileBlockIndex = diskMap.LastFileBlockIndex()
	var lastFileBlock = diskData[lastFileBlockIndex].(FileBlock)

	var firstEmptyBlockIndex = diskMap.FirstEmptyBlockIndexWith(lastFileBlock.size)
	var firstEmptyBlock = diskData[firstEmptyBlockIndex].(EmptyBlock)

	if firstEmptyBlockIndex > lastFileBlockIndex {
		return diskMap
	}

	var movedFileBlockAndCloseSpaces = []any{
		EmptyBlock{size: 0},
		lastFileBlock,
		EmptyBlock{size: firstEmptyBlock.size - lastFileBlock.size},
	}

	var middleRest = diskData[firstEmptyBlockIndex+1 : lastFileBlockIndex-1]

	var newLastEmptyBlockSize = diskData[lastFileBlockIndex-1].(EmptyBlock).size + lastFileBlock.size
	if lastFileBlockIndex+1 < len(diskData) {
		newLastEmptyBlockSize += diskData[lastFileBlockIndex+1].(EmptyBlock).size
	}

	var newLastEmptyBlock = EmptyBlock{size: newLastEmptyBlockSize}

	diskData = append(diskData[:firstEmptyBlockIndex],
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
