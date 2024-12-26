package day04

type Coordinate struct {
	X int
	Y int
}

func (this Coordinate) NorthEast() Coordinate {
	return Coordinate{this.X + 1, this.Y - 1}
}

func (this Coordinate) NorthOvest() Coordinate {
	return Coordinate{this.X - 1, this.Y - 1}
}

func (this Coordinate) SouthEast() Coordinate {
	return Coordinate{this.X + 1, this.Y + 1}
}

func (this Coordinate) SouthOvest() Coordinate {
	return Coordinate{this.X - 1, this.Y + 1}
}
