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

func (p *Position) IsConnectedWith(p2 Position) bool {
	return p.equal(p2) || p.isDirectlyConnected(p2) || p.isDiagonalConnected(p2)
}

func (p *Position) isDirectlyConnected(p2 Position) bool {
	return (p.X == p2.X && intAbs(p.Y-p2.Y) == 1) || (p.Y == p2.Y && intAbs(p.X-p2.X) == 1)
}

func (p *Position) isDiagonalConnected(p2 Position) bool {
	return intAbs(p.X-p2.X) == 1 && intAbs(p.Y-p2.Y) == 1
}

func (p *Position) equal(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func intAbs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
