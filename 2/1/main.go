package main

import (
	"flag"
	"log"

	"advent-of-code-2022/2/util"
	"advent-of-code-2022/helpers"
)

func main() {
	flag.Parse()

	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var score int

	for scanner.Scan() {
		actionPlayer1, actionPlayer2, err := util.ExtractActionsFromLine(scanner.Text())
		if err != nil {
			log.Println(err.Error())
			continue
		}

		score += util.GetPointsForAction(actionPlayer2)

		if actionPlayer1 == actionPlayer2 {
			score += 3
		} else {
			currentActionWinsAgainst := util.ActionWinsAgainst(actionPlayer2)
			if currentActionWinsAgainst == actionPlayer1 {
				score += 6
			}
		}
	}
	log.Printf("total score: %d", score)
}
