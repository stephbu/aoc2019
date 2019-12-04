package day4

import (
	"strconv"

	"github.com/stephbu/aoc2019/util"
)

func getValidPasswords(limitRepetition bool, low int, high int) int {
	valid := 0

	for index := range util.N(high - low) {

		attempt := strconv.Itoa(index + low)

		if passwordScanner(limitRepetition, attempt) {
			valid++
		}

	}

	return valid
}

func passwordScanner(limitRepetition bool, passwordAttempt string) bool {

	decreasingPair := false
	digitPair := false
	digitRepetition := 0

	indexLast := 0

	for index := 0; index < len(passwordAttempt); index++ {

		if index > 0 {

			if passwordAttempt[index] < passwordAttempt[indexLast] {
				decreasingPair = true
			}

			if passwordAttempt[index] == passwordAttempt[indexLast] {

				digitRepetition++
				if !digitPair && !limitRepetition {
					digitPair = true
				}

			} else {

				// change in digits
				if digitRepetition == 1 {
					digitPair = true
				}
				digitRepetition = 0

			}

		}

		indexLast = index
	}

	if digitRepetition == 1 {
		digitPair = true
	}

	return digitPair && !decreasingPair

}
