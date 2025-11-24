package day08

type AntennaGroup struct {
	locations []Coordinate
	antinodes []Coordinate // maybe a set (map with no duplicates) would be better
	mapSize   int
}

// with a map size field we can avoid outofbound antinodes
func (this *AntennaGroup) AddAntennaAt(toAdd Coordinate) {
	for _, alreadyPresent := range this.locations {
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y
		var firstAntinode = Coordinate{
			alreadyPresent.X - xDifference,
			alreadyPresent.Y - yDifference,
		}
		var secondAntinode = Coordinate{
			toAdd.X + xDifference,
			toAdd.Y + yDifference,
		}
		this.antinodes = append(this.antinodes, firstAntinode)
		this.antinodes = append(this.antinodes, secondAntinode)
	}

	this.locations = append(this.locations, toAdd)
}

// if we inject map size in construction we could avoid the mapSize param
func (this *AntennaGroup) AddAntennaAtWithResonantHarmonics(toAdd Coordinate, mapSize int) {

	// with a set (map with no duplicates) antinodes field we can avoid these checks
	if len(this.locations) > 0 {
		this.antinodes = append(this.antinodes, toAdd)
		if len(this.locations) == 1 {
			this.antinodes = append(this.antinodes, this.locations[0])
		}
	}

	for _, alreadyPresent := range this.locations {
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y

		var continueLoop = true
		for i := 1; continueLoop; i++ {
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
