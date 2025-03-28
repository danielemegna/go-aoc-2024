package day09

func FilesystemChecksumAfterDefrag(denseDiskMapString string) int {
	var denseDiskMap = ParseDenseDiskMap(denseDiskMapString)
	var expandedDiskMap = denseDiskMap.ToExpandedDiskMap()
	var defraggedDiskMap = Defrag(expandedDiskMap)
	return defraggedDiskMap.Checksum()
}

func FilesystemChecksumAfterDefragWholeFiles(denseDiskMapString string) int {
	var denseDiskMap = ParseDenseDiskMap(denseDiskMapString)
	var defraggedDiskMap = DefragWholeFiles(denseDiskMap)
	return defraggedDiskMap.ToExpandedDiskMap().Checksum()
}
