package util

import (
	"fmt"
	"strings"
)

type Action string
type GameOutcome string

const (
	Rock     Action = "Rock"
	Paper           = "Paper"
	Scissors        = "Scissors"

	Draw GameOutcome = "Draw"
	Lose GameOutcome = "Lose"
	Win  GameOutcome = "Win"
)

var identifierToOutcome = map[string]GameOutcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var identifierToAction = map[string]Action{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var pointsForAction = map[Action]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var actionWinsAgainst = map[Action]Action{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

func ExtractActionsFromLine(line string) (Action, Action, error) {
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

func ExtractActionAndGameOutcomeFromLine(line string) (Action, GameOutcome, error) {
	lineParts := strings.Split(line, " ")
	if len(lineParts) != 2 {
		return "", "", fmt.Errorf("invalid line, expected two characters split by space, got: %q", line)
	}

	actionPlayer1, foundAction := identifierToAction[lineParts[0]]
	outcome, foundOutcome := identifierToOutcome[lineParts[1]]
	if !(foundAction && foundOutcome) {
		return "", "", fmt.Errorf("unknown action in line: %q", line)
	}

	return actionPlayer1, outcome, nil
}

func GetPointsForAction(a Action) int {
	return pointsForAction[a]
}

func ActionWinsAgainst(a Action) Action {
	return actionWinsAgainst[a]
}

func ActionLosesAgainst(action Action) Action {
	for a, winsAgainst := range actionWinsAgainst {
		if winsAgainst == action {
			return a
		}
	}
	return ""
}
