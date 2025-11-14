package day15

type Warehouse struct {
	Width, Height int
	Walls         map[Coordinate]bool
	Boxes         map[Coordinate]bool
	Robot         Coordinate
}

func (w *Warehouse) Move(direction rune) {
	var dx, dy int
	switch direction {
	case '^':
		dy = -1
	case 'v':
		dy = 1
	case '<':
		dx = -1
	case '>':
		dx = 1
	}

	newRobotPos := Coordinate{X: w.Robot.X + dx, Y: w.Robot.Y + dy}

	if w.Walls[newRobotPos] {
		return
	}

	if w.Boxes[newRobotPos] {
		newBoxPos := Coordinate{X: newRobotPos.X + dx, Y: newRobotPos.Y + dy}
		if w.Walls[newBoxPos] || w.Boxes[newBoxPos] {
			return
		}
		delete(w.Boxes, newRobotPos)
		w.Boxes[newBoxPos] = true
	}

	w.Robot = newRobotPos
}

func (w *Warehouse) GpsCoordinatesSum() int {
	sum := 0
	for pos := range w.Boxes {
		sum += 100*pos.Y + pos.X
	}
	return sum
}
