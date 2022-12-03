package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"advent-of-code-2022/helpers"
)

type action string

const (
	rock     action = "Rock"
	paper           = "Paper"
	scissors        = "Scissors"
)

var identifierToAction = map[string]action{
	"X": rock,
	"Y": paper,
	"Z": scissors,
	"A": rock,
	"B": paper,
	"C": scissors,
}

var pointsForAction = map[action]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var actionWinsAgainst = map[action]action{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

func main() {
	flag.Parse()

	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var score int

	for scanner.Scan() {

		actionPlayer1, actionPlayer2, err := extractActionsFromLine(scanner.Text())
		if err != nil {
			log.Println(err.Error())
			continue
		}

		score += pointsForAction[actionPlayer2]

		if actionPlayer1 == actionPlayer2 {
			score += 3
		} else {
			currentActionWinsAgainst := actionWinsAgainst[actionPlayer2]
			if currentActionWinsAgainst == actionPlayer1 {
				score += 6
			}
		}
	}
	log.Printf("total score: %d", score)
}

func extractActionsFromLine(line string) (action, action, error) {
	lineParts := strings.Split(line, " ")
	if len(lineParts) != 2 {
		return "", "", fmt.Errorf("invalid line, expected two characters split by space, got: %q", line)
	}

	actionPlayer1, foundAction1 := identifierToAction[lineParts[0]]
	actionPlayer2, foundAction2 := identifierToAction[lineParts[1]]
	if !(foundAction1 && foundAction2) {
		return "", "", fmt.Errorf("unknown action in line: %q", line)
	}

	return actionPlayer1, actionPlayer2, nil
}
