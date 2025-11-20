package day08

type AntennaGroup struct {
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

func (this *AntennaGroup) AddAntennaAtWithResonantHarmonics(toAdd Coordinate, mapSize int) {
	for _, alreadyPresent := range this.locations {
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y

		var continueLoop = true
		for i := 0; continueLoop; i++ {
			var firstAntinode = Coordinate{
				alreadyPresent.X - (xDifference * i),
				alreadyPresent.Y - (yDifference * i),
			}
			var secondAntinode = Coordinate{
				toAdd.X + (xDifference * i),
				toAdd.Y + (yDifference * i),
			}

			continueLoop = false
			if !firstAntinode.IsOutOfBounds(mapSize) {
				this.antinodes = append(this.antinodes, firstAntinode)
				continueLoop = true
			}
			if !secondAntinode.IsOutOfBounds(mapSize) {
				this.antinodes = append(this.antinodes, secondAntinode)
				continueLoop = true
			}
		}

	}

	this.locations = append(this.locations, toAdd)
}
