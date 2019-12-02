package day1

import (
	"log"
	"strconv"

	"github.com/stephbu/aoc2019/util"
)

func Main() {
	calculateMass(getFuelForMass)
	calculateMass(getFullyBurdenedFuel)
}

type MassCalc func(int) int

func calculateMass(getMass MassCalc) {
	lines, err := util.ReadLines("day1/puzzleinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	totalFuel := 0
	moduleCount := 0
	for _, line := range lines {
		moduleMass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		fuel := getMass(moduleMass)
		log.Printf("getModuleFuel(%v)=%v", moduleMass, fuel)

		moduleCount++
		totalFuel += fuel
	}
	log.Printf("%v Modules, %v", moduleCount, totalFuel)
}

func getFullyBurdenedFuel(mass int) int {

	totalFuel := getFuelForMass(mass)

	remainingMass := totalFuel

	for remainingMass > 0 {
		remainingMass = getFuelForMass(remainingMass)
		if remainingMass > 0 {
			totalFuel += remainingMass
		}
	}

	return totalFuel
}

func getFuelForMass(mass int) int {

	quotient := mass / 3
	fuel := quotient - 2
	return fuel

}
