package day15

// Cell types in the warehouse map
const (
	Wall   = '#'
	Empty  = '.'
	Box    = 'O'
	Robot  = '@'
)

// Direction constants for robot movement
const (
	Up    = '^'
	Down  = 'v'
	Left  = '<'
	Right = '>'
)

// DirectionToCoordinate maps a direction character to a coordinate delta
var DirectionToCoordinate = map[rune]Coordinate{
	Up:    {Row: -1, Col: 0},
	Down:  {Row: 1, Col: 0},
	Left:  {Row: 0, Col: -1},
	Right: {Row: 0, Col: 1},
}

// WarehouseMap represents the warehouse with the robot and boxes.
type WarehouseMap struct {
	Grid       [][]rune
	RobotPos   Coordinate
	BoxesPos   map[Coordinate]bool
}

// NewWarehouseMap creates a new empty warehouse map.
func NewWarehouseMap() *WarehouseMap {
	return &WarehouseMap{
		Grid:     make([][]rune, 0),
		BoxesPos: make(map[Coordinate]bool),
	}
}

// AddRow adds a row to the warehouse map grid.
func (w *WarehouseMap) AddRow(row string) {
	rowRunes := []rune(row)
	w.Grid = append(w.Grid, rowRunes)
	
	// Detect robot and boxes in this row
	for col, r := range rowRunes {
		pos := Coordinate{Row: len(w.Grid) - 1, Col: col}
		if r == Robot {
			w.RobotPos = pos
			// The robot's position is actually an empty space in the grid
			w.Grid[pos.Row][pos.Col] = Empty
		} else if r == Box {
			w.BoxesPos[pos] = true
		}
	}
}

// IsWithinBounds checks if a coordinate is within the bounds of the grid.
func (w *WarehouseMap) IsWithinBounds(pos Coordinate) bool {
	return pos.Row >= 0 && pos.Row < len(w.Grid) && pos.Col >= 0 && pos.Col < len(w.Grid[0])
}

// GetCell returns the cell type at the given position.
// If the position is the robot's position, it returns Robot.
// If the position is a box's position, it returns Box.
// Otherwise, it returns the cell from the grid.
func (w *WarehouseMap) GetCell(pos Coordinate) rune {
	if pos.Equals(w.RobotPos) {
		return Robot
	}
	if w.BoxesPos[pos] {
		return Box
	}
	return w.Grid[pos.Row][pos.Col]
}

// IsWall checks if the cell at the given position is a wall.
func (w *WarehouseMap) IsWall(pos Coordinate) bool {
	return w.Grid[pos.Row][pos.Col] == Wall
}

// IsBox checks if there is a box at the given position.
func (w *WarehouseMap) IsBox(pos Coordinate) bool {
	return w.BoxesPos[pos]
}

// MoveRobot attempts to move the robot in the given direction.
// It returns true if the move was successful, false otherwise.
func (w *WarehouseMap) MoveRobot(direction rune) bool {
	delta := DirectionToCoordinate[direction]
	newPos := w.RobotPos.Add(delta)
	
	// Check if the new position is within bounds
	if !w.IsWithinBounds(newPos) {
		return false
	}
	
	// If the new position is a wall, the robot can't move
	if w.IsWall(newPos) {
		return false
	}
	
	// If the new position has a box, check if the box can be pushed
	if w.IsBox(newPos) {
		boxNewPos := newPos.Add(delta)
		
		// Check if the box's new position is within bounds
		if !w.IsWithinBounds(boxNewPos) {
			return false
		}
		
		// If the box's new position is a wall or another box, the box can't be pushed
		if w.IsWall(boxNewPos) || w.IsBox(boxNewPos) {
			return false
		}
		
		// Push the box
		delete(w.BoxesPos, newPos)
		w.BoxesPos[boxNewPos] = true
	}
	
	// Move the robot
	w.RobotPos = newPos
	return true
}

// GetBoxPositions returns a slice of all box positions.
func (w *WarehouseMap) GetBoxPositions() []Coordinate {
	positions := make([]Coordinate, 0, len(w.BoxesPos))
	for pos := range w.BoxesPos {
		positions = append(positions, pos)
	}
	return positions
}

// CalculateBoxesGPSSum calculates the sum of GPS coordinates for all boxes.
func (w *WarehouseMap) CalculateBoxesGPSSum() int {
	sum := 0
	for pos := range w.BoxesPos {
		sum += pos.GPSCoordinate()
	}
	return sum
}

// String returns a string representation of the warehouse map.
func (w *WarehouseMap) String() string {
	result := ""
	for i, row := range w.Grid {
		for j, cell := range row {
			pos := Coordinate{Row: i, Col: j}
			if pos.Equals(w.RobotPos) {
				result += string(Robot)
			} else if w.BoxesPos[pos] {
				result += string(Box)
			} else {
				result += string(cell)
			}
		}
		result += "\n"
	}
	return result
}