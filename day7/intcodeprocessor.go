package day7

import (
	"log"
	"strconv"

	"github.com/pkg/errors"
)

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

type IntCodeProcessor struct {
	InputChannel  chan int
	OutputChannel chan int
	memory        []int
	Registers     Registers
	Halted        bool
	opCodes       map[int]OpCode
	Label         string
}

func NewIntCodeProcessor(memory []int) *IntCodeProcessor {
	result := IntCodeProcessor{}
	result.memory = make([]int, len(memory))
	copy(result.memory, memory)
	result.InputChannel = make(chan int, 10000)
	result.OutputChannel = make(chan int, 10000)
	result.Registers = Registers{}
	result.opCodes = result.LoadOpCodes()
	result.Halted = false
	return &result
}

func (proc *IntCodeProcessor) LoadOpCodes() map[int]OpCode {
	opCodes := make(map[int]OpCode, 0)
	opCodes[ADD] = proc.add
	opCodes[MULT] = proc.multiply
	opCodes[HALT] = proc.halt
	opCodes[INPUT] = proc.input
	opCodes[OUTPUT] = proc.output
	opCodes[JIFT] = proc.jift
	opCodes[JIFF] = proc.jiff
	opCodes[LT] = proc.lt
	opCodes[EQ] = proc.eq

	return opCodes
}

func (proc *IntCodeProcessor) StoreInput(in string) {

	i, _ := strconv.Atoi(in)
	proc.InputChannel <- i
}

func (proc *IntCodeProcessor) DebugProcessor() error {

	proc.Registers.Debug = true
	return proc.execute()
}

func (proc *IntCodeProcessor) RunProcessor() error {

	return proc.execute()
}

func (proc *IntCodeProcessor) GetOutput() []int {

	var result []int
	for len(proc.OutputChannel) > 0 {
		result = append(result, <-proc.OutputChannel)
	}

	return result

}

func (proc *IntCodeProcessor) execute() error {

	for proc.Registers.PC < len(proc.memory) {
		if proc.memory[proc.Registers.PC] == HALT {
			proc.Halted = true
			break
		}

		opCode := proc.memory[proc.Registers.PC]

		parameterMask, instructionOpCode := divmod(opCode, 100)
		proc.Registers.ParameterMask = []int{}

		mask := 0
		last := parameterMask
		for last > 0 {
			last, mask = divmod(last, 10)
			proc.Registers.ParameterMask = append(proc.Registers.ParameterMask, mask)
		}

		instruction, ok := proc.opCodes[instructionOpCode]

		if ok {
			instruction()
		} else {
			err := errors.Errorf("unknown instruction %v, registers %v, memory %v", proc.memory[proc.Registers.PC], proc.Registers, proc.memory)
			return err
		}
	}
	return nil
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

type OpCode func()

type ParameterMask []int

type Registers struct {
	PC            int
	Debug         bool
	ParameterMask []int
}

func getValue(memory []int, registers Registers, offset int) int {

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

func (proc *IntCodeProcessor) debug(format string, values ...interface{}) {
	if proc.Registers.Debug {
		format = "Proc(%v):" + format + "\r\n"

		values = append([]interface{}{proc.Label}, values...)
		log.Printf(format, values...)
	}
}

func (proc *IntCodeProcessor) add() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	v2 := getValue(proc.memory, proc.Registers, 2)
	proc.memory[proc.memory[proc.Registers.PC+3]] = v1 + v2
	proc.debug("ADD(%v,%v) -> [%v]=%v", v1, v2, proc.memory[proc.Registers.PC+3], v1+v2)

	proc.Registers.PC += 4
}

func (proc *IntCodeProcessor) multiply() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	v2 := getValue(proc.memory, proc.Registers, 2)
	proc.memory[proc.memory[proc.Registers.PC+3]] = v1 * v2
	proc.debug("MULT(%v,%v) -> [%v]=%v", v1, v2, proc.memory[proc.Registers.PC+3], v1*v2)
	proc.Registers.PC += 4
}

func (proc *IntCodeProcessor) input() {
	v1 := <-proc.InputChannel
	proc.memory[proc.memory[proc.Registers.PC+1]] = v1
	proc.debug("INPUT(%v) -> [%v]=%v", proc.memory[proc.Registers.PC+1], proc.memory[proc.Registers.PC+1], v1)
	proc.Registers.PC += 2
}

func (proc *IntCodeProcessor) output() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	proc.debug("OUTPUT(%v) -> [%v]=%v", proc.memory[proc.Registers.PC+1], proc.memory[proc.Registers.PC+1], v1)
	proc.OutputChannel <- v1
	proc.Registers.PC += 2
}

func (proc *IntCodeProcessor) jift() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	v2 := getValue(proc.memory, proc.Registers, 2)

	if v1 != 0 {
		proc.Registers.PC = v2
		proc.debug("JIFT(%v,%v) -> PC=%v", v1, v2, v1)
	} else {
		proc.Registers.PC += 3
		proc.debug("JIFT(%v,%v) -> NOP", v1, v2)
	}
}

func (proc *IntCodeProcessor) jiff() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	v2 := getValue(proc.memory, proc.Registers, 2)

	if v1 == 0 {
		proc.Registers.PC = v2
		proc.debug("JIFF(%v,%v) -> PC=%v", v1, v2, v1)
	} else {
		proc.Registers.PC += 3
		proc.debug("JIFT(%v,%v) -> NOP", v1, v2)
	}
}

func (proc *IntCodeProcessor) lt() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	v2 := getValue(proc.memory, proc.Registers, 2)

	if v1 < v2 {
		proc.memory[proc.memory[proc.Registers.PC+3]] = 1
	} else {
		proc.memory[proc.memory[proc.Registers.PC+3]] = 0
	}
	proc.Registers.PC += 4
}

func (proc *IntCodeProcessor) eq() {
	v1 := getValue(proc.memory, proc.Registers, 1)
	v2 := getValue(proc.memory, proc.Registers, 2)

	if v1 == v2 {
		proc.memory[proc.memory[proc.Registers.PC+3]] = 1
	} else {
		proc.memory[proc.memory[proc.Registers.PC+3]] = 0
	}
	proc.Registers.PC += 4
}

func (proc *IntCodeProcessor) halt() {
	proc.Halted = true
}

func (proc *IntCodeProcessor) SetLabel(label string) *IntCodeProcessor {
	proc.Label = label
	return proc
}

func (proc *IntCodeProcessor) Send(value int) *IntCodeProcessor {
	proc.InputChannel <- value
	return proc
}
