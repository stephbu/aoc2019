package day10

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/stephbu/aoc2019/util"
)

type StarMap struct {
	Map   [][]Body
	Index map[Coord]Body
}

func (starMap *StarMap) Enumerate(match func(body Body) bool, enumerationFunc func(body Body)) {
	for _, body := range starMap.Index {
		if match(body) {
			enumerationFunc(body)
		}
	}
}

type Coord struct {
	X int
	Y int
}

func (coord *Coord) Subtract(other Coord) Coord {
	return Coord{X: other.X - coord.X, Y: other.Y - coord.Y}
}

func (coord *Coord) String() string {
	return fmt.Sprintf("(%d,%d)", coord.X, coord.Y)
}

type Body struct {
	MapChar string
	Coord   Coord
}

func (body *Body) String() string {
	return fmt.Sprintf("%s(%v)", body.MapChar, body.Coord)
}

func LoadMapFromFile(filename string) StarMap {

	result := StarMap{}
	result.Index = make(map[Coord]Body, 0)
	lines, err := util.ReadLines(filename)

	if err != nil {
		log.Fatal(err)
	}

	x := len(lines[0])
	y := len(lines)

	result.Map = make([][]Body, x)
	for i := range result.Map {
		result.Map[i] = make([]Body, y)
	}

	for indexY, line := range lines {
		for indexX, charX := range line {
			result.Map[indexX][indexY].MapChar = string(charX)
			result.Map[indexX][indexY].Coord = Coord{X: indexX, Y: indexY}
			result.Index[result.Map[indexX][indexY].Coord] = result.Map[indexX][indexY]
		}
	}

	return result
}

type LOSResult struct {
	Body     Body
	Distance float64
}

func (los *LOSResult) String() string {
	return fmt.Sprintf("%v %f", los.Body, los.Distance)
}

// mapchar blank matches all objects
func (asteroids *StarMap) CalculateLOS(self Coord, match func(body Body) bool) map[float64][]LOSResult {

	results := make(map[float64][]LOSResult, 0)

	for coords, body := range asteroids.Index {

		if coords == self {
			continue
		}

		if !match(body) {
			continue
		}

		delta := coords.Subtract(self)

		rad := math.Atan2(float64(coords.X-self.X), float64(self.Y-coords.Y)) // In radians

		if rad < 0 {
			rad = rad + (math.Pi * 2)
		}
		deg := rad * (180 / math.Pi)

		distance := math.Sqrt(math.Pow(float64(delta.X), 2) + math.Pow(float64(delta.Y), 2))

		results[deg] = append(results[deg], LOSResult{Body: body, Distance: distance})

	}

	for key, _ := range results {
		sort.Slice(results[key], func(i, j int) bool {
			return results[key][i].Distance < results[key][j].Distance
		})
	}

	return results
}

func UnrollLOS(los map[float64][]LOSResult) []Body {

	type Path struct {
		Angle float64
		LOS   []LOSResult
	}

	paths := make([]Path, 0)

	for key, bodies := range los {
		paths = append(paths, Path{Angle: key, LOS: bodies})
	}

	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Angle < paths[j].Angle
	})

	var output []Body

	repeat := true
	for repeat {
		repeat = false
		for index, _ := range paths {
			if len(paths[index].LOS) == 0 {
				continue
			}
			output = append(output, paths[index].LOS[0].Body)
			paths[index].LOS = paths[index].LOS[1:]

			if len(paths[index].LOS) > 0 {
				repeat = true
			}
		}
	}

	return output
}
