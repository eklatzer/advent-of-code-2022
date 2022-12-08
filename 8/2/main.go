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

	var maxScore = 0
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			curScore := treeMap.GetScenicScore(i, j)
			if curScore > maxScore {
				maxScore = curScore
			}
		}
	}
	log.Printf("the maximum scenic score is: %d", maxScore)
}
