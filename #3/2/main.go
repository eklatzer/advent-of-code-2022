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

	var itemsForElves = make([]map[byte]struct{}, 2)
	for i := range itemsForElves {
		itemsForElves[i] = map[byte]struct{}{}
	}
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
					itemsForElves = make([]map[byte]struct{}, 2)
					for i := range itemsForElves {
						itemsForElves[i] = map[byte]struct{}{}
					}
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

func hasItem(itemIdentifer byte, items map[byte]struct{}) bool {
	_, exists := items[itemIdentifer]
	return exists
}
