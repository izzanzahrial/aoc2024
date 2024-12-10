package day10

import (
	"os"
	"strconv"
	"strings"
)

type (
	coordinate [2]int
	direction  [2]int
)

func inBounds(x, y, maxRow, maxCol int) bool {
	return x >= 0 && x < maxRow && y >= 0 && y < maxCol
}

func Part2() int {
	dat, err := os.ReadFile("day10/input")
	if err != nil {
		panic(err)
	}

	islandMap := make([][]int, 0)
	values := strings.Split(string(dat), "\n")
	for i, val := range values {
		row := strings.Split(string(val), "")
		islandMap = append(islandMap, make([]int, 0))
		for _, curr := range row {
			intCurr, _ := strconv.Atoi(curr)
			islandMap[i] = append(islandMap[i], intCurr)
		}
	}

	maxRow := len(islandMap)
	maxCol := len(islandMap[0])
	// search starting coordinate with the number "0"
	var coordinates []coordinate
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			if islandMap[i][j] == 0 {
				coordinates = append(coordinates, coordinate{i, j})
			}
		}
	}

	directions := []direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 1; i <= 9; i++ {
		var nextCoordinates []coordinate
		for _, curr := range coordinates {
			// search all directions
			for _, dir := range directions {
				x := curr[0] + dir[0]
				y := curr[1] + dir[1]
				if inBounds(x, y, maxRow, maxCol) && islandMap[x][y] == i {
					nextCoordinates = append(nextCoordinates, coordinate{x, y})
				}
			}
		}

		coordinates = nextCoordinates
	}

	return len(coordinates)
}
