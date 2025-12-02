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
	dial, count1, count2 := 50, 0, 0

	for _, rotation := range rotations {
		var pair = strings.SplitN(rotation, "", 2)
		fmt.Println(pair, dial, count1, count2)

		direction := pair[0]
		value, _ := strconv.Atoi(pair[1])
		if direction == "R" {

			count2 += (dial + value) / 100
			dial = (dial + value) % 100

		} else {

			if dial != 0 {
				count2 += (100 - dial + value) / 100
			} else {
				count2 += (dial + value) / 100
			}

			if value%100 > dial {
				dial = 100 + dial - (value % 100)
			} else {
				dial -= value % 100
			}
		}
		if dial == 0 {
			count1++
		}
	}

	fmt.Println("Part 1", count1)
	fmt.Println("Part 2", count2)
}
