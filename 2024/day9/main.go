package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func repeatSlice(elem string, number int) []string {
	result := make([]string, number)
	for i := range result {
		result[i] = elem
	}
	return result
}

func formatDiskMap(diskInput []string) (diskMap []string, maxId int) {
	fileId := 0
	for i, block := range diskInput {
		blockNumber, _ := strconv.Atoi(block)
		if i%2 == 0 {
			file := repeatSlice(strconv.Itoa(fileId), blockNumber)
			diskMap = append(diskMap, file...)
			fileId++
		} else {
			file := repeatSlice(".", blockNumber)
			diskMap = append(diskMap, file...)
		}
	}
	return diskMap, fileId - 1
}

func checkIsCompressed(diskMap []string) bool {
	trimmed := strings.Trim(strings.Join(diskMap, ""), ".")
	return !strings.Contains(trimmed, ".")
}

func lastBlockFileIndex(diskMap []string) int {
	lastIndex := len(diskMap) - 1
	for lastIndex > 0 {
		lastElement := diskMap[lastIndex]
		if lastElement != "." {
			break
		}
		lastIndex--
	}
	return lastIndex
}

func compactPart1(diskMap []string) []string {
	diskMapCopy := make([]string, len(diskMap))
	copy(diskMapCopy, diskMap)

	for !checkIsCompressed(diskMapCopy) {
		freeBlockIndex := slices.Index(diskMapCopy, ".")
		endFileBlockIndex := lastBlockFileIndex(diskMapCopy)
		diskMapCopy[freeBlockIndex] = diskMapCopy[endFileBlockIndex]
		diskMapCopy[endFileBlockIndex] = "."
	}
	return diskMapCopy
}

func getWholeFile(diskMap []string, fileId int) (startIndex int, length int) {
	value := strconv.Itoa(fileId)
	startIndex = slices.Index(diskMap, value)
	if startIndex == -1 {
		return
	}
	i := startIndex
	for i < len(diskMap) {
		if diskMap[i] == value {
			length++
		}
		i++
	}
	return startIndex, length
}

func findFreeSpaceToMove(diskMap []string, capacity int) (freeIndex int) {
	i := 0
	count := 0
	for i < len(diskMap) {
		if diskMap[i] == "." {
			count++
		} else {
			count = 0
		}
		if count > 0 && count >= capacity {
			freeIndex = i - count + 1
			break
		}
		i++
	}
	return freeIndex
}

func compressPart2(diskMap []string, maxFileId int) []string {
	diskMapCopy := make([]string, len(diskMap))
	copy(diskMapCopy, diskMap)

	for maxFileId > 0 {
		startIndex, length := getWholeFile(diskMapCopy, maxFileId)
		freeIndex := findFreeSpaceToMove(diskMapCopy[:startIndex], length)
		if freeIndex > 0 {
			value := strconv.Itoa(maxFileId)
			for i := 0; i < length; i++ {
				diskMapCopy[freeIndex+i] = value
				diskMapCopy[startIndex+i] = "."
			}
		}
		maxFileId--
	}
	return diskMapCopy
}

func main() {
	inputPath := filepath.Join("2024", "day9", "input.txt")
	input, _ := os.ReadFile(inputPath)
	diskInput := strings.Split(string(input), "")
	diskMap, maxFileId := formatDiskMap(diskInput)

	// fmt.Println(diskMap)
	compactedDisk1 := compactPart1(diskMap)
	compactedDisk2 := compressPart2(diskMap, maxFileId)

	checksum1, checksum2 := 0, 0
	for i := 0; i < len(diskMap); i++ {
		block1 := compactedDisk1[i]
		block2 := compactedDisk2[i]
		if block1 != "." {
			blockNumber, _ := strconv.Atoi(block1)
			checksum1 += i * blockNumber
		}
		if block2 != "." {
			blockNumber, _ := strconv.Atoi(block2)
			checksum2 += i * blockNumber
		}
	}

	fmt.Println("Part 1", checksum1)
	fmt.Println("Part 2", checksum2)

}
