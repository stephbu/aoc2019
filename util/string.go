package util

import (
	"log"
	"strconv"
	"strings"
)

func StringSplitInt(str string) []int {

	intArr := make([]int, 0)

	for _, strValue := range strings.Split(str, ",") {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			log.Fatal(err)
		}
		intArr = append(intArr, intValue)
	}

	return intArr
}
