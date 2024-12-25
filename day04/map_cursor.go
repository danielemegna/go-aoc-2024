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

type DiagonalMapCursor struct{ c Coordinate }

func (this *DiagonalMapCursor) Increase()                { this.c.X++; this.c.Y++ }
func (this *DiagonalMapCursor) ToCoordinate() Coordinate { return this.c }

type InverseHorizontalMapCursor struct{ c Coordinate }

func (this *InverseHorizontalMapCursor) Increase()                { this.c.X-- }
func (this *InverseHorizontalMapCursor) ToCoordinate() Coordinate { return this.c }

type InverseVerticalMapCursor struct{ c Coordinate }

func (this *InverseVerticalMapCursor) Increase()                { this.c.Y-- }
func (this *InverseVerticalMapCursor) ToCoordinate() Coordinate { return this.c }

type InverseDiagonalMapCursor struct{ c Coordinate }

func (this *InverseDiagonalMapCursor) Increase()                { this.c.X--; this.c.Y-- }
func (this *InverseDiagonalMapCursor) ToCoordinate() Coordinate { return this.c }
