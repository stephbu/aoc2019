package day5

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

var opCodes map[int]OpCode

const (
	ADD    = 1
	MULT   = 2
	INPUT  = 3
	OUTPUT = 4
	JIFT   = 5
	JIFF   = 6
	LT     = 7
	EQ     = 8
	HALT   = 99
)

func init() {

	inputChannel = make(chan string, 100)
	outputChannel = make(chan int, 100)

	opCodes = make(map[int]OpCode, 0)
	opCodes[ADD] = add
	opCodes[MULT] = multiply
	opCodes[HALT] = halt
	opCodes[INPUT] = input
	opCodes[OUTPUT] = output
	opCodes[JIFT] = jift
	opCodes[JIFF] = jiff
	opCodes[LT] = lt
	opCodes[EQ] = eq
}

var inputChannel chan string
var outputChannel chan int

func StoreInput(in string) {
	inputChannel <- in
}

func RunProcessor(memory []int) (Registers, error) {

	registers := Registers{}

	for registers.PC < len(memory) {
		if memory[registers.PC] == HALT {
			break
		}

		opCode := memory[registers.PC]

		parameterMask, instructionOpCode := divmod(opCode, 100)
		registers.ParameterMask = []int{}

		mask := 0
		last := parameterMask
		for last > 0 {
			last, mask = divmod(last, 10)
			registers.ParameterMask = append(registers.ParameterMask, mask)
		}

		instruction, ok := opCodes[instructionOpCode]

		if ok {
			instruction(memory, &registers)
		} else {
			err := errors.Errorf("unknown instruction %v, registers %v, memory %v", memory[registers.PC], registers, memory)
			return registers, err
		}
	}
	return registers, nil
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

type OpCode func(memory []int, registers *Registers)

type ParameterMask []int

type Registers struct {
	PC            int
	ParameterMask []int
}

func getValue(memory []int, registers *Registers, offset int) int {

	v1 := 0

	// offset is 1-based
	// parametermask is 0-based
	parameterMask := 0
	if offset <= len(registers.ParameterMask) {
		parameterMask = registers.ParameterMask[offset-1]
	}

	if parameterMask == 0 {
		v1 = memory[memory[registers.PC+offset]]
	} else {
		v1 = memory[registers.PC+offset]
	}
	return v1
}

func add(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	v2 := getValue(memory, registers, 2)
	memory[memory[registers.PC+3]] = v1 + v2
	registers.PC += 4
}

func multiply(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	v2 := getValue(memory, registers, 2)
	memory[memory[registers.PC+3]] = v1 * v2
	registers.PC += 4
}

func input(memory []int, registers *Registers) {
	v1, err := strconv.Atoi(<-inputChannel)

	if err != nil {
		log.Fatal(err)
	}

	memory[memory[registers.PC+1]] = v1
	registers.PC += 2
}

func output(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	fmt.Printf("output-> %v\r\n", v1)
	outputChannel <- v1
	registers.PC += 2
}

func jift(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	v2 := getValue(memory, registers, 2)

	if v1 != 0 {
		registers.PC = v2
	} else {
		registers.PC += 3
	}
}

func jiff(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	v2 := getValue(memory, registers, 2)

	if v1 == 0 {
		registers.PC = v2
	} else {
		registers.PC += 3
	}
}

func lt(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	v2 := getValue(memory, registers, 2)

	if v1 < v2 {
		memory[memory[registers.PC+3]] = 1
	} else {
		memory[memory[registers.PC+3]] = 0
	}
	registers.PC += 4
}

func eq(memory []int, registers *Registers) {
	v1 := getValue(memory, registers, 1)
	v2 := getValue(memory, registers, 2)

	if v1 == v2 {
		memory[memory[registers.PC+3]] = 1
	} else {
		memory[memory[registers.PC+3]] = 0
	}
	registers.PC += 4
}

func halt(_ []int, _ *Registers) {

}
