package main

import (
	"fmt"
	"github.com/Erich-Reitz/commonGo"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func doesRangeContainOther(elf1begin, elf1end, elf2begin, elf2end int) bool {
	allMin := advent.MinOf(elf1begin, elf1end, elf2begin, elf2end)
	allMax := advent.MaxOf(elf1begin, elf1end, elf2begin, elf2end)

	return (allMin == elf1begin && allMax == elf1end) || (allMin == elf2begin && allMax == elf2end)
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

func part1(contents string) {
	lines := strings.Split(contents, "\n")
	score := 0
	for _, element := range lines {
		if element == "" {
			continue
		}
		assigments := strings.Split(element, ",")
		elf1begin, elf1end := advent.SplitTwoIntsByString(assigments[0], "-")
		elf2begin, elf2end := advent.SplitTwoIntsByString(assigments[1], "-")
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
		elf1begin, elf1end := advent.SplitTwoIntsByString(assigments[0], "-")
		elf2begin, elf2end := advent.SplitTwoIntsByString(assigments[1], "-")
		if doAnyRangesOverlap(elf1begin, elf1end, elf2begin, elf2end) {
			score = score + 1
		}
	}
	fmt.Println(score)
}

func main() {
	data := advent.ReadFileAsString("input.txt")
	part1(data)
	part2(data)
}
