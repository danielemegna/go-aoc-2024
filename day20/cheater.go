package day20

type Cheat struct {
	startPosition       Coordinate
	endPosition         Coordinate
	savingInPicoseconds int
}

func PossibleCheatsIn(racetrackMap RacetrackMap) []Cheat {
	var result = []Cheat{}

	var currentRacetrackPosition = racetrackMap.RacetrackStart()
	for currentRacetrackPosition != nil {
		var currentRacetrackValue = racetrackMap.ValueAt(currentRacetrackPosition.Coordinate)

		for _, firstStepCheat := range currentRacetrackPosition.Coordinate.CloseCoordinates() {
			if racetrackMap.ValueAt(firstStepCheat) != WALL {
				continue
			}

			for _, secondStepCheat := range firstStepCheat.CloseCoordinates() {
				if secondStepCheat.IsOutOfBounds(racetrackMap.MapSize()) {
					continue
				}
				if secondStepCheat == currentRacetrackPosition.Coordinate {
					continue
				}
				var secondStepValue = racetrackMap.ValueAt(secondStepCheat)
				if secondStepValue == WALL {
					continue
				}
				var savedPicoseconds = secondStepValue - currentRacetrackValue - 2
				if savedPicoseconds < 1 {
					continue
				}

				result = append(result, Cheat{
					startPosition:       currentRacetrackPosition.Coordinate,
					endPosition:         secondStepCheat,
					savingInPicoseconds: savedPicoseconds,
				})
			}
		}

		currentRacetrackPosition = currentRacetrackPosition.Next
	}

	return result
}

func PossibleLongCheatsIn(racetrackMap RacetrackMap, maxCheatDuration int) []Cheat {
	var result = []Cheat{}
	
	// Optimization: Collect all track positions into a slice for O(1) indexed access
	// This avoids repeated linked list traversals
	var trackPositions = []RacetrackElement{}
	var current = racetrackMap.RacetrackStart()
	for current != nil {
		trackPositions = append(trackPositions, *current)
		current = current.Next
	}
	
	// For each track position, check all other track positions
	// Complexity: O(n²) where n is track length, but this is unavoidable as we need
	// to check all pairs. The optimization is in early termination when Manhattan
	// distance exceeds maxCheatDuration.
	for i := 0; i < len(trackPositions); i++ {
		startPos := trackPositions[i]
		startDistance := racetrackMap.ValueAt(startPos.Coordinate)
		
		// Only check positions ahead in the track (j > i) to avoid duplicates
		// and ensure we're moving forward in time (endDistance > startDistance)
		for j := i + 1; j < len(trackPositions); j++ {
			endPos := trackPositions[j]
			endDistance := racetrackMap.ValueAt(endPos.Coordinate)
			
			// Calculate Manhattan distance between positions
			// This represents the minimum time needed to cheat from start to end
			manhattanDist := startPos.Coordinate.ManhattanDistanceTo(endPos.Coordinate)
			
			// Check if this is a valid cheat (within max duration)
			if manhattanDist <= maxCheatDuration {
				// Time saved = (normal path time from start to end) - (cheat duration)
				// = (endDistance - startDistance) - manhattanDist
				savedPicoseconds := endDistance - startDistance - manhattanDist
				
				if savedPicoseconds > 0 {
					result = append(result, Cheat{
						startPosition:       startPos.Coordinate,
						endPosition:         endPos.Coordinate,
						savingInPicoseconds: savedPicoseconds,
					})
				}
			}
		}
	}
	
	return result
}
