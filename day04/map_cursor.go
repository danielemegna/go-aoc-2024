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

type InverseHorizontalMapCursor struct{ c Coordinate }

func (this *InverseHorizontalMapCursor) Increase()                { this.c.X-- }
func (this *InverseHorizontalMapCursor) ToCoordinate() Coordinate { return this.c }

type InverseVerticalMapCursor struct{ c Coordinate }

func (this *InverseVerticalMapCursor) Increase()                { this.c.Y-- }
func (this *InverseVerticalMapCursor) ToCoordinate() Coordinate { return this.c }

type NorthDiagonalMapCursor struct{ c Coordinate }

func (this *NorthDiagonalMapCursor) Increase()                { this.c.X++; this.c.Y-- }
func (this *NorthDiagonalMapCursor) ToCoordinate() Coordinate { return this.c }

type SouthDiagonalMapCursor struct{ c Coordinate }

func (this *SouthDiagonalMapCursor) Increase()                { this.c.X++; this.c.Y++ }
func (this *SouthDiagonalMapCursor) ToCoordinate() Coordinate { return this.c }

type NorthInverseDiagonalMapCursor struct{ c Coordinate }

func (this *NorthInverseDiagonalMapCursor) Increase()                { this.c.X--; this.c.Y-- }
func (this *NorthInverseDiagonalMapCursor) ToCoordinate() Coordinate { return this.c }

type SouthInverseDiagonalMapCursor struct{ c Coordinate }

func (this *SouthInverseDiagonalMapCursor) Increase()                { this.c.X--; this.c.Y++ }
func (this *SouthInverseDiagonalMapCursor) ToCoordinate() Coordinate { return this.c }
