package day04

type MapCursor interface {
	Increase()
	ToCoordinate() Coordinate
}

type HorizontalMapCursor struct{ c Coordinate }

func (this *HorizontalMapCursor) Increase()                { this.c.X++ }
func (this *HorizontalMapCursor) ToCoordinate() Coordinate { return this.c }

type VerticalMapCursor struct{ c Coordinate }

func (this *VerticalMapCursor) Increase()                { this.c.Y++ }
func (this *VerticalMapCursor) ToCoordinate() Coordinate { return this.c }
