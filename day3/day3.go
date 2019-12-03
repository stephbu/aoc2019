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

func getDistanceFromTraceIntersection(manhattan bool, traceA string, traceB string) (int, error) {

	start := Coord{X: 0, Y: 0}

	path1, err := generatePathsFromString(traceA, start)
	if err != nil {
		return 0, err
	}

	path2, err := generatePathsFromString(traceB, start)
	if err != nil {
		return 0, err
	}

	return getComputeBestFromPaths(manhattan, path1, path2), nil
}

func intAbs(value int) int {
	return int(math.Abs(float64(value)))
}

func (coord *Coord) Distance(other Coord) int {
	return intAbs(other.X-coord.X) + intAbs(other.Y-coord.Y)
}

func getComputeBestFromPaths(manhattan bool, pathA []Coord, pathB []Coord) int {

	min := 0
	last := 0
	start := pathA[0]

	for stepA, coordA := range pathA {
		for stepB, coordB := range pathB {
			if (coordA.X == coordB.X) && (coordA.Y == coordB.Y) {

				if manhattan {
					last = getManhattanDistance(start, coordA)
				} else {
					last = stepA + stepB
				}

				if min == 0 || last < min {
					min = last
				}
			}
		}
	}

	return min
}

func getManhattanDistance(start Coord, coordA Coord) int {
	distance := start.Distance(coordA)
	log.Printf("intersection at %v, distance %v", coordA, distance)
	return distance
}

func generatePathsFromString(input string, start Coord) ([]Coord, error) {

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
