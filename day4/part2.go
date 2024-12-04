package day4

import (
	"os"
)

func Part2() int {
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
	for i := 1; i < len(input2D)-1; i++ {
		for j := 1; j < len(input2D[i])-1; j++ {
			if input2D[i][j] == "A" {
				if input2D[i-1][j-1] == "M" && input2D[i-1][j+1] == "M" && input2D[i+1][j-1] == "S" && input2D[i+1][j+1] == "S" {
					total++
				} else if input2D[i-1][j-1] == "S" && input2D[i-1][j+1] == "S" && input2D[i+1][j-1] == "M" && input2D[i+1][j+1] == "M" {
					total++
				} else if input2D[i-1][j-1] == "S" && input2D[i-1][j+1] == "M" && input2D[i+1][j-1] == "S" && input2D[i+1][j+1] == "M" {
					total++
				} else if input2D[i-1][j-1] == "M" && input2D[i-1][j+1] == "S" && input2D[i+1][j-1] == "M" && input2D[i+1][j+1] == "S" {
					total++
				}
			}
		}
	}

	return total
}
