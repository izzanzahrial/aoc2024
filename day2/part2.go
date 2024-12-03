package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2() int {
	dat, err := os.ReadFile("day2/input")
	if err != nil {
		panic(err)
	}

	var result int
	values := strings.Split(string(dat), "\n")
	for _, val := range values {
		curr := strings.Split(val, " ")
		if isSafe(curr) {
			result++
			continue
		}

		for i := 0; i < len(curr); i++ {
			dampened := append([]string(nil), curr[:i]...)
			dampened = append(dampened, curr[i+1:]...)
			fmt.Println("old nums ", curr, " new nums ", dampened)
			if isSafe(dampened) {
				result++
				break
			}
		}
	}

	return result
}

func isSafe(nums []string) bool {
	increaseCount, decreaseCount := 0, 0

	for i := 1; i < len(nums); i++ {
		curr, _ := strconv.Atoi(nums[i])
		prev, _ := strconv.Atoi(nums[i-1])

		if math.Abs(float64(curr-prev)) > 3 {
			return false
		}

		if curr > prev {
			increaseCount++
		} else if curr < prev {
			decreaseCount++
		}
	}

	return increaseCount == len(nums)-1 || decreaseCount == len(nums)-1
}
