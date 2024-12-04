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

func (this NumberList) SimilarityScoreWith(other NumberList) int {
	if(this.elements[0] == other.elements[0]) {
		return this.elements[0]
	}

	return 0
}

func (this NumberList) OccurenceCount(target int) int {
	var count = 0

	// really there is no other way !?
	for _, element := range this.elements {
		if(element == target){
			count++
		}
	}

	return count
}
