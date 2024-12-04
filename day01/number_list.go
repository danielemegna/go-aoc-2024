package day01

import "sort"

type NumberList struct {
	elements []int
}

func (this NumberList) Sort() {
	sort.Ints(this.elements)
}

func (this NumberList) Length() int {
	return len(this.elements)
}

func (this NumberList) At(index int) int {
	return this.elements[index]
}
