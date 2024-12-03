package day3

import (
	"os"
	"regexp"
	"strconv"
)

func Part2() int {
	dat, err := os.ReadFile("day3/input")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

	submatches := r.FindAllStringSubmatch(string(dat), -1)

	var total int
	doIt := true
	for _, match := range submatches {
		switch match[1] {
		case "do()":
			doIt = true
		case "don't()":
			doIt = false
		default:
			if doIt {
				n, _ := strconv.Atoi(match[2])
				m, _ := strconv.Atoi(match[3])
				total += n * m
			}
		}
	}

	return total
}
