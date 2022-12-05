package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"advent-of-code-2022/5/util"
	"advent-of-code-2022/helpers"
)

func main() {
	// trying out different version of file-reading
	filecontent, err := ioutil.ReadFile(helpers.GetInputFilePath())
	if err != nil {
		log.Fatalf("failed to read input file at %q: %v", helpers.GetInputFilePath(), err)
	}
	lines := strings.Split(string(filecontent), "\n")

	ship, moves := util.ExtractShipAndMoves(lines)
	for _, move := range moves {
		ship.ExecuteCommand(move, true)
	}
	fmt.Println("Top Crates:")
	for _, stack := range ship {
		fmt.Printf("%c", stack[0])
	}
}
