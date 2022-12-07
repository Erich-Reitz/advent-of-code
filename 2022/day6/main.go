package main

import (
	"fmt"
	"github.com/Erich-Reitz/commonGo"
)

func fleshFreqMapWithFirstNCharacteres(line string, queueLength int) map[string]int {
	freqMap := make(map[string]int, 0)
	for i := 0; i < queueLength; i++ {
		freqMap[string(line[i])] += 1
	}
	return freqMap
}

func firstIndexWhereWindowIsAllUnqiueChars(line string, queueLength int) int {
	freqMap := fleshFreqMapWithFirstNCharacteres(line, queueLength)
	leftIndex := 0
	rightIndex := queueLength
	for ; rightIndex < len(line); rightIndex++ {
		if len(freqMap) == queueLength {
			return rightIndex
		}
		leftMostElement := string(line[leftIndex])
		freqMap[leftMostElement] -= 1
		leftIndex += 1
		if freqMap[leftMostElement] == 0 {
			delete(freqMap, leftMostElement)
		}
		rightMostElement := string(line[rightIndex])
		freqMap[rightMostElement] += 1
	}
	return -1
}

func part1(line string) {
	fmt.Println(firstIndexWhereWindowIsAllUnqiueChars(line, 4))
}

func part2(line string) {
	fmt.Println(firstIndexWhereWindowIsAllUnqiueChars(line, 14))
}

func main() {
	data := advent.ReadFileAsString("input.txt")
	part1(data)
	part2(data)
}
