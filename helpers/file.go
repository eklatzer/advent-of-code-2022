package helpers

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var inputFilePath *string

func init() {
	inputFilePath = flag.String("file", "../input.txt", "Path to input file")
}

// GetInputFilePath returns the value of the flag -file.
func GetInputFilePath() string {
	return *inputFilePath
}

// GetInput returns the input file split by lines.
func GetInput(filepath string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, file, fmt.Errorf("failed to open file: %v", err)
	}
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	return s, file, nil
}
