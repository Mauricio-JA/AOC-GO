package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func insertStone(stonesMap map[int]int, stone int, addCount int) {
	val, ok := stonesMap[stone]
	if ok {
		stonesMap[stone] = val + addCount
	} else {
		stonesMap[stone] = addCount
	}
}

func intLen(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func blink(stones map[int]int, blinkCount int) map[int]int {
	// fmt.Println(blinkCount)
	blinkCount--
	if blinkCount < 0 {
		return stones
	}
	newStones := make(map[int]int)
	for stone := range maps.Keys(stones) {
		if stone == 0 {
			insertStone(newStones, 1, stones[stone])
			continue
		}
		stoneLen := intLen(stone)
		if stoneLen%2 == 0 {
			halfIndex := intLen(stone) / 2
			divisor := int(math.Pow10(halfIndex))
			leftStone := stone / divisor
			rigthStone := stone % divisor
			insertStone(newStones, leftStone, stones[stone])
			insertStone(newStones, rigthStone, stones[stone])
		} else {
			insertStone(newStones, stone*2024, stones[stone])
		}
	}
	return blink(newStones, blinkCount)
}

func sumStonesCount(stonesMap map[int]int) int {
	sum := 0
	for k := range maps.Keys(stonesMap) {
		sum += stonesMap[k]
	}
	return sum
}

func main() {
	inputPath := filepath.Join("2024", "day11", "input.txt")
	input, _ := os.ReadFile(inputPath)
	initialStones := strings.Split(string(input), " ")
	stonesMap := make(map[int]int)
	for _, stone := range initialStones {
		number, _ := strconv.Atoi(stone)
		insertStone(stonesMap, number, 1)
	}
	// startTime2 := time.Now()
	newStonesPart1 := blink(stonesMap, 25)
	fmt.Println("Part 1", sumStonesCount(newStonesPart1))

	newStonesPart2 := blink(stonesMap, 75)
	fmt.Println("Part 2", sumStonesCount(newStonesPart2))
	// fmt.Println("Part 2", sum, "tooks", time.Since(startTime2))

}
