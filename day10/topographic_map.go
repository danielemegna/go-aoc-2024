package day10

type TopographicMap [][]int

type Coordinate struct {
	X int
	Y int
}

type Trailhead struct {
	startingPosition Coordinate
	reachableNineHeightPositions []Coordinate
}

func (this TopographicMap) FindTrailheads() []Trailhead {
	return []Trailhead{}
}
