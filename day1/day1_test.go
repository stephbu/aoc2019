package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	Main()
}

func TestPart1Sample1(t *testing.T) {
	assert.Equal(t, getFuelForMass(12), 2)
}

func TestPart1Sample2(t *testing.T) {
	assert.Equal(t, getFuelForMass(14), 2)
}

func TestPart1Sample3(t *testing.T) {
	assert.Equal(t, getFuelForMass(1969), 654)
}

func TestPart1Sample4(t *testing.T) {
	assert.Equal(t, getFuelForMass(100756), 33583)
}

func TestPart2Sample1(t *testing.T) {
	assert.Equal(t, getFullyBurdenedFuel(14), 2)
}

func TestPart2Sample2(t *testing.T) {
	assert.Equal(t, getFullyBurdenedFuel(1969), 966)
}

func TestPart2Sample3(t *testing.T) {
	assert.Equal(t, getFullyBurdenedFuel(100756), 50346)
}
