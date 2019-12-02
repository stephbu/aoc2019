package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stephbu/aoc2019/util"
)

func TestPart1Sample1(t *testing.T) {
	memory := util.StringSplitInt("1,0,0,0,99")
	_ = RunProcessor(memory)
	assert.Equal(t, 2, memory[0])
}

func TestPart1Sample2(t *testing.T) {
	memory := util.StringSplitInt("2,3,0,3,99")
	_ = RunProcessor(memory)
	assert.Equal(t, 6, memory[3])
}

func TestPart1Sample3(t *testing.T) {
	memory := util.StringSplitInt("2,4,4,5,99,0")
	_ = RunProcessor(memory)
	assert.Equal(t, 9801, memory[5])
}

func TestPart1Sample4(t *testing.T) {
	memory := util.StringSplitInt("1,1,1,4,99,5,6,0,99")
	_ = RunProcessor(memory)
	assert.Equal(t, 30, memory[0])
}

func TestPart1PuzzleInput(t *testing.T) {
	memory := util.StringSplitInt("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,6,23,2,6,23,27,2,27,9,31,1,5,31,35,1,35,10,39,2,39,9,43,1,5,43,47,2,47,10,51,1,51,6,55,1,5,55,59,2,6,59,63,2,63,6,67,1,5,67,71,1,71,9,75,2,75,10,79,1,79,5,83,1,10,83,87,1,5,87,91,2,13,91,95,1,95,10,99,2,99,13,103,1,103,5,107,1,107,13,111,2,111,9,115,1,6,115,119,2,119,6,123,1,123,6,127,1,127,9,131,1,6,131,135,1,135,2,139,1,139,10,0,99,2,0,14,0")
	memory[1] = 12
	memory[2] = 2
	_ = RunProcessor(memory)
	t.Logf("memory[0]=%v", memory[0])
}
