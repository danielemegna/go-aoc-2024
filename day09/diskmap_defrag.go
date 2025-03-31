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
	var diskData = defragData(diskMap.data, len(diskMap.data)-1)
	return DenseDiskMap{data: diskData}
}

func defragData(diskData []any, bottomCursor int) []any {
	var selectedFileBlockIndex = LastFileBlockIndexIn(diskData, bottomCursor)
	if selectedFileBlockIndex == -1 {
		return diskData
	}

	var selectedFileBlock = diskData[selectedFileBlockIndex].(FileBlock)

	var selectedEmptyBlockIndex = FirstEmptyBlockIndexWith(selectedFileBlock.size, diskData)
	if selectedEmptyBlockIndex == -1 {
		return defragData(diskData, selectedFileBlockIndex-1)
	}
	if selectedEmptyBlockIndex > selectedFileBlockIndex {
		return defragData(diskData, selectedFileBlockIndex-1)
	}

	var selectedEmptyBlock = diskData[selectedEmptyBlockIndex].(EmptyBlock)
	var head = diskData[:selectedEmptyBlockIndex]

	var tail = []any{}
	if selectedFileBlockIndex < (len(diskData) - 2) {
		tail = diskData[selectedFileBlockIndex+2:]
	}

	var middleRest = []any{}
	var areBlocksClose = selectedFileBlockIndex-selectedEmptyBlockIndex < 2
	if !areBlocksClose {
		middleRest = diskData[selectedEmptyBlockIndex+1 : selectedFileBlockIndex-1]
	}

	var newLastEmptyBlockSize = selectedFileBlock.size
	if !areBlocksClose {
		newLastEmptyBlockSize += diskData[selectedFileBlockIndex-1].(EmptyBlock).size
	}

	if selectedFileBlockIndex+1 < len(diskData) {
		newLastEmptyBlockSize += diskData[selectedFileBlockIndex+1].(EmptyBlock).size
	}

	var movedFileBlockAndCloseSpaces = []any{
		EmptyBlock{size: 0},
		selectedFileBlock,
	}

	if !areBlocksClose {
		movedFileBlockAndCloseSpaces = append(
			movedFileBlockAndCloseSpaces,
			EmptyBlock{size: selectedEmptyBlock.size - selectedFileBlock.size},
		)
	}

	var newLastEmptyBlock = EmptyBlock{size: newLastEmptyBlockSize}
	middleRest = append(middleRest, newLastEmptyBlock)

	diskData = append(head,
		append(movedFileBlockAndCloseSpaces,
			append(middleRest,
				tail...,
			)...,
		)...,
	)

	return defragData(diskData, selectedFileBlockIndex)
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
