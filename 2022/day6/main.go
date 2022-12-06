package main

import (
	"fmt"
	"github.com/Erich-Reitz/commonGo"
)

func removeLeftMostElement(queue []string, counts map[string]int) ([]string, map[string]int) {
	topElement := queue[0]
	counts[topElement] -= 1
	queue = queue[1:]
	return queue, counts
}

func addElementToRight(queue []string, counts map[string]int, element string) ([]string, map[string]int) {
	queue = append(queue, element)
	counts[element] += 1
	return queue, counts
}

func allElementsInQueueAreUnique(queue []string, counts map[string]int) bool {
	allUnique := true
	for _, item := range queue {
		if counts[item] != 1 {
			allUnique = false
		}
	}
	return allUnique
}

func firstIndexWhereWindowIsAllUnqiueChars(line string, queueLength int) int {
	queue := make([]string, 0)
	counts := make(map[string]int)
	for index, char := range line {
		if len(queue) >= queueLength {
			queue, counts = removeLeftMostElement(queue, counts)
		}
		queue, counts = addElementToRight(queue, counts, string(char))
		if len(queue) >= queueLength && allElementsInQueueAreUnique(queue, counts) {
			return index
		}
	}
	return -1
}

func part1(line string) {
	fmt.Println(firstIndexWhereWindowIsAllUnqiueChars(line, 4) + 1)
}

func part2(line string) {
	fmt.Println(firstIndexWhereWindowIsAllUnqiueChars(line, 14) + 1)
}

func main() {
	data := advent.ReadFileAsString("input.txt")
	part1(data)
	part2(data)
}
