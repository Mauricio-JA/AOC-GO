package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func calcJoltageRating(batteries []string) int {
	var max = int(0)
	for i := 0; i < len(batteries)-1; i++ {
		digit, _ := strconv.Atoi(batteries[i])
		numbersToEvaluate := batteries[i+1:]
		for j := 0; j < len(numbersToEvaluate); j++ {
			digit2, _ := strconv.Atoi(numbersToEvaluate[j])
			joltage := digit*10 + digit2
			if joltage > max {
				max = joltage
			}
		}
	}
	return max
}

type maxJoltage struct {
	number int
	index  int
}

func calcJoltageRating12(batteries []string) int {
	var digits []int = []int{}
	var startIndex = 0

	for len(digits) < 12 {

		// fmt.Println(startIndex, len(batteries)-12+len(digits))

		numbersToEvaluate := batteries[startIndex : len(batteries)-11+len(digits)]
		var max maxJoltage = maxJoltage{number: 0, index: startIndex}

		for j := 0; j < len(numbersToEvaluate); j++ {
			joltage, _ := strconv.Atoi(numbersToEvaluate[j])
			if joltage > max.number {
				max.number = joltage
				max.index = j
			}
		}

		digits = append(digits, max.number)
		startIndex += max.index + 1

	}
	// fmt.Println(digits)

	var result int
	for i := 0; i < len(digits); i++ {
		// fmt.Println(i, digits[i], int(math.Pow10(12-i-1))*digits[i])
		result += int(math.Pow10(12-i-1)) * digits[i]
	}
	return result

}

func main() {
	inputPath := filepath.Join("2025", "day3", "input.txt")
	input, _ := os.ReadFile(inputPath)
	banks := strings.Split(string(input), "\r\n")
	sum1, sum2 := 0, 0
	for _, bank := range banks {
		joltageRating := calcJoltageRating(strings.Split(bank, ""))
		sum1 += joltageRating

		joltageRating12 := calcJoltageRating12(strings.Split(bank, ""))
		// fmt.Println(joltageRating12)
		sum2 += joltageRating12
	}

	fmt.Println("Part 1", sum1)
	fmt.Println("Part 2", sum2)

}
