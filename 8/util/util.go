package util

import (
	"log"
	"os"
	"strings"

	"advent-of-code-2022/helpers"
)

// TreeMap represents a two-dimensional map of the height of trees in the range from 0-9
type TreeMap [][]byte

// NewTreeMapFromFile reads the TreeMap from the given file
func NewTreeMapFromFile(filepath string) TreeMap {
	filecontent, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("failed to read input file at %q: %v", helpers.GetInputFilePath(), err)
	}

	lines := strings.Split(string(filecontent), "\n")

	var treeMap = make(TreeMap, len(lines[0]))

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		if line != "" {
			treeMap[lineIndex] = make([]byte, len(line))
			for colIndex := 0; colIndex < len(line); colIndex++ {
				treeMap[lineIndex][colIndex] = line[colIndex] - '0'
			}
		}
	}
	return treeMap
}

// IsVisible checks if the tree in the row/col is visible from the outside of the forest.
func (t *TreeMap) IsVisisble(row, col int) bool {
	value := (*t)[row][col]
	colValues := t.getElementsForCol(col)

	if row == 0 || col == 0 || col >= (len(*t)-1) || row >= (len((*t)[0])-1) { // elments at the edge
		return true
	}

	return treeIsVisible(value, (*t)[row][0:col]) || treeIsVisible(value, (*t)[row][(col+1):]) || treeIsVisible(value, colValues[0:row]) || treeIsVisible(value, colValues[(row+1):])
}

// GetScenicScore returns the scenic score of the tree at the requested position.
func (t *TreeMap) GetScenicScore(row, col int) int {
	sizeOfCurrentTree := (*t)[row][col]
	var score = 1

	leftHigherTreeIndex := getIndexOfFirstHigherValueFromEnd(sizeOfCurrentTree, (*t)[row][0:col])
	score *= getScenicScoreInFrontOfTree(col, leftHigherTreeIndex)

	rightHigherIndexTree := getIndexOfFirstHigherValue(sizeOfCurrentTree, (*t)[row][(col+1):])
	score *= getScenicScoreBehindTree(col, rightHigherIndexTree, len((*t)[row]))

	colValues := t.getElementsForCol(col)
	topHigherTreeIndex := getIndexOfFirstHigherValueFromEnd(sizeOfCurrentTree, colValues[0:row])
	score *= getScenicScoreInFrontOfTree(row, topHigherTreeIndex)

	bottomHigherTreeIndex := getIndexOfFirstHigherValue(sizeOfCurrentTree, colValues[(row+1):])
	score *= getScenicScoreBehindTree(row, bottomHigherTreeIndex, len(colValues))

	return score
}

func getScenicScoreInFrontOfTree(position, indexOfHigherTree int) int {
	if indexOfHigherTree == -1 {
		return position
	}
	return (position - indexOfHigherTree)
}

func getScenicScoreBehindTree(position, indexOfHigherTree, numberOfElementsBehindTree int) int {
	if indexOfHigherTree == -1 {
		return (numberOfElementsBehindTree - position - 1)
	}
	return indexOfHigherTree + 1
}

func (t *TreeMap) getElementsForCol(colIndex int) []byte {
	var values []byte
	for i := 0; i < len(*t); i++ {
		values = append(values, (*t)[i][colIndex])
	}
	return values
}

func treeIsVisible(val byte, values []byte) bool {
	return getIndexOfFirstHigherValue(val, values) == -1
}

func getIndexOfFirstHigherValue(val byte, values []byte) int {
	for index, v := range values {
		if v >= val {
			return index
		}
	}
	return -1
}

func getIndexOfFirstHigherValueFromEnd(val byte, values []byte) int {
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] >= val {
			return i
		}
	}
	return -1
}
