package day20

type Cheat struct {
	startPosition       Coordinate
	endPosition         Coordinate
	savingInPicoseconds int
}

func PossibleCheatsIn(racetrackMap RacetrackMap) []Cheat {
	var cheats []Cheat

	// Iterate through all track positions using the linked list
	current := &racetrackMap.start
	for current != nil {
		startCoord := current.coordinate
		startDistance := racetrackMap.ValueAt(startCoord)

		// Try all possible 2-step cheats from this position
		// A cheat can go in any direction for 2 steps
		directions := []Coordinate{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

		for _, dir := range directions {
			// First step (through wall or track)
			mid := Coordinate{startCoord.X + dir.X, startCoord.Y + dir.Y}

			// Second step (continue in same direction)
			end := Coordinate{mid.X + dir.X, mid.Y + dir.Y}

			// Check if end position is valid and on track
			if end.Y >= 0 && end.Y < len(racetrackMap.values) &&
				end.X >= 0 && end.X < len(racetrackMap.values[end.Y]) {

				endValue := racetrackMap.ValueAt(end)

				// Valid cheat: end must be on track (not a wall) and further ahead
				if endValue != WALL && endValue > startDistance {
					saving := endValue - startDistance - 2
					if saving > 0 {
						cheats = append(cheats, Cheat{
							startPosition:       startCoord,
							endPosition:         end,
							savingInPicoseconds: saving,
						})
					}
				}
			}
		}

		current = current.next
	}

	return cheats
}
