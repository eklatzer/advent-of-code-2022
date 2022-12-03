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
		actionPlayer1, expectedOutcome, err := util.ExtractActionAndGameOutcomeFromLine(scanner.Text())
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var actionPlayer2 string

		switch expectedOutcome {
		case util.Draw:
			score += 3
			actionPlayer2 = actionPlayer1
			break
		case util.Lose:
			actionPlayer2 = util.ActionWinsAgainst(actionPlayer1)
			break
		case util.Win:
			score += 6
			actionPlayer2 = util.ActionLosesAgainst(actionPlayer1)
		}
		score += util.GetPointsForAction(actionPlayer2)
	}
	log.Printf("total score: %d\n", score)
}
