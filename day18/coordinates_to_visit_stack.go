package day18

import "slices"

type CoordinatesToVisitStack []CoordinateWithCost

func (this *CoordinatesToVisitStack) PopFirstElement() CoordinateWithCost {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *CoordinatesToVisitStack) AppendSortedByCost(toAppend CoordinateWithCost) {
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
