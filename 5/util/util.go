package util

import (
	"strconv"
	"strings"
)

// Ship represents a complete ship containing Stacks.s
type Ship []Stack

// Stack represents one stack of crates.
type Stack []byte

// ExtractShipAndGetEndIndex extracts the ship and returns the index of the first line after the ship.
func ExtractShipAndGetEndIndex(lines []string) (Ship, int) {
	var ship = Ship{}
	for lineIndex, line := range lines {
		for i := 1; i < len(line); i += 4 {
			var c = line[i]
			if isDigit(c) {
				return ship, lineIndex + 1
			}
			if !isSpace(c) {
				stackIndex := (i - 1) / 4
				ship.addElement(c, stackIndex)
			}
		}
	}
	return nil, 0
}

func isDigit(in byte) bool {
	return in >= '0' && in <= '9'
}

func isSpace(in byte) bool {
	return in == ' '
}

func (s *Ship) addElement(value byte, index int) {
	var initialSize = len(*s)
	for i := 0; i <= (index - initialSize); i++ {
		*s = append(*s, Stack{})
	}
	(*s)[index] = append((*s)[index], value)
}

type command struct {
	amount int
	from   int
	to     int
}

// ExecuteCommand executes a command in the format 'move 1 from 6 to 8'.
func (s Ship) ExecuteCommand(cmd command, reverseItems bool) {
	movedItems := append(Stack{}, s[cmd.from][0:cmd.amount]...)
	if reverseItems {
		movedItems.reverse()
	}
	s[cmd.from] = s[cmd.from][cmd.amount:]
	s[cmd.to] = append(movedItems, s[cmd.to]...)

}

func (st Stack) reverse() Stack {
	for i, j := 0, len(st)-1; i < j; i, j = i+1, j-1 {
		st[i], st[j] = st[j], st[i]
	}
	return st
}

func ParseCommand(commandLine string) command {
	if commandLine == "" {
		return command{}
	}
	parts := strings.Split(commandLine, " ")
	return command{
		amount: toInt(parts[1]),
		from:   toInt(parts[3]) - 1,
		to:     toInt(parts[5]) - 1,
	}
}

func toInt(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return val
}
