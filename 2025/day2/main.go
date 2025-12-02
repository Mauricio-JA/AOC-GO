package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func isValidID1(id int) bool {
	text := strconv.Itoa(id)
	if len(text)%2 == 0 {
		half := len(text) / 2
		firstHalf := text[:half]
		secondHalf := text[half:]
		return firstHalf != secondHalf
	} else {
		return true
	}
}

func isValidID2(id int) bool {
	text := strconv.Itoa(id)
	for i := 2; i <= len(text); i++ {
		if len(text)%i == 0 {
			chunkSize := len(text) / i
			firstChunk := text[:chunkSize]
			isValid := false
			for j := 1; j < i; j++ {
				nextChunk := text[j*chunkSize : (j*chunkSize)+chunkSize]
				if firstChunk != nextChunk {
					isValid = true
					break
				}
			}
			if !isValid {
				return false
			}
			continue
		}
	}
	return true

}

func main() {
	inputPath := filepath.Join("2025", "day2", "input.txt")
	input, _ := os.ReadFile(inputPath)
	var productRanges = strings.Split(string(input), ",")
	sum1, sum2 := 0, 0

	for _, rangeIDs := range productRanges {
		var IDs = strings.Split(rangeIDs, "-")
		startID, _ := strconv.Atoi(IDs[0])
		endID, _ := strconv.Atoi(IDs[1])

		for i := startID; i <= endID; i++ {
			if !isValidID1(i) {
				//fmt.Println(i)
				sum1 += i
			}
			if !isValidID2(i) {
				//fmt.Println(i)
				sum2 += i
			}

		}

	}
	fmt.Println("Part 1", sum1)
	fmt.Println("Part 2", sum2)

}
