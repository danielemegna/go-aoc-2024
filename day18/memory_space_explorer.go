package day18

import "slices"

type CoordinateWithCost struct {
	coordinate Coordinate
	cost       int
}

type MemorySpaceExplorer struct {
	memorySpace       MemorySpace
	currentCoordinate Coordinate
	toVisitStack      ToVisitStack
	visited           []Coordinate
}

type ToVisitStack []CoordinateWithCost

func NewMemorySpaceExplorer(memorySpace MemorySpace) MemorySpaceExplorer {
	var explorer = MemorySpaceExplorer{
		memorySpace:       memorySpace,
		currentCoordinate: Coordinate{0, 0},
		toVisitStack: []CoordinateWithCost{
			{coordinate: Coordinate{0, 1}, cost: 1},
			{coordinate: Coordinate{1, 0}, cost: 1},
		},
		visited: []Coordinate{{0, 0}},
	}
	return explorer
}

func (this MemorySpaceExplorer) ShortestPathFromTopLeftToBottomRight() int {
	var targetCoordinate = this.memorySpace.BottomRightCoordinate()
	for len(this.toVisitStack) > 0 {
		var toVisit = this.toVisitStack.PopFirstElement()

		if toVisit.coordinate == targetCoordinate {
			return toVisit.cost
		}

		for _, closeCoordinate := range toVisit.coordinate.CloseCoordinates() {
			this.appendToStackWithCostIfVisitable(closeCoordinate, toVisit.cost+1)
		}

		this.visited = append(this.visited, toVisit.coordinate)
	}

	return -1
}

func (this *MemorySpaceExplorer) appendToStackWithCostIfVisitable(coordinate Coordinate, cost int) {
	if !this.memorySpace.IsVisitable(coordinate) {
		return
	}

	if this.alreadyVisited(coordinate) {
		return
	}

	var coordinateWithCost = CoordinateWithCost{coordinate, cost}
	this.toVisitStack.AppendSortedByCost(coordinateWithCost)
}

func (this *MemorySpaceExplorer) alreadyVisited(c Coordinate) bool {
	return slices.Contains(this.visited, c)
}

func (this *ToVisitStack) PopFirstElement() CoordinateWithCost {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *ToVisitStack) AppendSortedByCost(toAppend CoordinateWithCost) {
	var cost = toAppend.cost
	var stack = (*this)
	for index, existingElement := range stack {
		if cost <= existingElement.cost {
			(*this) = slices.Insert(stack, index, toAppend)
			return
		}
	}

	(*this) = append(stack, toAppend)
}
