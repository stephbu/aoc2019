package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stephbu/aoc2019/util"
)

func TestPart1Sample1(t *testing.T) {

	steps1 := "R8,U5,L5,D3"
	steps2 := "U7,R6,D4,L4"

	dist, err := getDistanceFromSteps(true, steps1, steps2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 6, dist)
}

func TestPart1Sample2(t *testing.T) {

	steps1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	steps2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	dist, err := getDistanceFromSteps(true, steps1, steps2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 159, dist)
}

func TestPart1Sample3(t *testing.T) {

	steps1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	steps2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	dist, err := getDistanceFromSteps(true, steps1, steps2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 135, dist)
}

func TestPart1PuzzleInput(t *testing.T) {

	traces, err := util.ReadLines("puzzleinput1.txt")
	dist, err := getDistanceFromSteps(true, traces[0], traces[1])
	if err != nil {
		t.Error(err)
	}
	t.Logf("distance %v", dist)
}

func TestPart2Sample1(t *testing.T) {

	steps1 := "R8,U5,L5,D3"
	steps2 := "U7,R6,D4,L4"

	dist, err := getDistanceFromSteps(false, steps1, steps2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 30, dist)
}

func TestPart2Sample2(t *testing.T) {

	steps1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	steps2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	dist, err := getDistanceFromSteps(false, steps1, steps2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 610, dist)
}

func TestPart2Sample3(t *testing.T) {

	steps1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	steps2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	dist, err := getDistanceFromSteps(false, steps1, steps2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 410, dist)
}

func TestPart2PuzzleInput(t *testing.T) {

	traces, err := util.ReadLines("puzzleinput2.txt")
	dist, err := getDistanceFromSteps(false, traces[0], traces[1])
	if err != nil {
		t.Error(err)
	}
	t.Logf("distance %v", dist)
}
