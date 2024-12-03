package day3

import (
	"os"
	"regexp"
	"strconv"
)

func Part1() int {
	dat, err := os.ReadFile("day3/input")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`mul\(([\d]{1,3}),([\d]{1,3})\)`)

	submatches := r.FindAllStringSubmatch(string(dat), -1)

	var total int
	for i := 0; i < len(submatches); i++ {
		n, _ := strconv.Atoi(submatches[i][1])
		m, _ := strconv.Atoi(submatches[i][2])

		total += n * m
	}

	return total
}
