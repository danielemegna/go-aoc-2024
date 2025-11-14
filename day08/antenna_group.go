package day08

type AntennaGroup struct {
	frequency rune
	locations []Coordinate
	antinodes []Coordinate
}

func (this *AntennaGroup) AddAntennaAt(toAdd Coordinate) {
	if len(this.locations) > 0 {
		var alreadyPresent = this.locations[0]
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y
		this.antinodes = append(this.antinodes, Coordinate{alreadyPresent.X - xDifference, alreadyPresent.Y - yDifference})
		this.antinodes = append(this.antinodes, Coordinate{toAdd.X + xDifference, toAdd.Y + yDifference})
	}

	this.locations = append(this.locations, toAdd)
}
