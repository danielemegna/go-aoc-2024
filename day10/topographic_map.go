package day10

type TopographicMap [][]int

type Trailhead struct {
	startingPosition             Coordinate
	reachableNineHeightPositions []Coordinate
	rate                         int
}

func (this Trailhead) GetRate() int {
	return this.rate
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
			if reachableNineHeightPositions.GetLength() > 0 {
				result = append(result, Trailhead{
					startingPosition,
					reachableNineHeightPositions.ToSlice(),
					reachableNineHeightPositions.GetLengthWithDuplicates(),
				})
			}

		}
	}

	return result
}

func (this TopographicMap) reachableNineHeightPositionsFrom(currentPosition Coordinate, currentHeight int) CoordinateSet {
	var result = NewCoordinateSet()
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
