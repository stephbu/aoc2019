package day7

import (
	"log"
	"runtime"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stephbu/aoc2019/util"
)

func TestPart1Sample1(t *testing.T) {
	program := util.StringSplitInt("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
	out := RunProcessorChain(program, []int{4, 3, 2, 1, 0})
	t.Logf("out:%v", out)
	assert.Equal(t, 43210, out)
}

func TestPart1Sample2(t *testing.T) {
	program := util.StringSplitInt("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")
	out := RunProcessorChain(program, []int{0, 1, 2, 3, 4})
	t.Logf("out:%v", out)
	assert.Equal(t, 54321, out)
}

func TestPart1Sample3(t *testing.T) {
	program := util.StringSplitInt("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")
	out := RunProcessorChain(program, []int{1, 0, 4, 3, 2})
	t.Logf("out:%v", out)
	assert.Equal(t, 65210, out)
}

func TestPart1Puzzle(t *testing.T) {
	program := util.StringSplitInt("3,8,1001,8,10,8,105,1,0,0,21,38,63,76,89,106,187,268,349,430,99999,3,9,1001,9,5,9,102,3,9,9,1001,9,2,9,4,9,99,3,9,101,4,9,9,102,3,9,9,101,4,9,9,1002,9,3,9,101,2,9,9,4,9,99,3,9,101,5,9,9,1002,9,4,9,4,9,99,3,9,101,2,9,9,1002,9,5,9,4,9,99,3,9,1001,9,5,9,1002,9,5,9,1001,9,5,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99")
	permutations := permutations([]int{0, 1, 2, 3, 4})

	max := 0

	for _, permutation := range permutations {
		out := RunProcessorChain(program, permutation)
		if out > max {
			t.Logf("out:%v", out)
			max = out
		}
	}
	t.Logf("max:%v", max)
}

func TestPart2Puzzle(t *testing.T) {
	program := util.StringSplitInt("3,8,1001,8,10,8,105,1,0,0,21,38,63,76,89,106,187,268,349,430,99999,3,9,1001,9,5,9,102,3,9,9,1001,9,2,9,4,9,99,3,9,101,4,9,9,102,3,9,9,101,4,9,9,1002,9,3,9,101,2,9,9,4,9,99,3,9,101,5,9,9,1002,9,4,9,4,9,99,3,9,101,2,9,9,1002,9,5,9,4,9,99,3,9,1001,9,5,9,1002,9,5,9,1001,9,5,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99")
	permutations := permutations([]int{9, 8, 7, 6, 5})

	max := 0

	for _, permutation := range permutations {
		out := Evaluate(program, permutation)
		if out > max {
			t.Logf("out:%v", out)
			max = out
		}
	}
	t.Logf("max:%v", max)

}

func Evaluate(program []int, phaseSettings []int) int {
	procs := make([]*IntCodeProcessor, 5)
	procs[0] = NewIntCodeProcessor(program).SetLabel("A").Send(phaseSettings[0])
	procs[1] = NewIntCodeProcessor(program).SetLabel("B").Send(phaseSettings[1])
	procs[2] = NewIntCodeProcessor(program).SetLabel("C").Send(phaseSettings[2])
	procs[3] = NewIntCodeProcessor(program).SetLabel("D").Send(phaseSettings[3])
	procs[4] = NewIntCodeProcessor(program).SetLabel("E").Send(phaseSettings[4])

	go func() {
		err := procs[0].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[1].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[2].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[3].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[4].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	procs[0].Send(0)

	currentProc := 0
	nextProc := 0
	lastOut := 0

	for {

		if (currentProc == 0) && procs[currentProc].Halted {
			return lastOut
		}

		lastOut = <-procs[currentProc].OutputChannel

		nextProc = currentProc + 1
		if nextProc == len(procs) {
			nextProc = 0
		}

		procs[nextProc].Send(lastOut)
		currentProc = nextProc
	}

}

func TestPart2Sample1(t *testing.T) {

	runtime.GOMAXPROCS(100)

	program := util.StringSplitInt("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5")

	procs := make([]*IntCodeProcessor, 5)
	procs[0] = NewIntCodeProcessor(program).SetLabel("A").Send(9)
	procs[1] = NewIntCodeProcessor(program).SetLabel("B").Send(8)
	procs[2] = NewIntCodeProcessor(program).SetLabel("C").Send(7)
	procs[3] = NewIntCodeProcessor(program).SetLabel("D").Send(6)
	procs[4] = NewIntCodeProcessor(program).SetLabel("E").Send(5)

	go func() {
		err := procs[0].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[1].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[2].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[3].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[4].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	procs[0].Send(0)

	currentProc := 0
	nextProc := 0
	lastOut := 0

	for {

		if (currentProc == 0) && procs[currentProc].Halted {
			t.Logf("last output :%v", lastOut)
			assert.Equal(t, 139629729, lastOut)
			return
		}

		lastOut = <-procs[currentProc].OutputChannel

		nextProc = currentProc + 1
		if nextProc == len(procs) {
			nextProc = 0
		}

		procs[nextProc].Send(lastOut)
		currentProc = nextProc
	}

}

func TestPart2Sample2(t *testing.T) {

	runtime.GOMAXPROCS(100)

	program := util.StringSplitInt("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")

	procs := make([]*IntCodeProcessor, 5)
	procs[0] = NewIntCodeProcessor(program).SetLabel("A").Send(9)
	procs[1] = NewIntCodeProcessor(program).SetLabel("B").Send(7)
	procs[2] = NewIntCodeProcessor(program).SetLabel("C").Send(8)
	procs[3] = NewIntCodeProcessor(program).SetLabel("D").Send(5)
	procs[4] = NewIntCodeProcessor(program).SetLabel("E").Send(6)

	go func() {
		err := procs[0].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[1].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[2].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[3].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := procs[4].DebugProcessor()
		if err != nil {
			log.Fatal(err)
		}
	}()

	procs[0].Send(0)

	currentProc := 0
	nextProc := 0
	lastOut := 0

	for {

		if (currentProc == 0) && procs[currentProc].Halted {
			t.Logf("last output :%v", lastOut)
			assert.Equal(t, 18216, lastOut)
			return
		}

		lastOut = <-procs[currentProc].OutputChannel

		nextProc = currentProc + 1
		if nextProc == len(procs) {
			nextProc = 0
		}

		procs[nextProc].Send(lastOut)
		currentProc = nextProc
	}

}

func TestPermute(t *testing.T) {
	permutations := permutations([]int{1, 2, 3})
	t.Logf("permutations: %v", permutations)
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func RunProcessorChain(program []int, phaseSetting []int) int {
	out := 0
	out = RunProc(program, phaseSetting[0], out)
	out = RunProc(program, phaseSetting[1], out)
	out = RunProc(program, phaseSetting[2], out)
	out = RunProc(program, phaseSetting[3], out)
	out = RunProc(program, phaseSetting[4], out)
	return out
}

func RunProc(program []int, ampInput int, previousOut int) int {
	a := NewIntCodeProcessor(program)
	go func() {
		err := a.RunProcessor()
		if err != nil {
			log.Fatalf("error processor a: %v", err)
		}
	}()

	a.StoreInput(strconv.Itoa(ampInput))
	a.StoreInput(strconv.Itoa(previousOut))

	out := <-a.OutputChannel
	return out
}
