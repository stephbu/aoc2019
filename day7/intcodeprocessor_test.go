package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stephbu/aoc2019/util"
)

func TestInputOutput(t *testing.T) {
	memory := util.StringSplitInt("3,0,4,0,99")

	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("123")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 123, memory[0])
}

func Test2(t *testing.T) {
	memory := util.StringSplitInt("3,0,4,0,99")

	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("123")
	err := proc.DebugProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 123, memory[0])
}

func Test3(t *testing.T) {
	memory := util.StringSplitInt("1002,4,3,4,33")

	proc := NewIntCodeProcessor(memory)
	err := proc.DebugProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 99, memory[4])
}

func Test4(t *testing.T) {
	memory := util.StringSplitInt("1101,100,-1,4,0")
	proc := NewIntCodeProcessor(memory)
	err := proc.DebugProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 99, memory[4])
}

func Test5(t *testing.T) {

	memory := util.StringSplitInt("3,225,1,225,6,6,1100,1,238,225,104,0,1102,68,5,225,1101,71,12,225,1,117,166,224,1001,224,-100,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1001,66,36,224,101,-87,224,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1101,26,51,225,1102,11,61,224,1001,224,-671,224,4,224,1002,223,8,223,1001,224,5,224,1,223,224,223,1101,59,77,224,101,-136,224,224,4,224,1002,223,8,223,1001,224,1,224,1,223,224,223,1101,11,36,225,1102,31,16,225,102,24,217,224,1001,224,-1656,224,4,224,102,8,223,223,1001,224,1,224,1,224,223,223,101,60,169,224,1001,224,-147,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1102,38,69,225,1101,87,42,225,2,17,14,224,101,-355,224,224,4,224,102,8,223,223,1001,224,2,224,1,224,223,223,1002,113,89,224,101,-979,224,224,4,224,1002,223,8,223,1001,224,7,224,1,224,223,223,1102,69,59,225,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,7,677,677,224,1002,223,2,223,1006,224,329,1001,223,1,223,1007,226,226,224,1002,223,2,223,1006,224,344,1001,223,1,223,1108,226,677,224,102,2,223,223,1005,224,359,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,374,101,1,223,223,1107,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,7,226,677,224,1002,223,2,223,1005,224,404,101,1,223,223,1008,677,226,224,102,2,223,223,1005,224,419,101,1,223,223,1008,226,226,224,102,2,223,223,1006,224,434,101,1,223,223,107,226,226,224,1002,223,2,223,1005,224,449,1001,223,1,223,108,226,677,224,102,2,223,223,1005,224,464,101,1,223,223,1108,677,226,224,102,2,223,223,1005,224,479,101,1,223,223,1007,226,677,224,102,2,223,223,1006,224,494,101,1,223,223,107,677,677,224,102,2,223,223,1005,224,509,101,1,223,223,108,677,677,224,102,2,223,223,1006,224,524,1001,223,1,223,8,226,677,224,102,2,223,223,1005,224,539,101,1,223,223,107,677,226,224,102,2,223,223,1005,224,554,1001,223,1,223,8,226,226,224,102,2,223,223,1006,224,569,1001,223,1,223,7,677,226,224,1002,223,2,223,1005,224,584,1001,223,1,223,1108,226,226,224,102,2,223,223,1005,224,599,1001,223,1,223,1107,677,677,224,1002,223,2,223,1006,224,614,1001,223,1,223,1007,677,677,224,1002,223,2,223,1006,224,629,1001,223,1,223,108,226,226,224,102,2,223,223,1005,224,644,1001,223,1,223,8,677,226,224,1002,223,2,223,1005,224,659,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("1")
	err := proc.DebugProcessor()
	assert.Nil(t, err)
}

func Test6(t *testing.T) {

	memory := util.StringSplitInt("3,9,8,9,10,9,4,9,99,-1,8")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("8")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 1, memory[9])
}

func Test7(t *testing.T) {

	memory := util.StringSplitInt("3,9,8,9,10,9,4,9,99,-1,8")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("16")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 0, memory[9])
}

func Test8(t *testing.T) {

	memory := util.StringSplitInt("3,9,7,9,10,9,4,9,99,-1,8")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("7")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 1, memory[9])
}

func Test9(t *testing.T) {

	memory := util.StringSplitInt("3,9,7,9,10,9,4,9,99,-1,8")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("8")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 0, memory[9])
}

func Test10(t *testing.T) {

	memory := util.StringSplitInt("3,3,1108,-1,8,3,4,3,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("8")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 1, memory[3])
}

func Test11(t *testing.T) {

	memory := util.StringSplitInt("3,3,1108,-1,8,3,4,3,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("16")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 0, memory[3])
}

func Test12(t *testing.T) {

	memory := util.StringSplitInt("3,3,1107,-1,8,3,4,3,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("7")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 1, memory[3])
}

func Test13(t *testing.T) {

	memory := util.StringSplitInt("3,3,1107,-1,8,3,4,3,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("8")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 0, memory[3])
}

func Test14(t *testing.T) {

	memory := util.StringSplitInt("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("0")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 0, memory[13])
}

func Test15(t *testing.T) {

	memory := util.StringSplitInt("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("100")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 1, memory[13])
}

func Test16(t *testing.T) {

	memory := util.StringSplitInt("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("0")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 0, memory[12])
}

func Test17(t *testing.T) {

	memory := util.StringSplitInt("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("10")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	assert.Equal(t, 1, memory[12])
}

func Test18(t *testing.T) {
	memory := util.StringSplitInt("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("7")
	err := proc.RunProcessor()
	assert.Nil(t, err)

	output := proc.GetOutput()
	assert.GreaterOrEqual(t, len(output), 1)
	assert.Equal(t, 999, output[len(output)-1])
}

func Test19(t *testing.T) {
	memory := util.StringSplitInt("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("8")
	err := proc.RunProcessor()
	assert.Nil(t, err)

	output := proc.GetOutput()
	assert.GreaterOrEqual(t, len(output), 1)
	assert.Equal(t, 1000, output[len(output)-1])
}

func Test20(t *testing.T) {
	memory := util.StringSplitInt("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	proc := NewIntCodeProcessor(memory)
	proc.StoreInput("9")
	err := proc.RunProcessor()
	assert.Nil(t, err)
	output := proc.GetOutput()
	assert.GreaterOrEqual(t, len(output), 1)
	assert.Equal(t, 1001, output[len(output)-1])
}
