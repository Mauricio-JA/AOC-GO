package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type height struct {
	x     int
	y     int
	value int
}

func findHikingTrails(trailhead height, topoMap [][]int) (peaks []height) {

	if trailhead.value == 9 {
		peaks = append(peaks, trailhead)
		return peaks
	} else {
		x, y, value := trailhead.x, trailhead.y, trailhead.value // solo para hacerlo menos verboso
		//check up
		if x > 0 && topoMap[x-1][y] == value+1 {
			findedPeaks := findHikingTrails(height{x - 1, y, value + 1}, topoMap)
			peaks = append(peaks, findedPeaks...)
		}
		//check left
		if y > 0 && topoMap[x][y-1] == value+1 {
			findedPeaks := findHikingTrails(height{x, y - 1, value + 1}, topoMap)
			peaks = append(peaks, findedPeaks...)
		}
		//check down
		if x < len(topoMap)-1 && topoMap[x+1][y] == value+1 {
			findedPeaks := findHikingTrails(height{x + 1, y, value + 1}, topoMap)
			peaks = append(peaks, findedPeaks...)
		}
		//check left
		if y < len(topoMap[x])-1 && topoMap[x][y+1] == value+1 {
			findedPeaks := findHikingTrails(height{x, y + 1, value + 1}, topoMap)
			peaks = append(peaks, findedPeaks...)
		}
	}

	return peaks
}

func removeRepeated(peaks []height) (result []height) {
	for _, peak := range peaks {
		if !slices.Contains(result, peak) {
			result = append(result, peak)
		}
	}
	return result
}

func main() {
	inputPath := filepath.Join("2024", "day10", "input.txt")
	input, _ := os.ReadFile(inputPath)
	rows := strings.Split(string(input), "\r\n")

	topoMap := make([][]int, len(rows))
	for i, row := range rows {
		line := strings.Split(row, "")
		for _, char := range line {
			number, _ := strconv.Atoi(char)
			topoMap[i] = append(topoMap[i], number)
		}
	}

	score, rating := 0, 0
	for i, row := range topoMap {
		for j, value := range row {
			if value == 0 {
				peaks := findHikingTrails(height{i, j, value}, topoMap)
				score += len(removeRepeated(peaks))
				rating += len(peaks)
			}
		}
	}
	fmt.Println("Part 1", score)
	fmt.Println("Part 2", rating)

}
