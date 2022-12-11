package util

import (
	"os"
	"sort"
	"strings"

	"advent-of-code-2022/helpers"
)

type Monkey struct {
	Items     []int
	Operation string
	Test      Test
}

type Test struct {
	Divisor       int
	TargetIfTrue  int
	TargetIfFalse int
}

type InspectionCount map[int]int

func (c InspectionCount) GetSortedKeysByValue() []int {
	keys := make([]int, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return c[keys[i]] < c[keys[j]]
	})

	return keys
}

func ReadInput(path string) ([]Monkey, error) {
	filecontent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(filecontent), "\n\n")

	var monkey = []Monkey{}
	for _, part := range parts {
		lines := strings.Split(part, "\n")

		monkey = append(monkey, Monkey{
			Items:     extractItems(lines[1]),
			Operation: extractOperation(lines[2]),
			Test:      extractTest(lines[3:]),
		})
	}
	return monkey, nil
}

func RunOperation(worryLevel int, operation string) int {
	operand := "*"
	if strings.Contains(operation, "+") {
		operand = "+"
	}
	operationParts := strings.Split(operation, operand)
	firstValue := extractValue(operationParts[0], worryLevel)
	secondValue := extractValue(operationParts[1], worryLevel)
	if operand == "+" {
		return firstValue + secondValue
	}
	return firstValue * secondValue
}

func extractValue(in string, oldWorryLevel int) int {
	if strings.Contains(in, "old") {
		return oldWorryLevel
	}
	return trimSpaceToInt(in)
}

func extractItems(in string) []int {
	itemList := strings.Split(in, ":")
	parts := strings.Split(itemList[1], ",")

	var out = []int{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		out = append(out, helpers.ParseInt(part))
	}
	return out
}

func extractOperation(in string) string {
	operation := strings.Split(in, "=")[1]
	return strings.TrimSpace(operation)
}

func extractTest(in []string) Test {
	return Test{
		Divisor:       trimSpaceToInt(strings.Split(in[0], "by")[1]),
		TargetIfTrue:  trimSpaceToInt(strings.Split(in[1], "monkey")[1]),
		TargetIfFalse: trimSpaceToInt(strings.Split(in[2], "monkey")[1]),
	}
}

func trimSpaceToInt(in string) int {
	return helpers.ParseInt(strings.TrimSpace(in))
}
