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
	var defraggedData = diskMap.data

	var diskMapLength = len(defraggedData)
	var lastFile = defraggedData[diskMapLength-1].(FileBlock)

	var emptyBlockIndex = diskMap.FirstEmptyBlockIndexWith(lastFile.size)
	var emptyBlock = defraggedData[emptyBlockIndex].(EmptyBlock)
	defraggedData[emptyBlockIndex] = EmptyBlock{size: emptyBlock.size - lastFile.size}
	// WIP

	return DenseDiskMap{data: defraggedData}
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
