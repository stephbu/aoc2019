package day11

import (
	"fmt"
)

var (
	Up    Coord
	Down  Coord
	Left  Coord
	Right Coord
)

func init() {
	Up = Coord{0, -1}
	Down = Coord{0, 1}
	Left = Coord{-1, 0}
	Right = Coord{1, 0}
}

type Turtle struct {
	Location      Coord
	CurrentVector Coord
	Panels        [][]*Panel
	PanelsPainted int
}

func (turtle *Turtle) MoveForward() {
	turtle.Location = turtle.Location.Add(turtle.CurrentVector)
}

func (turtle *Turtle) MoveBackwards() {
	turtle.Location = turtle.Location.Add(turtle.CurrentVector)
}

func (turtle *Turtle) SetDirection(coord Coord) {
	turtle.CurrentVector = coord
}

func (turtle *Turtle) SetLocation(coord Coord) {
	turtle.Location = coord
}

func (turtle *Turtle) GetPaint() int {
	panel := turtle.Panels[turtle.Location.Y][turtle.Location.X]
	if panel.MapChar == "." {
		return 0
	}

	return 1
}

func (turtle *Turtle) PaintTurnAndMove(paint int, turn int) {

	panel := turtle.Panels[turtle.Location.Y][turtle.Location.X]
	paintCount := panel.Paint(paint)

	if paintCount == 1 {
		turtle.PanelsPainted++
	}

	if turn == 0 {
		turtle.TurnLeft()
	} else {
		turtle.TurnRight()
	}

	turtle.MoveForward()

}

func (turtle *Turtle) String() string {
	return fmt.Sprintf("Turtle %v facing %v", turtle.Location, turtle.CurrentVector)
}

func (turtle *Turtle) DumpPanels() {

	for y := range turtle.Panels {
		for x := range turtle.Panels[y] {
			print(turtle.Panels[y][x].MapChar)
		}
		print("\r\n")
	}
}

func (turtle *Turtle) TurnLeft() {
	if turtle.CurrentVector == Up {
		turtle.CurrentVector = Left
	} else if turtle.CurrentVector == Left {
		turtle.CurrentVector = Down
	} else if turtle.CurrentVector == Down {
		turtle.CurrentVector = Right
	} else if turtle.CurrentVector == Right {
		turtle.CurrentVector = Up
	}
}

func (turtle *Turtle) TurnRight() {
	if turtle.CurrentVector == Up {
		turtle.CurrentVector = Right
	} else if turtle.CurrentVector == Right {
		turtle.CurrentVector = Down
	} else if turtle.CurrentVector == Down {
		turtle.CurrentVector = Left
	} else if turtle.CurrentVector == Left {
		turtle.CurrentVector = Up
	}
}
