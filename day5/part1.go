package day5

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() int {
	rawRules, err := os.ReadFile("day5/rules")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(rawRules), "\n")
	rulesMap := make(map[int][]int, len(values))
	for _, val := range values {
		nums := strings.Split(string(val), "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		rulesMap[num1] = append(rulesMap[num1], num2)
	}
	// fmt.Println(rulesMap)

	rawInput, err := os.ReadFile("day5/input")
	if err != nil {
		panic(err)
	}

	var total int
	for _, val := range strings.Split(string(rawInput), "\n") {
		allInOrder := true
		curr := strings.Split(val, ",")
		intCurr := convertSlice(curr)
		for i, num := range curr {
			num, _ := strconv.Atoi(num)
			if !inOrder(i, intCurr, rulesMap[num]) {
				allInOrder = false
			}
		}

		if allInOrder {
			middleIdx := len(intCurr) / 2
			total += intCurr[middleIdx]
		}
	}

	return total
}

func inOrder(index int, collection []int, rules []int) bool {
	for i, val := range collection {
		if i == index {
			continue
		}

		if slices.Contains(rules, val) && index > i {
			return false
		}
	}

	return true
}

func convertSlice(slice []string) []int {
	result := make([]int, len(slice))
	for i, val := range slice {
		num, _ := strconv.Atoi(val)
		result[i] = num
	}
	return result
}
