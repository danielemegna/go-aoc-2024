package day08

type AntennaGroup struct {
	frequency rune
	locations []Coordinate
	antinodes []Coordinate
}

func (this *AntennaGroup) AddAntennaAt(toAdd Coordinate) {
	for _, alreadyPresent := range this.locations {
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y
		this.antinodes = append(this.antinodes, Coordinate{alreadyPresent.X - xDifference, alreadyPresent.Y - yDifference})
		this.antinodes = append(this.antinodes, Coordinate{toAdd.X + xDifference, toAdd.Y + yDifference})
	}

	this.locations = append(this.locations, toAdd)
}
