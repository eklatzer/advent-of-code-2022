package util

import (
	"strings"

	"advent-of-code-2022/helpers"
)

type Direction string

const (
	Right Direction = "R"
	Left            = "L"
	Up              = "U"
	Down            = "D"
)

type Command struct {
	Distance  int
	Direction Direction
}

func ExtractCommand(line string) Command {
	parts := strings.Split(line, " ")
	return Command{
		Direction: Direction(parts[0]),
		Distance:  helpers.ParseInt(parts[1]),
	}
}

type Position struct {
	X int
	Y int
}

func (tail *Position) Follow(head Position) {
	deltaX := head.X - tail.X
	deltaY := head.Y - tail.Y

	if intAbs(deltaX) > 1 || intAbs(deltaY) > 1 {
		tail.X = tail.X + deltaToNeededSteps(deltaX)
		tail.Y = tail.Y + deltaToNeededSteps(deltaY)
	}
}

func deltaToNeededSteps(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func intAbs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
