package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(filepath string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, file, fmt.Errorf("failed to open file: %v", err)
	}
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	return s, file, nil
}
