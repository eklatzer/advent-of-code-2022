package util

import (
	"fmt"
	"strings"
)

// Contains the possible values for actions and outcomes of Rock Paper Scissors games.
const (
	Rock     = "Rock"
	Paper    = "Paper"
	Scissors = "Scissors"

	Draw = "Draw"
	Lose = "Lose"
	Win  = "Win"
)

var identifierToOutcome = map[string]string{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var identifierToAction = map[string]string{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var pointsForAction = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var actionWinsAgainst = map[string]string{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

// ExtractActionsFromLine extracts the actions in a line, represented by two characters separated by a space.
func ExtractActionsFromLine(line string) (string, string, error) {
	return extractValues(line, identifierToAction, identifierToAction)
}

// ExtractActionAndGameOutcomeFromLine extracts the actions/expected outcome in a line, represented by two characters separated by a space.
func ExtractActionAndGameOutcomeFromLine(line string) (string, string, error) {
	return extractValues(line, identifierToAction, identifierToOutcome)
}

func extractValues(line string, mappingForFirstValue map[string]string, mappingForSecondValue map[string]string) (string, string, error) {
	lineParts := strings.Split(line, " ")
	if len(lineParts) != 2 {
		return "", "", fmt.Errorf("invalid line, expected two characters split by space, got: %q", line)
	}

	actionPlayer1, foundAction := mappingForFirstValue[lineParts[0]]
	outcome, foundOutcome := mappingForSecondValue[lineParts[1]]
	if !(foundAction && foundOutcome) {
		return "", "", fmt.Errorf("unknown action in line: %q", line)
	}

	return actionPlayer1, outcome, nil
}

// GetPointsForAction returns the points a player gets when using a specific action.
func GetPointsForAction(a string) int {
	return pointsForAction[a]
}

// ActionWinsAgainst returns the action the given action wins against.
func ActionWinsAgainst(a string) string {
	return actionWinsAgainst[a]
}

// ActionLosesAgainst returns the action the given action loses against.
func ActionLosesAgainst(currentAction string) string {
	for action, winsAgainst := range actionWinsAgainst {
		if winsAgainst == currentAction {
			return action
		}
	}
	return ""
}
