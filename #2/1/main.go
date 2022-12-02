package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var identifierToAction = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var pointsForUsage = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var actionWinsAgainst = map[string]string{
	"Rock":     "Scissors",
	"Paper":    "Rock",
	"Scissors": "Paper",
}

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var score int

	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, " ")
		if len(lineParts) != 2 {
			log.Printf("invalid number of parts per lined split by space: %d\n", len(lineParts))
			continue
		}

		action0, found0 := identifierToAction[lineParts[0]]
		action1, found1 := identifierToAction[lineParts[1]]
		if !(found0 && found1) {
			log.Printf("unknown action in line: %s", line)
			continue
		}

		score += pointsForUsage[action1]

		if action0 == action1 {
			score += 3
		} else {
			currentActionWinsAgainst := actionWinsAgainst[action1]
			if currentActionWinsAgainst == action0 {
				score += 6
			}
		}
	}
	log.Printf("total score: %d\n", score)
}
