package day6

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stephbu/aoc2019/util"
)

func TestPart1Sample1a(t *testing.T) {

	d, i := expandOrbits([]string{"AAA)BBB"})
	assert.Equal(t, 1, len(d))
	assert.Equal(t, 0, len(i))

}

func TestPart1Sample1b(t *testing.T) {

	d, i := expandOrbits([]string{"AAA)BBB", "BBB)CCC"})
	assert.Equal(t, 2, len(d))
	assert.Equal(t, 1, len(i))
}

func TestPart1Sample1c(t *testing.T) {

	d, i := expandOrbits([]string{"COM)B", "B)C", "C)D"})
	assert.Equal(t, 3, len(d))
	assert.Equal(t, 3, len(i))
}

func TestPart1Sample2a(t *testing.T) {

	lines, err := util.ReadLines("inputpart1sample2.txt")
	if err != nil {
		log.Fatalf("error reading file :%v", err)
	}
	d, i := expandOrbits(lines)
	assert.Equal(t, 1, checkChain("D", d))
	assert.Equal(t, 2, checkChain("D", i))
	assert.Equal(t, 1, checkChain("L", d))
	assert.Equal(t, 6, checkChain("L", i))
	assert.Equal(t, 0, checkChain("COM", d))
	assert.Equal(t, 0, checkChain("COM", i))
	assert.Equal(t, 42, len(d)+len(i))
}

func TestPart1Puzzle(t *testing.T) {
	lines, err := util.ReadLines("inputpuzzle.txt")
	if err != nil {
		log.Fatalf("error reading file :%v", err)
	}
	d, i := expandOrbits(lines)
	checkBodies(d, i)
	log.Printf("Direct:%v, Indirect:%v, Total:%v", len(d), len(i), len(d)+len(i))
	assert.Equal(t, 1438, len(d))

}

func TestPart2Sample2a(t *testing.T) {
	lines, err := util.ReadLines("inputpart2sample1.txt")
	if err != nil {
		log.Fatalf("error reading file :%v", err)
	}
	d, i := expandOrbits(lines)
	orbits := append(d, i...)

	shortestPath, commonCOM := findShortestPath(orbits, "YOU", "SAN")
	assert.Equal(t, 4, shortestPath)
	assert.Equal(t, "D", commonCOM)

}

func TestPart2Puzzle(t *testing.T) {
	lines, err := util.ReadLines("inputpuzzle.txt")
	if err != nil {
		log.Fatalf("error reading file :%v", err)
	}
	d, i := expandOrbits(lines)
	orbits := append(d, i...)
	shortestPath, commonCOM := findShortestPath(orbits, "YOU", "SAN")
	log.Printf("Shortest Hops:%v, Common Body:%v", shortestPath, commonCOM)
}

func checkBodies(direct []*Orbit, indirect []*Orbit) {

	orbitIndex := make(map[string]int, 0)

	for _, orbit := range append(direct, indirect...) {
		_, ok := orbitIndex[orbit.Body]
		if ok {
			orbitIndex[orbit.Body]++
		} else {
			orbitIndex[orbit.Body] = 1
		}
	}

	for key, value := range orbitIndex {
		if value < 2 {
			log.Printf("key:%v=%v", key, value)
		}
	}

}

func checkChain(bodyName string, orbits []*Orbit) (count int) {

	for _, orbit := range orbits {
		if orbit.Body == bodyName {
			count++
		}
	}

	return
}
