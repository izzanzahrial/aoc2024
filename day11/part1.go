package day11

import (
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	dat, err := os.ReadFile("day11/input")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(dat), " ")

	for i := 0; i < 25; i++ {
		var nextValues []string
		for _, val := range values {
			intVal, _ := strconv.Atoi(val)
			if intVal == 0 {
				intVal = 1
				nextValues = append(nextValues, strconv.Itoa(intVal))
			} else if len(val)%2 == 0 {
				currVal1 := val[:len(val)/2]
				currVal2 := val[len(val)/2:]
				intVal2, _ := strconv.Atoi(currVal2)
				strVal2 := strconv.Itoa(intVal2)
				nextValues = append(nextValues, currVal1)
				nextValues = append(nextValues, strVal2)
			} else {
				intVal = intVal * 2024
				nextValues = append(nextValues, strconv.Itoa(intVal))
			}
		}

		values = nextValues
	}

	return len(values)
}
