package day09

func FilesystemChecksumAfterDefrag(denseDiskMapString string) int {
	var denseDiskMap = ParseDenseDiskMap(denseDiskMapString)
	var expandedDiskMap = denseDiskMap.ToExpandedDiskMap()
	var defraggedDiskMap = Defrag(expandedDiskMap)
	return defraggedDiskMap.Checksum()
}
