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
	var fileBlockToMoveIndex = LastFileBlockIndexIn(diskData, bottomCursor)
	var thereAreNoMoreFileBlockToCheck = fileBlockToMoveIndex == -1
	if thereAreNoMoreFileBlockToCheck {
		return diskData
	}

	var fileBlockToMove = diskData[fileBlockToMoveIndex].(FileBlock)

	var selectedEmptyBlockIndex = FirstEmptyBlockIndexIn(diskData, fileBlockToMove.size)
	var thereAreNoBigEnoughEmptyBlock = selectedEmptyBlockIndex == -1
	if thereAreNoBigEnoughEmptyBlock {
		return defragData(diskData, fileBlockToMoveIndex-1)
	}
	var isTheEmptyBlockBlockBeforeTheFileBlock = selectedEmptyBlockIndex < fileBlockToMoveIndex
	if !isTheEmptyBlockBlockBeforeTheFileBlock {
		return defragData(diskData, fileBlockToMoveIndex-1)
	}

	var selectedEmptyBlockSize = diskData[selectedEmptyBlockIndex].(EmptyBlock).size
	var areSelectedBlocksCloseBlocks = fileBlockToMoveIndex == selectedEmptyBlockIndex+1
	var isFileBlockToMoveLastBlock = fileBlockToMoveIndex == len(diskData)-1

	// ================ HEAD
	var head = diskData[:selectedEmptyBlockIndex]

	// ================ MOVED FILE
	var movedFileBlockAndCloseSpaces = []any{
		EmptyBlock{size: 0},
		fileBlockToMove,
	}

	if !areSelectedBlocksCloseBlocks {
		movedFileBlockAndCloseSpaces = append(
			movedFileBlockAndCloseSpaces,
			EmptyBlock{size: selectedEmptyBlockSize - fileBlockToMove.size},
		)
	}

	// ================ MIDDLE REST
	var middleRest = []any{}
	if !areSelectedBlocksCloseBlocks {
		middleRest = diskData[selectedEmptyBlockIndex+1 : fileBlockToMoveIndex-1]
	}

	var newFreedEmptyBlockSize = fileBlockToMove.size
	if !isFileBlockToMoveLastBlock {
		newFreedEmptyBlockSize += diskData[fileBlockToMoveIndex+1].(EmptyBlock).size
	}
	if !areSelectedBlocksCloseBlocks {
		newFreedEmptyBlockSize += diskData[fileBlockToMoveIndex-1].(EmptyBlock).size
	}
	var newFreedEmptyBlock = EmptyBlock{size: newFreedEmptyBlockSize}
	middleRest = append(middleRest, newFreedEmptyBlock)

	// ================ TAIL
	var tail = []any{}
	if fileBlockToMoveIndex < (len(diskData) - 2) {
		tail = diskData[fileBlockToMoveIndex+2:]
	}

	// ================ COMPOSE
	diskData = append(head, append(movedFileBlockAndCloseSpaces, append(middleRest, tail...)...)...)
	return defragData(diskData, fileBlockToMoveIndex)
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
