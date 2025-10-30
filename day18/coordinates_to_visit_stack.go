package day18

import "slices"

type CoordinatesToVisitStack []CoordinateInPath

func (this *CoordinatesToVisitStack) PopFirstElement() CoordinateInPath {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *CoordinatesToVisitStack) AppendSortedByPathLength(toAppend CoordinateInPath) {
	var stack = (*this)
	for index, existingElement := range stack {
		if existingElement.coordinate == toAppend.coordinate {
			return
		}
		if toAppend.pathLength < existingElement.pathLength {
			(*this) = slices.Insert(stack, index, toAppend)
			return
		}
	}

	(*this) = append(stack, toAppend)
}
