package main

import (
	"log"

	"advent-of-code-2022/helpers"
)

func main() {
	scanner, file, err := helpers.GetInput("../input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

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
