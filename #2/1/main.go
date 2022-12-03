package main

import (
	"log"
	"strings"

	"advent-of-code-2022/helpers"
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
	scanner, file, err := helpers.GetInput("../input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

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
