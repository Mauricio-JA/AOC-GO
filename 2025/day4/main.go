package main

import (
	"os"
	"path/filepath"
	"strings"
)

type position struct {
	X, Y int
}

func canAccess(x, y int, grid [][]string) bool {
	count := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i == x && j == y {
				continue
			}
			if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
				continue
			}
			if grid[i][j] == "@" {
				count++
			}
		}

	}
	return count < 4
}

func main() {
	inputPath := filepath.Join("2025", "day4", "input.txt")
	input, _ := os.ReadFile(inputPath)
	diagram := strings.Split(string(input), "\r\n")

	lenX := len(diagram[0])
	lenY := len(diagram)

	grid := make([][]string, lenY)
	for i := range grid {
		grid[i] = make([]string, lenX)
	}

	// fill grid
	for i := 0; i < lenY; i++ {
		for j := 0; j < lenX; j++ {
			chars := strings.Split(diagram[i], "")
			grid[i][j] = string(chars[j])
		}
	}

	canBeRemoved := make([]position, 0)
	isFirstLoopp := true
	sum1, sum2 := 0, 0

	for isFirstLoopp || len(canBeRemoved) > 0 {

		for _, pos := range canBeRemoved {
			grid[pos.X][pos.Y] = "x"
		}
		canBeRemoved = make([]position, 0)

		for i, row := range grid {
			for j, roll := range row {
				if roll == "@" {
					if canAccess(i, j, grid) {
						if isFirstLoopp {
							sum1++
						}
						sum2++
						canBeRemoved = append(canBeRemoved, position{i, j})
					}
				}
			}
		}

		isFirstLoopp = false

	}

	println("Part 1", sum1)
	println("Part 2", sum2)

}
