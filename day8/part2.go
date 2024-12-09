package day8

import (
	"os"
	"strings"
)

func Part2() int {
	dat, err := os.ReadFile("day8/input")
	if err != nil {
		panic(err)
	}

	// Parse the grid
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	// Create a map of frequencies to their positions
	frequencies := make(map[string][][2]int)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				freq := string(char)
				frequencies[freq] = append(frequencies[freq], [2]int{i, j})
			}
		}
	}

	// Track antinodes
	antinodes := make(map[[2]int]bool)

	// For each frequency
	for _, positions := range frequencies {
		// Check each pair of antennas with the same frequency
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				p1 := positions[i]
				p2 := positions[j]

				// Add the antenna positions themselves as antinodes
				// (as they are collinear with other antennas)
				if len(positions) > 1 {
					antinodes[p1] = true
					antinodes[p2] = true
				}

				// Find all points collinear with these two antennas
				findCollinearPoints(p1, p2, len(lines), len(lines[0]), antinodes)
			}
		}
	}

	return len(antinodes)
}

func findCollinearPoints(p1, p2 [2]int, rows, cols int, antinodes map[[2]int]bool) {
	// Calculate direction vector
	dy := p2[0] - p1[0]
	dx := p2[1] - p1[1]

	// Get GCD to find the smallest step size
	gcd := GCD(abs(dx), abs(dy))
	if gcd != 0 {
		dx /= gcd
		dy /= gcd
	}

	// Check all points along the line in both directions
	curr := [2]int{p1[0], p1[1]}

	// Forward direction
	for curr[0] >= 0 && curr[0] < rows && curr[1] >= 0 && curr[1] < cols {
		antinodes[curr] = true
		curr[0] += dy
		curr[1] += dx
	}

	// Backward direction
	curr = [2]int{p1[0] - dy, p1[1] - dx}
	for curr[0] >= 0 && curr[0] < rows && curr[1] >= 0 && curr[1] < cols {
		antinodes[curr] = true
		curr[0] -= dy
		curr[1] -= dx
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
