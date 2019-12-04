package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample1(t *testing.T) {
	assert.Equal(t, true, passwordScanner(false, "111111"))
}

func TestPart1Sample2(t *testing.T) {
	assert.Equal(t, false, passwordScanner(false, "223450"))
}

func TestPart1Sample3(t *testing.T) {
	assert.Equal(t, false, passwordScanner(false, "123789"))
}

func TestPart1PuzzleInput(t *testing.T) {
	valid := getValidPasswords(false, 171309, 643603)
	t.Logf("Valid:%v", valid)
}

func TestPart2Sample1(t *testing.T) {
	assert.Equal(t, true, passwordScanner(true, "112233"))
}

func TestPart2Sample2(t *testing.T) {
	assert.Equal(t, false, passwordScanner(true, "123444"))
}

func TestPart2Sample3(t *testing.T) {
	assert.Equal(t, true, passwordScanner(true, "111122"))
}

func TestPart2Sample4(t *testing.T) {
	assert.Equal(t, false, passwordScanner(true, "555666"))
}

func TestPart2Sample5(t *testing.T) {
	assert.Equal(t, false, passwordScanner(true, "555664"))
}

func TestPart2Sample6(t *testing.T) {
	assert.Equal(t, false, passwordScanner(true, "556644"))
}

func TestPart2Sample8(t *testing.T) {
	assert.Equal(t, true, passwordScanner(true, "445555"))
}

func TestPart2Sample9(t *testing.T) {
	assert.Equal(t, false, passwordScanner(true, "111111"))
}

func TestPart2PuzzleInput(t *testing.T) {
	valid := getValidPasswords(true, 171309, 643603)
	t.Logf("Valid:%v", valid)
}
