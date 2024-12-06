package day5

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

type RulesMap map[int][]int

// Compare each number within each other's rules
// if not in each other's rules return 0
func (r RulesMap) Compare(k1, k2 int) int {
	// if second number rules contain first number return -1
	// meaning second number should come before first number
	if slices.Contains(r[k2], k1) {
		return -1
	}

	// if first number rules contain second number return 1
	// meaning first number should come before second number
	if slices.Contains(r[k1], k2) {
		return 1
	}

	// first number and second number are not in each other's rules
	return 0
}

func Part2() int {
	rawRules, err := os.ReadFile("day5/rules")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(rawRules), "\n")
	rulesMap := make(RulesMap, len(values))
	for _, val := range values {
		nums := strings.Split(string(val), "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		rulesMap[num1] = append(rulesMap[num1], num2)
	}

	rawInput, err := os.ReadFile("day5/input")
	if err != nil {
		panic(err)
	}

	var total int
	for _, val := range strings.Split(string(rawInput), "\n") {
		curr := strings.Split(val, ",")
		intCurr := convertSlice(curr)
	out:
		for i, num := range curr {
			num, _ := strconv.Atoi(num)
			if !inOrder(i, intCurr, rulesMap[num]) {
				slices.SortFunc(intCurr, rulesMap.Compare)
				total += intCurr[len(intCurr)/2]
				break out
			}
		}
	}

	return total
}
