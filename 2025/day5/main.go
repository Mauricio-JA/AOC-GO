package main

import (
	"aoc2024/2025/utils"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

type rangePair struct {
	start int
	end   int
}

func checkIsFreshID(id int, freshRanges []rangePair) bool {

	for _, pair := range freshRanges {
		if id >= pair.start && id <= pair.end {
			return true
		}
	}
	return false

}

func main() {
	inputPath := filepath.Join("2025", "day5", "input.txt")
	input, _ := os.ReadFile(inputPath)
	eol := utils.GetEOL()
	databaseLists := strings.Split(string(input), eol+eol)

	ingredientsRanges := strings.Split(databaseLists[0], eol)
	availableIDs := strings.Split(databaseLists[1], eol)

	var freshRanges []rangePair

	start1 := time.Now()

	for _, row := range ingredientsRanges {
		pair := strings.Split(row, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])

		freshRanges = append(freshRanges, rangePair{start, end})
	}

	count := 0
	for _, id := range availableIDs {
		num, _ := strconv.Atoi(id)
		if checkIsFreshID(num, freshRanges) {
			count++
		}
	}

	fmt.Println("Part 1", count, time.Since(start1))

	start2 := time.Now()

	sum := 0
	slices.SortFunc(freshRanges, func(a, b rangePair) int {
		return a.start - b.start
	})

	for i, pair := range freshRanges {
		start, end := pair.start, pair.end
		if i == 0 {
			sum += end - start + 1
			// fmt.Println(pair, end-start+1)
			continue
		}

		prevPair := freshRanges[i-1]
		if end <= prevPair.end {
			// fmt.Println(pair, "ZERO")
			freshRanges[i] = rangePair{start, prevPair.end}
			continue
		}
		if start <= prevPair.end && end > prevPair.end {
			sum += end - prevPair.end
			// fmt.Println(pair, end-prevPair.end, "TRUNC")
			continue
		}
		if start > prevPair.end {
			sum += end - start + 1
			// fmt.Println(pair, end-start+1)
		}

	}

	fmt.Println("Part 2", sum, time.Since(start2))

}
