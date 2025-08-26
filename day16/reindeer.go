package day16

type Reindeer struct {
	Coordinate Coordinate
	Direction  Direction
}

func (this Reindeer) ForehandCoordinate() Coordinate {
	return this.Coordinate.NextFor(this.Direction)
}

func (this Reindeer) OppositeCoordinate() Coordinate {
	return this.Coordinate.NextFor(this.Direction.Opposite())
}
