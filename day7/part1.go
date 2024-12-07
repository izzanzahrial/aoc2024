package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
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
		if possible1(target, nums[0], 1, nums) {
			total += target
			fmt.Println("possible target:", target, "nums:", nums)
		} else {
			fmt.Println("not possible target:", target, "nums:", nums)
		}

		// using possible2
		// if possible2(target, nums) {
		// 	total += target
		// 	fmt.Println("possible target:", target, "nums:", nums)
		// } else {
		// 	fmt.Println("not possible target:", target, "nums:", nums)
		// }
	}

	return total
}

func possible1(target, total, idx int, nums []int) bool {
	if total == target {
		return true
	}

	if idx >= len(nums) || total > target {
		return false
	}

	return possible1(target, total+nums[idx], idx+1, nums) || possible1(target, total*nums[idx], idx+1, nums)
}

func possible2(target int, nums []int) bool {
	if nums[0] == target {
		return true
	}

	if nums[0] > target || len(nums) == 1 {
		return false
	}

	numsAdd := make([]int, len(nums)-1)
	numsMulti := make([]int, len(nums)-1)
	numsAdd[0] = nums[0] + nums[1]
	numsMulti[0] = nums[0] * nums[1]

	if len(nums) == 2 {
		return numsAdd[0] == target || numsMulti[0] == target
	}

	if len(nums) > 2 {
		copy(numsAdd[1:], nums[2:])
		copy(numsMulti[1:], nums[2:])
	}

	return possible2(target, numsAdd) || possible2(target, numsMulti)
}

// Part 1
// actual : 538191549061
// got    : 538191549251
