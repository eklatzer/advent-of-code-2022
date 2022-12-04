package main

import (
	"log"

	"advent-of-code-2022/3/util"
	"advent-of-code-2022/helpers"
)

func main() {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var itemPrioritySum = 0

	var itemsForElves = []util.ByteSet{{}, {}}
	var elveNumber = 1

	for scanner.Scan() {
		bytesPerLine := scanner.Bytes()

		if elveNumber == 1 || elveNumber == 2 {
			for _, itemIdentifier := range bytesPerLine {
				itemsForElves[elveNumber-1][itemIdentifier] = struct{}{}
			}
		} else {
			for _, itemIdentifier := range bytesPerLine {
				if hasItem(itemIdentifier, itemsForElves[0]) && hasItem(itemIdentifier, itemsForElves[1]) {
					itemPrioritySum += int(getScoreForItem(itemIdentifier))
					itemsForElves = []util.ByteSet{{}, {}}
					elveNumber = 0
					break
				}
			}
		}

		elveNumber++
	}
	log.Println(itemPrioritySum)
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

func hasItem(itemIdentifer byte, items util.ByteSet) bool {
	_, exists := items[itemIdentifer]
	return exists
}
