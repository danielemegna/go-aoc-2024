package day15

func BoxesGPSCoordinatesSumAfterAllRobotMoves(fileContent string) int {
	warehouse, moves := parseWarehouse(fileContent)
	for _, move := range moves {
		warehouse.Move(move)
	}
	return warehouse.GpsCoordinatesSum()
}

