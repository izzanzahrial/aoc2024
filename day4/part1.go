package day4

import (
	"os"
)

func Part1() int {
	dat, err := os.ReadFile("day4/input")
	if err != nil {
		panic(err)
	}

	input2D := make([][]string, 0)
	i := 0
	for j := 0; j < len(dat); j++ {
		if dat[j] == '\n' {
			i++
			continue
		}
		if len(input2D) <= i {
			input2D = append(input2D, make([]string, 0))
		}
		input2D[i] = append(input2D[i], string(dat[j]))
	}

	var total int
	for i := 0; i < len(input2D); i++ {
		for j := 0; j < len(input2D[i]); j++ {
			if input2D[i][j] == "X" {
				// search east
				if j+1 < len(input2D[i]) {
					if searchXMAS("MAS", "east", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search southeast
				if i+1 < len(input2D) && j+1 < len(input2D[i]) {
					if searchXMAS("MAS", "southeast", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search south
				if i+1 < len(input2D) {
					if searchXMAS("MAS", "south", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search southwest
				if i+1 < len(input2D) && j-1 >= 0 {
					if searchXMAS("MAS", "southwest", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search west
				if j-1 >= 0 {
					if searchXMAS("MAS", "west", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search northwest
				if i-1 >= 0 && j-1 >= 0 {
					if searchXMAS("MAS", "northwest", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search north
				if i-1 >= 0 {
					if searchXMAS("MAS", "north", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}

				// search northeast
				if i-1 >= 0 && j+1 < len(input2D[i]) {
					if searchXMAS("MAS", "northeast", i, j, len(input2D), len(input2D[i]), input2D) {
						total++
					}
				}
			}
		}
	}

	return total
}

func searchXMAS(keyword, direction string, i, j, maxRow, maxCol int, input2D [][]string) bool {
	if len(keyword) == 0 {
		return true
	}

	currKeyword := string(keyword[0])
	switch direction {
	case "east":
		// search east
		if j+1 < maxCol && input2D[i][j+1] == currKeyword {
			if searchXMAS(keyword[1:], "east", i, j+1, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "southeast":
		// search southeast
		if i+1 < maxRow && j+1 < maxCol && input2D[i+1][j+1] == currKeyword {
			if searchXMAS(keyword[1:], "southeast", i+1, j+1, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "south":
		// search south
		if i+1 < maxRow && input2D[i+1][j] == currKeyword {
			if searchXMAS(keyword[1:], "south", i+1, j, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "southwest":
		// search southwest
		if i+1 < maxRow && j-1 >= 0 && input2D[i+1][j-1] == currKeyword {
			if searchXMAS(keyword[1:], "southwest", i+1, j-1, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "west":
		// search west
		if j-1 >= 0 && input2D[i][j-1] == currKeyword {
			if searchXMAS(keyword[1:], "west", i, j-1, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "northwest":
		// search northwest
		if i-1 >= 0 && j-1 >= 0 && input2D[i-1][j-1] == currKeyword {
			if searchXMAS(keyword[1:], "northwest", i-1, j-1, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "north":
		// search north
		if i-1 >= 0 && input2D[i-1][j] == currKeyword {
			if searchXMAS(keyword[1:], "north", i-1, j, maxRow, maxCol, input2D) {
				return true
			}
		}
	case "northeast":
		// search northeast
		if i-1 >= 0 && j+1 < maxCol && input2D[i-1][j+1] == currKeyword {
			if searchXMAS(keyword[1:], "northeast", i-1, j+1, maxRow, maxCol, input2D) {
				return true
			}
		}
	}

	return false
}
