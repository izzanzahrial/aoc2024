package day2

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	dat, err := os.ReadFile("day2/input")
	if err != nil {
		panic(err)
	}

	var result int
	values := strings.Split(string(dat), "\n")
	for _, val := range values {
		var safe bool
		curr := strings.Split(val, " ")
		num1, _ := strconv.Atoi(curr[0])
		num2, _ := strconv.Atoi(curr[1])
		if num1 > num2 {
			safe = isSafe1(curr, "desc")
		} else {
			safe = isSafe1(curr, "asc")
		}

		if safe {
			result++
		}
	}

	return result
}

func isSafe1(nums []string, orderKey string) bool {
	if orderKey == "desc" {
		for i := 0; i < len(nums)-1; i++ {
			num1, _ := strconv.Atoi(nums[i])
			num2, _ := strconv.Atoi(nums[i+1])
			if num1 < num2 || num1 == num2 || math.Abs(float64(num1-num2)) > 3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(nums)-1; i++ {
			num1, _ := strconv.Atoi(nums[i])
			num2, _ := strconv.Atoi(nums[i+1])
			if num1 > num2 || num1 == num2 || math.Abs(float64(num1-num2)) > 3 {
				return false
			}
		}
	}

	return true
}
