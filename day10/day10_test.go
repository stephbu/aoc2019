package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadMapFromFile(t *testing.T) {
	asteroidMap := LoadMapFromFile("sample1.txt")
	assert.Equal(t, 5, len(asteroidMap.Map))
	assert.Equal(t, 5, len(asteroidMap.Map[0]))

	assert.Equal(t, "#", asteroidMap.Map[3][4].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[0][0].MapChar)

	losMap := asteroidMap.CalculateLOS(Coord{X: 3, Y: 4}, AsteroidHash)

	t.Logf("%v", losMap)

}

func TestPart1Sample2(t *testing.T) {

	asteroidMap := LoadMapFromFile("sample2.txt")
	assert.Equal(t, 10, len(asteroidMap.Map))
	assert.Equal(t, 10, len(asteroidMap.Map[0]))

	assert.Equal(t, ".", asteroidMap.Map[0][0].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[9][0].MapChar)
	assert.Equal(t, "#", asteroidMap.Map[9][9].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[0][9].MapChar)

	maxCount := 0
	maxCoord := Coord{}

	asteroidMap.Enumerate(AsteroidHash, func(body Body) {
		losMap := asteroidMap.CalculateLOS(body.Coord, AsteroidHash)
		if len(losMap) > maxCount {
			maxCount = len(losMap)
			maxCoord = body.Coord
		}

	})

	assert.Equal(t, 33, maxCount)
	assert.Equal(t, Coord{5, 8}, maxCoord)
}

func TestPart1Sample3(t *testing.T) {

	asteroidMap := LoadMapFromFile("sample3.txt")
	assert.Equal(t, 10, len(asteroidMap.Map))
	assert.Equal(t, 10, len(asteroidMap.Map[0]))

	assert.Equal(t, "#", asteroidMap.Map[0][0].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[9][0].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[9][9].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[0][9].MapChar)

	maxCount := 0
	maxCoord := Coord{}

	asteroidMap.Enumerate(AsteroidHash, func(body Body) {
		losMap := asteroidMap.CalculateLOS(body.Coord, AsteroidHash)
		if len(losMap) > maxCount {
			maxCount = len(losMap)
			maxCoord = body.Coord
		}

	})

	assert.Equal(t, 35, maxCount)
	assert.Equal(t, Coord{1, 2}, maxCoord)
}

func TestPart1Sample4(t *testing.T) {

	asteroidMap := LoadMapFromFile("sample4.txt")
	assert.Equal(t, 10, len(asteroidMap.Map))
	assert.Equal(t, 10, len(asteroidMap.Map[0]))

	assert.Equal(t, ".", asteroidMap.Map[0][0].MapChar)
	assert.Equal(t, "#", asteroidMap.Map[9][0].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[9][9].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[0][9].MapChar)

	maxCount := 0
	maxCoord := Coord{}

	asteroidMap.Enumerate(AsteroidHash, func(body Body) {
		losMap := asteroidMap.CalculateLOS(body.Coord, AsteroidHash)
		if len(losMap) > maxCount {
			maxCount = len(losMap)
			maxCoord = body.Coord
		}

	})

	assert.Equal(t, 41, maxCount)
	assert.Equal(t, Coord{6, 3}, maxCoord)
}

func TestPart1Sample5(t *testing.T) {

	asteroidMap := LoadMapFromFile("sample5.txt")
	assert.Equal(t, 20, len(asteroidMap.Map))
	assert.Equal(t, 20, len(asteroidMap.Map[0]))

	assert.Equal(t, ".", asteroidMap.Map[0][0].MapChar)
	assert.Equal(t, "#", asteroidMap.Map[19][0].MapChar)
	assert.Equal(t, "#", asteroidMap.Map[19][19].MapChar)
	assert.Equal(t, "#", asteroidMap.Map[0][19].MapChar)

	maxCount := 0
	maxCoord := Coord{}

	asteroidMap.Enumerate(AsteroidHash, func(body Body) {
		losMap := asteroidMap.CalculateLOS(body.Coord, AsteroidHash)
		if len(losMap) > maxCount {
			maxCount = len(losMap)
			maxCoord = body.Coord
		}

	})

	assert.Equal(t, 210, maxCount)
	assert.Equal(t, Coord{11, 13}, maxCoord)
}

func TestPart1Puzzle(t *testing.T) {

	asteroidMap := LoadMapFromFile("inputpuzzle.txt")
	assert.Equal(t, 36, len(asteroidMap.Map))
	assert.Equal(t, 36, len(asteroidMap.Map[0]))

	assert.Equal(t, ".", asteroidMap.Map[0][0].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[35][0].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[35][35].MapChar)
	assert.Equal(t, ".", asteroidMap.Map[0][35].MapChar)

	maxCount := 0
	maxCoord := Coord{}

	asteroidMap.Enumerate(AsteroidHash, func(body Body) {
		losMap := asteroidMap.CalculateLOS(body.Coord, AsteroidHash)
		if len(losMap) > maxCount {
			maxCount = len(losMap)
			maxCoord = body.Coord
		}

	})

	t.Logf("Max:%v coords:%v", maxCount, maxCoord)

}

func TestPart2Sample6(t *testing.T) {

	asteroidMap := LoadMapFromFile("sample6.txt")
	losMap := asteroidMap.CalculateLOS(Coord{8, 3}, AsteroidHash)
	output := UnrollLOS(losMap)
	t.Logf("%v", output)

}

func TestPart2Sample7(t *testing.T) {

	asteroidMap := LoadMapFromFile("sample7.txt")
	losMap := asteroidMap.CalculateLOS(Coord{11, 13}, AsteroidHash)
	output := UnrollLOS(losMap)

	assert.Equal(t, Coord{11, 12}, output[0].Coord)
	assert.Equal(t, Coord{12, 1}, output[1].Coord)
	assert.Equal(t, Coord{12, 2}, output[2].Coord)
	assert.Equal(t, Coord{12, 8}, output[9].Coord)
	assert.Equal(t, Coord{16, 0}, output[19].Coord)
	assert.Equal(t, Coord{16, 9}, output[49].Coord)
	assert.Equal(t, Coord{10, 16}, output[99].Coord)
	assert.Equal(t, Coord{9, 6}, output[198].Coord)
	assert.Equal(t, Coord{8, 2}, output[199].Coord)
	assert.Equal(t, Coord{10, 9}, output[200].Coord)
	assert.Equal(t, Coord{11, 1}, output[298].Coord)

	index := 1
	for _, body := range output {
		t.Logf("Index:%d Coord:%v", index, body.Coord)
		index++
	}
}

func TestPart2Puzzle(t *testing.T) {

	asteroidMap := LoadMapFromFile("inputpuzzle.txt")
	losMap := asteroidMap.CalculateLOS(Coord{25, 31}, AsteroidHash)
	output := UnrollLOS(losMap)

	index := 1
	for _, body := range output {
		t.Logf("Index:%d Coord:%v Compute:%d", index, body.Coord, body.Coord.X*100+body.Coord.Y)
		index++
	}
}

func AsteroidHash(body Body) bool {
	if body.MapChar == "#" {
		return true
	}
	return false
}
