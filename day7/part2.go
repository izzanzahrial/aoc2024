package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() int {
	dat, err := os.ReadFile("day7/input")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(dat), "\n")
	calibrationsMap := make(map[int][]int, len(values))
	for _, val := range values {
		vals := strings.Split(string(val), ":")
		result, _ := strconv.Atoi(vals[0])
		currs := strings.Split(vals[1], " ")
		for _, curr := range currs {
			num, _ := strconv.Atoi(curr)
			calibrationsMap[result] = append(calibrationsMap[result], num)
		}
	}

	var total int
	for target, nums := range calibrationsMap {
		if possibleConcat(target, nums) {
			total += target
		}
	}

	return total
}

func possibleConcat(target int, nums []int) bool {
	if nums[0] == target {
		return true
	}

	if nums[0] > target || len(nums) == 1 {
		return false
	}

	numsAdd := make([]int, len(nums)-1)
	numsMulti := make([]int, len(nums)-1)
	numsConcat := make([]int, len(nums)-1)
	numsAdd[0] = nums[0] + nums[1]
	numsMulti[0] = nums[0] * nums[1]
	currNum, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[1]))
	numsConcat[0] = currNum

	if len(nums) == 2 {
		return numsAdd[0] == target || numsMulti[0] == target || numsConcat[0] == target
	}

	if len(nums) > 2 {
		copy(numsAdd[1:], nums[2:])
		copy(numsMulti[1:], nums[2:])
		copy(numsConcat[1:], nums[2:])
	}

	return possibleConcat(target, numsAdd) || possibleConcat(target, numsMulti) || possibleConcat(target, numsConcat)
}

// Part 2
// actual : 34612812972206
// got    : 34612814781146
