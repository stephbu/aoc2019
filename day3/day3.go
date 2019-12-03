package day3

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func getDistanceFromSteps(manhattan bool, steps1 string, steps2 string) (int, error) {

	start := Coord{X: 0, Y: 0}

	path1, err := generatePathFromSteps(steps1, start)
	if err != nil {
		return 0, err
	}

	path2, err := generatePathFromSteps(steps2, start)
	if err != nil {
		return 0, err
	}

	return computeBestDistanceFromPaths(manhattan, path1, path2), nil
}

func intAbs(value int) int {
	return int(math.Abs(float64(value)))
}

func (coord *Coord) ManhattanDistance(other Coord) int {
	return intAbs(other.X-coord.X) + intAbs(other.Y-coord.Y)
}

func computeBestDistanceFromPaths(manhattan bool, pathA []Coord, pathB []Coord) int {

	min := 0
	last := 0
	start := pathA[0]

	for stepA, coordA := range pathA {
		for stepB, coordB := range pathB {
			if (coordA.X == coordB.X) && (coordA.Y == coordB.Y) {

				if manhattan {
					last = start.ManhattanDistance(coordA)
				} else {
					last = stepA + stepB
				}
				log.Printf("intersection at %v, distance %v", coordA, last)

				if min == 0 || last < min {
					min = last
				}
			}
		}
	}

	return min
}

func generatePathFromSteps(input string, start Coord) ([]Coord, error) {

	steps := strings.Split(input, ",")

	coords := make([]Coord, 0)
	coords = append(coords, start)

	last := start
	for _, step := range steps {

		count, err := strconv.Atoi(step[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid instruction %v", step)
		}

		switch step[0] {
		case 'U':
			{
				for ; count > 0; count-- {
					last.Y++
					coords = append(coords, last)
				}
			}
		case 'D':
			{
				for ; count > 0; count-- {
					last.Y--
					coords = append(coords, last)
				}
			}
		case 'L':
			{
				for ; count > 0; count-- {
					last.X--
					coords = append(coords, last)
				}
			}
		case 'R':
			{
				for ; count > 0; count-- {
					last.X++
					coords = append(coords, last)
				}
			}
		default:
			return nil, fmt.Errorf("invalid instruction %v", step)
		}

	}

	return coords, nil
}
