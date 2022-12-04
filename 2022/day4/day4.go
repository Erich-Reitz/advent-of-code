package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileAsString(filepath string) string {
	dat, err := os.ReadFile(filepath)
	check(err)
	return string(dat)
}

func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min < i {
			min = i
		}
	}

	return min
}

func doesRangeContainOther(elf1begin, elf1end, elf2begin, elf2end int) bool {
	allMin := MinOf(elf1begin, elf1end, elf2begin, elf2end)
	allMax := MaxOf(elf1begin, elf1end, elf2begin, elf2end)

	if (allMin == elf1begin && allMax == elf1end) || (allMin == elf2begin && allMax == elf2end) {
		return true
	}

	return false
}

func doAnyRangesOverlap(elf1begin, elf1end, elf2begin, elf2end int) bool {
	if elf2begin >= elf1begin && elf2begin <= elf1end {
		return true
	}
	if elf1begin >= elf2begin && elf1begin <= elf2end {
		return true
	}
	return false
}

func splitTwoIntsByString(line, splitBy string) (int, int) {
	splitString := strings.Split(line, splitBy)
	num1, err := strconv.Atoi(splitString[0])
	check(err)
	num2, err := strconv.Atoi(splitString[1])
	check(err)
	return num1, num2
}

func part1(contents string) {
	lines := strings.Split(contents, "\n")
	score := 0
	for _, element := range lines {
		if element == "" {
			continue
		}
		assigments := strings.Split(element, ",")
		elf1begin, elf1end := splitTwoIntsByString(assigments[0], "-")
		elf2begin, elf2end := splitTwoIntsByString(assigments[1], "-")
		if doesRangeContainOther(elf1begin, elf1end, elf2begin, elf2end) {
			score = score + 1
		}
	}
	fmt.Println(score)
}

func part2(contents string) {
	lines := strings.Split(contents, "\n")
	score := 0
	for _, element := range lines {
		if element == "" {
			continue
		}
		assigments := strings.Split(element, ",")
		elf1begin, elf1end := splitTwoIntsByString(assigments[0], "-")
		elf2begin, elf2end := splitTwoIntsByString(assigments[1], "-")
		if doAnyRangesOverlap(elf1begin, elf1end, elf2begin, elf2end) {
			score = score + 1
		}
	}
	fmt.Println(score)
}

func main() {
	data := readFileAsString("input.txt")
	part1(data)
	part2(data)
}
