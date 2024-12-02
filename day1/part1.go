package day1

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	dat, err := os.ReadFile("day1/input")
	if err != nil {
		panic(err)
	}
	values := strings.Split(string(dat), "\n")

	list1 := []int{}
	list2 := []int{}
	for _, val := range values {
		curr := strings.Split(val, "   ")
		num1, _ := strconv.Atoi(curr[0])
		num2, _ := strconv.Atoi(curr[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var result int
	for i := 0; i < len(list1); i++ {
		curr := math.Abs(float64(list1[i] - list2[i]))
		result += int(curr)
	}

	return result
}
