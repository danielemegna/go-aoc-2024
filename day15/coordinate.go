package day15

// Coordinate represents a position in the warehouse map.
type Coordinate struct {
	Row int
	Col int
}

// NewCoordinate creates a new coordinate with the given row and column.
func NewCoordinate(row, col int) Coordinate {
	return Coordinate{Row: row, Col: col}
}

// Add returns a new coordinate by adding the given delta to this coordinate.
func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row + other.Row,
		Col: c.Col + other.Col,
	}
}

// Equals checks if two coordinates are the same.
func (c Coordinate) Equals(other Coordinate) bool {
	return c.Row == other.Row && c.Col == other.Col
}

// GPSCoordinate calculates the GPS coordinate value as described in the problem.
// It's 100 * (distance from top) + (distance from left).
// The distance is measured from the top-left corner (0,0), not counting the wall.
// For example, a box at (row=1, col=4) has a GPS of 100*1 + 4 = 104.
func (c Coordinate) GPSCoordinate() int {
	return 100*c.Row + c.Col
}