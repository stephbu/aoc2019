package day11

import (
	"log"

	"github.com/stephbu/aoc2019/util"
)

type Panel struct {
	PaintCount int
	MapChar    string
}

func (panel *Panel) Paint(paint int) int {
	panel.PaintCount++
	if paint == 0 {
		panel.MapChar = "."
	} else {
		panel.MapChar = "#"
	}

	return panel.PaintCount
}

func NewPanelsFromFile(filename string) [][]*Panel {
	lines, err := util.ReadLines(filename)
	if err != nil {
		log.Fatal(err)
	}

	y := len(lines)
	x := len(lines[0])

	panels := make([][]*Panel, y)
	for i := range panels {
		panels[i] = make([]*Panel, x)
	}

	for indexY := range lines {
		for indexX := range lines[indexY] {
			panels[indexY][indexX] = &Panel{MapChar: string(lines[indexY][indexX]), PaintCount: 0}
		}
	}

	return panels
}
