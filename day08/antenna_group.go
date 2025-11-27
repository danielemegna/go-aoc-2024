package day08

type AntennaGroup struct {
	locations []Coordinate
	antinodes Set[Coordinate]
	mapSize   int
}

func NewAntennaGroup(mapSize int) AntennaGroup {
	return AntennaGroup{
		locations: []Coordinate{},
		antinodes: NewSet[Coordinate](),
		mapSize:   mapSize,
	}
}

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
		if !firstAntinode.IsOutOfBounds(this.mapSize) {
			this.antinodes.Add(firstAntinode)
		}
		if !secondAntinode.IsOutOfBounds(this.mapSize) {
			this.antinodes.Add(secondAntinode)
		}
	}

	this.locations = append(this.locations, toAdd)
}

func (this *AntennaGroup) AddAntennaAtWithResonantHarmonics(toAdd Coordinate) {
	for _, alreadyPresent := range this.locations {
		this.antinodes.Add(toAdd)
		this.antinodes.Add(alreadyPresent)

		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y

		var isLeftAntinodeInBounds = true
		var isRightAntinodeInBounds = true
		for i := 1; isLeftAntinodeInBounds || isRightAntinodeInBounds; i++ {
			if isLeftAntinodeInBounds {
				var nextLeftAntinode = Coordinate{
					alreadyPresent.X - (xDifference * i),
					alreadyPresent.Y - (yDifference * i),
				}
				if nextLeftAntinode.IsOutOfBounds(this.mapSize) {
					isLeftAntinodeInBounds = false
				} else {
					this.antinodes.Add(nextLeftAntinode)
				}
			}

			if isRightAntinodeInBounds {
				var nextRightAntinode = Coordinate{
					toAdd.X + (xDifference * i),
					toAdd.Y + (yDifference * i),
				}
				if(nextRightAntinode.IsOutOfBounds(this.mapSize)) {
					isRightAntinodeInBounds = false
				} else {
					this.antinodes.Add(nextRightAntinode)
				}
			}
		}
	}

	this.locations = append(this.locations, toAdd)
}

func (this AntennaGroup) GetAntennas() []Coordinate {
	return this.locations
}

func (this AntennaGroup) GetAntinodes() []Coordinate {
	return this.antinodes.ToSlice()
}
