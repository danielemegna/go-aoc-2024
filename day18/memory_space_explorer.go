package day18

import "slices"

type MemorySpaceExplorer struct {
	memorySpace       MemorySpace
	currentCoordinate Coordinate
	toVisitStack      CoordinatesToVisitStack
	visited           []Coordinate
}

func NewMemorySpaceExplorer(memorySpace MemorySpace) MemorySpaceExplorer {
	return MemorySpaceExplorer{
		memorySpace:       memorySpace,
		currentCoordinate: Coordinate{0, 0},
		toVisitStack: []CoordinateInPath{
			{coordinate: Coordinate{0, 0}, parent: nil, pathLength: 0},
		},
		visited: []Coordinate{},
	}
}

func (this MemorySpaceExplorer) ShortestPathFromTopLeftToBottomRight() []Coordinate {
	var targetCoordinate = this.memorySpace.BottomRightCoordinate()
	for len(this.toVisitStack) > 0 {
		var toVisit = this.toVisitStack.PopFirstElement()

		if toVisit.coordinate == targetCoordinate {
			return toVisit.PathCoordinates()
		}

		for _, closeCoordinate := range toVisit.coordinate.CloseCoordinates() {
			var closeCoordinateInPath = LinkCoordinateInPath(closeCoordinate, &toVisit)
			this.appendToStackIfVisitable(closeCoordinateInPath)
		}

		this.visited = append(this.visited, toVisit.coordinate)
	}

	return []Coordinate{}
}

func (this *MemorySpaceExplorer) appendToStackIfVisitable(toAppend CoordinateInPath) {
	if !this.memorySpace.IsVisitable(toAppend.coordinate) {
		return
	}

	if this.alreadyVisited(toAppend.coordinate) {
		return
	}

	this.toVisitStack.AppendSortedByPathLength(toAppend)
}

func (this *MemorySpaceExplorer) alreadyVisited(c Coordinate) bool {
	return slices.Contains(this.visited, c)
}
