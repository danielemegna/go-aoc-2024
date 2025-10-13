package day10

type TopographicMap [][]int

type Trailhead struct {
	startingPosition             Coordinate
	reachableNineHeightPositions []Coordinate
}

func (this Trailhead) GetScore() int {
	return len(this.reachableNineHeightPositions)
}

func (this TopographicMap) FindTrailheads() []Trailhead {
	var result = []Trailhead{}

	for y := 0; y < len(this); y++ {
		for x := 0; x < len(this[y]); x++ {
			var startingPosition = Coordinate{x, y}
			if this.valueAt(startingPosition) != 0 {
				continue
			}

			var reachableNineHeightPositions = this.reachableNineHeightPositionsFrom(startingPosition, 0)
			if len(reachableNineHeightPositions) > 0 {
				result = append(result, Trailhead{
					startingPosition,
					reachableNineHeightPositions.ToSlice(),
				})
			}

		}
	}

	return result
}

func (this TopographicMap) reachableNineHeightPositionsFrom(currentPosition Coordinate, currentHeight int) CoordinateSet {
	var result = CoordinateSet{}
	var targetHeigth = currentHeight + 1

	for _, closeCoordinate := range currentPosition.CloseCoordinates() {
		if closeCoordinate.IsOutOfBoundsOf(this) {
			continue
		}

		if this.valueAt(closeCoordinate) != targetHeigth {
			continue
		}

		if targetHeigth == 9 {
			result.Add(closeCoordinate)
			continue
		}

		result.Merge(this.reachableNineHeightPositionsFrom(closeCoordinate, currentHeight+1))
	}
	return result
}

func (this TopographicMap) valueAt(c Coordinate) int {
	return this[c.Y][c.X]
}
