package day8

import (
	"os"
	"strings"
)

type Direction [2]int

func Part1() int {
	dat, err := os.ReadFile("day8/input")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(dat), "\n")
	anthenasMap := make([][]string, len(values))
	for i, val := range values {
		row := strings.Split(string(val), "")
		for _, curr := range row {
			anthenasMap[i] = append(anthenasMap[i], curr)
		}
	}

	anthenas := make(map[string][][2]int)
	for i := 0; i < len(anthenasMap); i++ {
		for j := 0; j < len(anthenasMap[i]); j++ {
			if anthenasMap[i][j] != "." {
				anthenas[anthenasMap[i][j]] = append(anthenas[anthenasMap[i][j]], [2]int{i, j})
			}
		}
	}

	total := new(int)
	antinodesMap := make(map[[2]int]struct{}, 0)
	for _, v := range anthenas {
		for j := 0; j < len(v)-1; j++ {
			for k := j + 1; k < len(v); k++ {
				isAntinode(total, v[j], v[k], anthenasMap, antinodesMap)
			}
		}
	}

	return len(antinodesMap)
}

func isAntinode(total *int, x, y [2]int, anthenasMap [][]string, antinodesMap map[[2]int]struct{}) {
	ix, iy := x[0], x[1]
	jx, jy := y[0], y[1]

	kx, ky := ix-(jx-ix), iy-(jy-iy)
	lx, ly := jx+(jx-ix), jy+(jy-iy)

	if inBounds(kx, ky, anthenasMap) {
		*total += 1
		antinodesMap[[2]int{kx, ky}] = struct{}{}
	}

	if inBounds(lx, ly, anthenasMap) {
		*total += 1
		antinodesMap[[2]int{lx, ly}] = struct{}{}
	}
}

func inBounds(x, y int, athenasMap [][]string) bool {
	return x >= 0 && x < len(athenasMap) && y >= 0 && y < len(athenasMap[0])
}
