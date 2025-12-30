package main

import (
	"aoc2024/2025/utils"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func SplitByIndexes(s string, indexes []int) []string {
	var result []string
	lastIndex := 0

	for _, i := range indexes {
		if i > len(s) {
			i = len(s)
		}
		if i > lastIndex {
			result = append(result, s[lastIndex:i])
			lastIndex = i
		}
	}

	if lastIndex < len(s) {
		result = append(result, s[lastIndex:])
	}

	return result
}
func SplitByColumns(rows []string) [][]string {
	spacePositions := make(map[int]int)
	max := len(rows)
	for _, row := range rows {
		for i, char := range strings.Split(row, "") {
			if char == " " {
				spacePositions[i]++
			}
		}
	}

	dividerIndexes := []int{}
	for key, value := range spacePositions {
		if value == max {
			dividerIndexes = append(dividerIndexes, key)
		}
	}
	sort.Slice(dividerIndexes, func(i, j int) bool {
		return dividerIndexes[i] < dividerIndexes[j]
	})

	result := [][]string{}
	for _, row := range rows {
		result = append(result, SplitByIndexes(row, dividerIndexes))
	}
	return result

}

func main() {
	inputPath := filepath.Join("2025", "day6", "input.txt")
	input, _ := os.ReadFile(inputPath)
	eol := utils.GetEOL()
	rows := strings.Split(string(input), eol)

	var trimmed [][]string
	for _, row := range rows {
		fields := strings.Fields(row)
		trimmed = append(trimmed, fields)
	}

	problemsCount := len(trimmed[0])
	problemLength := len(trimmed)
	total := 0

	start := time.Now()

	for i := 0; i < problemsCount; i++ {
		sign := trimmed[problemLength-1][i]
		result := 0
		for j := 0; j < problemLength-1; j++ {
			value, _ := strconv.Atoi(trimmed[j][i])
			if j == 0 {
				result += value
				continue
			}
			if sign == "+" {
				result += value
			} else {
				result *= value
			}

		}
		total += result
	}

	fmt.Println("Part 1", total, time.Since(start))

	start2 := time.Now()

	splittedProblems := SplitByColumns(rows[:len(rows)-1])

	signs := strings.Fields(rows[len(rows)-1])

	// fmt.Println(splittedProblems)
	// fmt.Println(signs)

	problemsCount = len(splittedProblems[0])
	problemLength = len(splittedProblems)
	total = 0

	for i := problemsCount - 1; i >= 0; i-- {
		sign := signs[i]
		result := 0
		numberLength := len(splittedProblems[0][i])

		for k := numberLength - 1; k >= 0; k-- {
			cephalodNumber := 0
			for j := problemLength - 1; j >= 0; j-- {
				char := strings.Split(splittedProblems[j][i], "")[k]
				if char != " " {
					value, _ := strconv.Atoi(char)
					cephalodNumberLength := 0
					if cephalodNumber != 0 {
						cephalodNumberLength = int(math.Floor(math.Log10(float64(cephalodNumber)))) + 1
					}
					// fmt.Println("char", char, value, cephalodNumber, cephalodNumberLength)

					cephalodNumber += value * int(math.Pow10(cephalodNumberLength))
				}
			}
			// fmt.Println("CEPHALOD NUMBER", cephalodNumber)
			if result == 0 || cephalodNumber == 0 {
				result += cephalodNumber
				continue
			}
			if sign == "+" {
				result += cephalodNumber
			} else {
				result *= cephalodNumber
			}
		}
		// fmt.Println("RESULT", result)

		total += result
	}

	fmt.Println("Part 2", total, time.Since(start2))
}
