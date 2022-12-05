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

// ExecuteCommand executes the given command on the ship.
func (s Ship) ExecuteCommand(cmd Command, reverseItems bool) {
	movedItems := append(Stack{}, s[cmd.from][0:cmd.amount]...)
	if reverseItems {
		movedItems.reverse()
	}
	s[cmd.from] = s[cmd.from][cmd.amount:]
	s[cmd.to] = append(movedItems, s[cmd.to]...)

}

// Command is one command handling the movement of amount items from the index from to the index to.
type Command struct {
	amount int
	from   int
	to     int
}

// ParseCommand parses a command in the format 'move 1 from 6 to 8'.
func ParseCommand(commandLine string) Command {
	if commandLine == "" {
		return Command{}
	}
	parts := strings.Split(commandLine, " ")
	return Command{
		amount: toInt(parts[1]),
		from:   toInt(parts[3]) - 1,
		to:     toInt(parts[5]) - 1,
	}
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

func (st Stack) reverse() Stack {
	for i, j := 0, len(st)-1; i < j; i, j = i+1, j-1 {
		st[i], st[j] = st[j], st[i]
	}
	return st
}

func toInt(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return val
}
