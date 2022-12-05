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

	ship, endIndex := util.ExtractShipAndGetEndIndex(lines)
	for _, move := range lines[endIndex:] {
		cmd := util.ParseCommand(move)
		ship.ExecuteCommand(cmd, false)
	}
	fmt.Println("Top Crates:")
	for _, stack := range ship {
		fmt.Printf("%c", stack[0])
	}
}
