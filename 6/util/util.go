package util

import (
	"os"

	"advent-of-code-2022/helpers"
)

func GetInputLine() []byte {
	content, err := os.ReadFile(helpers.GetInputFilePath())
	if err != nil {
		panic(err)
	}

	return content
}

func GetNumberOfDifferentChars(chars []byte) int {
	var charSet = helpers.Set[byte]{}
	for _, v := range chars {
		charSet[v] = struct{}{}
	}
	return len(charSet)
}

func GetNValuesStartingFromIndex(in []byte, index, n int) []byte {
	if (index + n) >= len(in) {
		return in[index:]
	}
	return in[index:(index + n)]
}
