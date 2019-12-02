package day2

import (
	"log"
)

var opcodes map[int]OpCode

const (
	ADD  = 1
	MULT = 2
	HALT = 99
)

func init() {
	opcodes = make(map[int]OpCode, 0)
	opcodes[ADD] = add
	opcodes[MULT] = multiply
	opcodes[HALT] = halt
}

func RunProcessor(memory []int) Registers {

	registers := Registers{}

	for registers.PC < len(memory) {
		if memory[registers.PC] == HALT {
			break
		}
		instruction, ok := opcodes[memory[registers.PC]]

		if ok {
			instruction(memory, &registers)
		} else {
			log.Fatalf("unknown instruction %v, registers %v, memory %v", memory[registers.PC], registers, memory)
		}
	}
	return registers
}

type OpCode func(memory []int, registers *Registers)

type Registers struct {
	PC int
}

func add(memory []int, registers *Registers) {
	v1 := memory[memory[registers.PC+1]]
	v2 := memory[memory[registers.PC+2]]
	memory[memory[registers.PC+3]] = v1 + v2
	registers.PC += 4
}

func multiply(memory []int, registers *Registers) {
	v1 := memory[memory[registers.PC+1]]
	v2 := memory[memory[registers.PC+2]]
	memory[memory[registers.PC+3]] = v1 * v2
	registers.PC += 4
}

func halt(memory []int, registers *Registers) {

}
