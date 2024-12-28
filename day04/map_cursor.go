package day04

type MapCursor struct {
	coordinate Coordinate
	increaseFn func(c Coordinate) Coordinate
}

func (this *MapCursor) Increase() {
	this.coordinate = this.increaseFn(this.coordinate)
}
