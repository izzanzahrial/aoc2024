package day9

import (
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	dat, err := os.ReadFile("day9/input")
	if err != nil {
		panic(err)
	}

	diskMap := strings.Split(string(dat), "")
	newDiskMap := make([]string, 0)

	isFile := true
	id := 0
	for _, val := range diskMap {
		strId := strconv.Itoa(id)
		intVal, _ := strconv.Atoi(val)
		if isFile {
			for i := 0; i < intVal; i++ {
				newDiskMap = append(newDiskMap, strId)
			}
			id++
			isFile = false
		} else {
			for i := 0; i < intVal; i++ {
				newDiskMap = append(newDiskMap, ".")
			}
			isFile = true
		}

	}

	left := 0
	right := len(newDiskMap) - 1
	for left < right {
		if newDiskMap[left] == "." && newDiskMap[right] != "." {
			newDiskMap[left], newDiskMap[right] = newDiskMap[right], newDiskMap[left]
			left++
			right--
		}

		if newDiskMap[left] != "." {
			left++
		}

		if newDiskMap[right] == "." {
			right--
		}
	}

	var result int
	for i := 0; i < len(newDiskMap); i++ {
		if newDiskMap[i] == "." {
			continue
		}

		currNum, _ := strconv.Atoi(string(newDiskMap[i]))
		result += (i * currNum)
	}

	return result
}
