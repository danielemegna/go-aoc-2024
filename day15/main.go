package day15

func BoxesGPSCoordinatesSum(fileContent string) int {
	var warehouseMap, robotMoves = ParseWarehouseMapAndMoves(fileContent)

	for _, move := range robotMoves {
		warehouseMap.MoveRobot(move)
	}

	return warehouseMap.GetBoxesGPSCoordinatesSum()
}
