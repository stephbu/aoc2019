package day11

type Coord struct {
	X int
	Y int
}

func (coord *Coord) Add(other Coord) Coord {
	return Coord{X: other.X + coord.X, Y: other.Y + coord.Y}
}

func (coord *Coord) Subtract(other Coord) Coord {
	return Coord{X: other.X - coord.X, Y: other.Y - coord.Y}
}

/*
func (coord *Coord) String() string {
	return fmt.Sprintf("(%d,%d)", coord.X, coord.Y)
}*/
