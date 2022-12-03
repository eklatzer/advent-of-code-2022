package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var itemPrioritySum = 0

	for scanner.Scan() {
		bytesPerLine := scanner.Bytes()

		var itemExistsInFirstCompartment = map[byte]struct{}{}

		for i, itemIdentifier := range bytesPerLine {
			if isItemOfFirstCompartment(i, len(bytesPerLine)) {
				itemExistsInFirstCompartment[itemIdentifier] = struct{}{}
			} else if _, existsInFirstCompartment := itemExistsInFirstCompartment[itemIdentifier]; existsInFirstCompartment {
				itemPrioritySum += int(getScoreForItem(itemIdentifier))
				break
			}
		}
	}
	log.Println(itemPrioritySum)
}

func isItemOfFirstCompartment(index, totalItems int) bool {
	return index < totalItems/2
}

func getScoreForItem(item byte) byte {
	if isUpper(item) {
		return item - 38
	}
	return item - 96
}

func isUpper(x byte) bool {
	return 'A' <= x && x <= 'Z'
}
