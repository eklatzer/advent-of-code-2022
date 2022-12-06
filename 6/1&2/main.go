package main

import (
	"flag"
	"log"

	"advent-of-code-2022/6/util"
)

var numberOfDistinctCharsForMarker int

func init() {
	flag.IntVar(&numberOfDistinctCharsForMarker, "n", 4, "Number of distinct chars in a row for a marker")
	flag.Parse()
}

func main() {
	line := util.GetInputLine()
	for i := 0; i < len(line); i++ {
		elements := util.GetNValuesStartingFromIndex(line, i, numberOfDistinctCharsForMarker)
		if util.GetNumberOfDifferentChars(elements) == numberOfDistinctCharsForMarker {
			log.Printf("index after start of marker: %d", (i + numberOfDistinctCharsForMarker))
			break
		}
	}
}
