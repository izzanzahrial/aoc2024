package day12

import (
	"fmt"
	"os"
	"strings"
)

type direction [2]int

func Part1() int {
	dat, err := os.ReadFile("day12/input")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(dat), "\n")
	gardenMap := make([][]string, len(values))
	for i, val := range values {
		row := strings.Split(string(val), "")
		currCol := make([]string, 0)
		for _, curr := range row {
			currCol = append(currCol, curr)
		}
		gardenMap[i] = currCol
	}

	directions := []direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var total int
	visited := make(map[[2]int]struct{}, 0)
	for i := 0; i < len(gardenMap); i++ {
		for j := 0; j < len(gardenMap[i]); j++ {
			if _, ok := visited[[2]int{i, j}]; ok {
				continue
			}

			var gates, plot int
			currPlant := gardenMap[i][j]
			queueCoors := [][2]int{{i, j}}
			visited[[2]int{i, j}] = struct{}{}
			gates, plot = checkGatesAndPlots(currPlant, directions, queueCoors, gardenMap, visited)
			fmt.Println(currPlant, "gates:", gates, "plots:", plot, "total:", gates*plot)
			// plot + 1 because we add the current plot
			total += gates * (plot + 1)
		}
	}

	return total
}

func checkGatesAndPlots(plant string, directions []direction, queue [][2]int, gardenMap [][]string, visited map[[2]int]struct{}) (gates int, plot int) {
	// keep checking neighbours
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]

		// check neighbours for each direction
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]

			// check if inbound
			if inBounds(newX, newY, gardenMap) {

				// if inbound but not same plant then count as gate
				if gardenMap[newX][newY] != plant {
					gates++

					// if inbound and same plant then check if it has been visited
					// if visited then continue
				} else if _, ok := visited[[2]int{newX, newY}]; ok {
					continue

					// if not visited then add count of plots
					// add to queue
					// mark visited
				} else if gardenMap[newX][newY] == plant {
					queue = append(queue, [2]int{newX, newY})
					visited[[2]int{newX, newY}] = struct{}{}
					plot++
				}

				// if not in bounds then count as gate
			} else {
				gates++
			}
		}
		queue = queue[1:]
	}

	return gates, plot
}

func inBounds(x, y int, gardenMap [][]string) bool {
	return x >= 0 && x < len(gardenMap) && y >= 0 && y < len(gardenMap[0])
}
