package day11

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part2() int {
	dat, err := os.ReadFile("day11/input")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(dat), " ")
	var result int
	cache := make(map[int][]int)
	for _, val := range values {
		intVal, _ := strconv.Atoi(val)
		result += getResultBlinks(intVal, 75, cache)
	}
	return result
}

func getResultBlinks(stone, blink int, cache map[int][]int) int {
	// check if stone already in cache
	if _, ok := cache[stone]; ok {
		// check if stone for the current blink already in cache
		if cache[stone][blink-1] != 0 {
			return cache[stone][blink-1]
		}
	} else {
		cache[stone] = make([]int, 75)
	}

	if blink == 1 {
		cache[stone][blink-1] = len(createNewStone(stone))
		return len(createNewStone(stone))
	}

	// recursively count the number of blinks by totaling the number of blinks
	// for each previous blink
	// blinks = 75 nums = [1, 2, 3, 4, 5, 6]
	// blinks = 74 nums = [1] , [2], [3], [4], [5], [6]
	// blinks = 73 the result from blinks = 74
	// .....
	// all the result will go up back again to the blinks = 75
	// in short divide the work into smaller pieces
	var result int
	for _, stone := range createNewStone(stone) {
		result += getResultBlinks(stone, blink-1, cache)
	}

	cache[stone][blink-1] = result
	return result
}

func createNewStone(stone int) []int {
	var result []int

	if stone == 0 {
		result = append(result, 1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		s1, s2 := splitStone(stone)
		result = append(result, s1, s2)
	} else {
		result = append(result, stone*2024)
	}

	return result
}

func splitStone(stone int) (int, int) {
	stoneString := strconv.Itoa(stone)
	stone1, stone2 := stoneString[:len(stoneString)/2], stoneString[len(stoneString)/2:]
	return FetchNumFromStringIgnoringNonNumeric(stone1), FetchNumFromStringIgnoringNonNumeric(stone2)
}

func FetchNumFromStringIgnoringNonNumeric(line string) int {
	var build strings.Builder
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(localNum)
	}
	return 0
}
