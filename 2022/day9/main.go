package main

import (
	"fmt"
	"github.com/Erich-Reitz/commonGo"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Rope struct {
	items []Coordinate
}

func (rope *Rope) Head() *Coordinate {
	return &rope.items[0]
}

func (rope *Rope) MoveHead(direction string) {
	if direction == "U" {
		rope.Head().Y += 1
	} else if direction == "D" {
		rope.Head().Y -= 1
	} else if direction == "R" {
		rope.Head().X += 1
	} else if direction == "L" {
		rope.Head().X -= 1
	}

	rope.SimulateMovement()
}

func (rope *Rope) buildKnots(numKnots int) {
	rope.items = make([]Coordinate, numKnots)
	for i := 0; i < numKnots; i++ {
		rope.items[i] = Coordinate{0, 0}
	}
}

func (rope *Rope) SimulateMovementBetweenKnots(head, tail *Coordinate) {
	xDifference := advent.AbsInt(head.X - tail.X)
	yDifference := advent.AbsInt(head.Y - tail.Y)

	if xDifference > 1 || yDifference > 1 {
		if head.X-tail.X > 0 {
			tail.X += 1
		} else if head.X-tail.X < 0 {
			tail.X -= 1
		}
		if head.Y-tail.Y > 0 {
			tail.Y += 1
		} else if head.Y-tail.Y < 0 {
			tail.Y -= 1
		}
	}

}

func (rope *Rope) GetTail() *Coordinate {
	return &rope.items[len(rope.items)-1]
}

func (rope *Rope) SimulateMovement() {
	for i := 1; i < len(rope.items); i++ {
		rope.SimulateMovementBetweenKnots(&rope.items[i-1], &rope.items[i])
	}
}

func simulateRopeMovementOverDirections(lengthOfRope int, input string) {
	visitedCount := make(map[Coordinate]int)
	lines := strings.Split(input, "\n")
	rope := Rope{}
	rope.buildKnots(lengthOfRope)
	for _, line := range lines {
		var direction string
		var amount int
		fmt.Sscanf(line, "%s %d", &direction, &amount)
		for iteration := 0; iteration < amount; iteration++ {
			rope.MoveHead(direction)
			visitedCount[*rope.GetTail()] += 1
		}
	}
	fmt.Println(len(visitedCount))
}

func part1(input string) {
	simulateRopeMovementOverDirections(2, input)
}

func part2(input string) {
	simulateRopeMovementOverDirections(10, input)
}

func main() {
	data := advent.ReadFileAsString("input.txt")
	part1(data)
	part2(data)
}
