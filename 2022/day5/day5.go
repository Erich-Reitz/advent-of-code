package main

import (
	"fmt"
	"github.com/Erich-Reitz/commonGo"
	"reflect"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}



func parseCargoInitialStateLine(line string) map[int]string {
	result_map := make(map[int]string)
	for index, char := range line {
		if unicode.IsUpper(char) {
			result_map[((index+2)/4)+1] = string(char)
		}
	}
	return result_map
}

func getMoveInstructions(lines []string) []string {
	moveInstructions := make([]string, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "move") {
			moveInstructions = append(moveInstructions, line)
		}
	}
	return moveInstructions
}

func ReverseSlice(s []string) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func part1(cargo map[int][]string, moveInstructions []string) {
	for _, line := range moveInstructions {
		var amount, indexFrom, indexTo int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &indexFrom, &indexTo)
		check(err)
		toMove := cargo[indexFrom][len(cargo[indexFrom])-amount:]
		cargo[indexFrom] = cargo[indexFrom][:len(cargo[indexFrom])-amount]
		cargo[indexTo] = append(cargo[indexTo], toMove...)
	}
	printLastEntryOfEachCargoStack(cargo)
}

func printLastEntryOfEachCargoStack(cargo map[int][]string) {
	for i := 1; i <= 9; i++ {
		values := cargo[i]
		if len(values) != 0 {
			fmt.Print(cargo[i][len(cargo[i])-1])
		}
	}
}

func getInitialCargoState(lines []string) map[int][]string {
	cargo := make(map[int][]string)
	for _, line := range lines {
		if line == "\n" {
			break
		} else {
			parsedLineResult := parseCargoInitialStateLine(line)
			for index, value := range parsedLineResult {
				cargo[index] = append([]string{value}, cargo[index]...)
			}
		}
	}
	return cargo
}

func part2(cargo map[int][]string, moveInstructions []string) {
	for _, line := range moveInstructions {
		var amount, indexFrom, indexTo int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &indexFrom, &indexTo)
		check(err)
		toMove := cargo[indexFrom][len(cargo[indexFrom])-amount:]
		cargo[indexFrom] = cargo[indexFrom][:len(cargo[indexFrom])-amount]
		ReverseSlice(toMove)
		cargo[indexTo] = append(cargo[indexTo], toMove...)
	}
	printLastEntryOfEachCargoStack(cargo)
}

func main() {
	data := advent.ReadFileAsString("input.txt")
	lines := strings.Split(data, "\n")
	cargo := getInitialCargoState(lines)
	moveInstructions := getMoveInstructions(lines)
	part1(cargo, moveInstructions)
	fmt.Println()
	// part2(cargo, moveInstructions)
	// fmt.Println()
}
