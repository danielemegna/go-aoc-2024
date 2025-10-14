package day10

type CoordinateSet struct {
	elements      map[Coordinate]bool
	addOperations int
}

func NewCoordinateSet() CoordinateSet {
	return CoordinateSet{
		elements:      map[Coordinate]bool{},
		addOperations: 0,
	}
}

func (this *CoordinateSet) Add(c Coordinate) {
	this.elements[c] = true
	this.addOperations++
}

func (this *CoordinateSet) Merge(other CoordinateSet) {
	for c := range other.elements {
		this.elements[c] = true
	}
	this.addOperations += other.addOperations
}

func (this CoordinateSet) ToSlice() []Coordinate {
	var result = []Coordinate{}
	for c := range this.elements {
		result = append(result, c)
	}
	return result
}

func (this CoordinateSet) GetLength() int {
	return len(this.elements)
}

func (this CoordinateSet) GetLengthWithDuplicates() int {
	return this.addOperations
}
