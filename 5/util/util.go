package util

import (
	"strconv"
	"strings"
)

// Ship represents a complete ship containing Stacks.s
type Ship []Stack

// Stack represents one stack of crates.
type Stack []byte

// ExtractShipAndMoves extracts the ship at the start and the commands.
func ExtractShipAndMoves(lines []string) (Ship, []string) {
	var ship = Ship{}
	for lineIndex, line := range lines {
		for i := 1; i < len(line); i += 4 {
			var c = line[i]
			if isDigit(c) {
				return ship, lines[(lineIndex + 2):]
			}
			if !isSpace(c) {
				stackIndex := (i - 1) / 4
				ship.addElement(c, stackIndex)
			}
		}
	}
	return nil, nil
}

func isDigit(in byte) bool {
	return in >= 48 && in <= 57
}

func isSpace(in byte) bool {
	return in == 32
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
func (s Ship) ExecuteCommand(commandLine string, reverseItems bool) {
	if commandLine != "" {
		cmd := getCommand(commandLine)
		movedItems := shallowCopy(s[cmd.from][0:cmd.amount])
		if reverseItems {
			movedItems.reverse()
		}
		s[cmd.from] = s[cmd.from][cmd.amount:]
		s[cmd.to] = append(movedItems, s[cmd.to]...)
	}
}

func shallowCopy(in Stack) Stack {
	newArray := make(Stack, len(in))
	copy(newArray, in)
	return newArray
}

func (st Stack) reverse() Stack {
	for i, j := 0, len(st)-1; i < j; i, j = i+1, j-1 {
		st[i], st[j] = st[j], st[i]
	}
	return st
}

func getCommand(commandLine string) command {
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
