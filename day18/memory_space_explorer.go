package day18

import "slices"

type CoordinateWithCost struct {
	coordinate Coordinate
	cost       int
	parent     *CoordinateWithCost
}

func (this CoordinateWithCost) PathCoordinates() []Coordinate {
	var result = []Coordinate{}
	var currentNode *CoordinateWithCost = &this
	for {
		if currentNode == nil {
			break
		}

		result = append(result, currentNode.coordinate)
		currentNode = currentNode.parent
	}

	return result
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
			{coordinate: Coordinate{0, 0}, cost: 0, parent: nil},
		},
		visited: []Coordinate{},
	}
	return explorer
}

func (this MemorySpaceExplorer) ShortestPathFromTopLeftToBottomRight() []Coordinate {
	var targetCoordinate = this.memorySpace.BottomRightCoordinate()
	for len(this.toVisitStack) > 0 {
		var toVisit = this.toVisitStack.PopFirstElement()

		if toVisit.coordinate == targetCoordinate {
			return toVisit.PathCoordinates()
		}

		for _, closeCoordinate := range toVisit.coordinate.CloseCoordinates() {
			var closeCoordinateWithCost = CoordinateWithCost{closeCoordinate, toVisit.cost + 1, &toVisit}
			this.appendToStackWithCostIfVisitable(closeCoordinateWithCost)
		}

		this.visited = append(this.visited, toVisit.coordinate)
	}

	return []Coordinate{}
}

func (this *MemorySpaceExplorer) appendToStackWithCostIfVisitable(toAppend CoordinateWithCost) {
	if !this.memorySpace.IsVisitable(toAppend.coordinate) {
		return
	}

	if this.alreadyVisited(toAppend.coordinate) {
		return
	}

	this.toVisitStack.AppendSortedByCost(toAppend)
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
	var stack = (*this)
	for index, existingElement := range stack {
		if existingElement.coordinate == toAppend.coordinate {
			return
		}
		if toAppend.cost < existingElement.cost {
			(*this) = slices.Insert(stack, index, toAppend)
			return
		}
	}

	(*this) = append(stack, toAppend)
}
