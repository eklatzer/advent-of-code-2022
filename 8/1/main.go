package main

import (
	"flag"
	"log"

	"advent-of-code-2022/8/util"
	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

func main() {
	treeMap := util.NewTreeMapFromFile(helpers.GetInputFilePath())

	var visibleCount = 0
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			if treeMap.IsVisisble(i, j) {
				visibleCount++
			}
		}
	}
	log.Printf("nubmer of trees visible from the outside: %d", visibleCount)
}
