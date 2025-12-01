package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	inputPath := filepath.Join("2025", "day1", "input.txt")
	input, _ := os.ReadFile(inputPath)

	var rotations = strings.Split(string(input), "\r\n")
	dial, count := 50, 0

	for _, rotation := range rotations {
		var pair = strings.SplitN(rotation, "", 2)
		direction := pair[0]
		value, _ := strconv.Atoi(pair[1])
		if direction == "R" {
			dial = (dial + value) % 100
		} else {
			if value%100 > dial {
				dial = 100 + dial - (value % 100)
			} else {
				dial -= value % 100
			}
		}
		if dial == 0 {
			count++
		}
	}

	fmt.Println("Part 1", count)
}
