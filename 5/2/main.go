package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent-of-code-2022/5/util"
	"advent-of-code-2022/helpers"
)

func main() {
	// trying out different version of file-reading
	filecontent, err := os.ReadFile(helpers.GetInputFilePath())
	if err != nil {
		log.Fatalf("failed to read input file at %q: %v", helpers.GetInputFilePath(), err)
	}
	lines := strings.Split(string(filecontent), "\n")

	ship, moves := util.ExtractShipAndMoves(lines)
	for _, move := range moves {
		ship.ExecuteCommand(move, false)
	}
	fmt.Println("Top Crates:")
	for _, stack := range ship {
		fmt.Printf("%c", stack[0])
	}
}
