package day15

func BoxesGPSCoordinatesSumAfterAllRobotMoves(fileContent string) int {
	// Parse the input
	parser := NewWarehouseParser()
	warehouseMap, moves := parser.Parse(fileContent)
	
	// Execute all the robot moves
	for i, move := range moves {
		success := warehouseMap.MoveRobot(move)
		if !success {
			// This is just for debugging to see which move failed
			// fmt.Printf("Move %d (%c) failed\n", i, move)
			_ = i // avoid unused variable warning
		}
	}
	
	// Calculate and return the sum of GPS coordinates for all boxes
	return warehouseMap.CalculateBoxesGPSSum()
}

