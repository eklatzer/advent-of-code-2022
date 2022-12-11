package main

import (
	"flag"
	"log"

	"advent-of-code-2022/11/util"
	"advent-of-code-2022/helpers"
)

var isPartOne bool
var rounds int

func init() {
	flag.BoolVar(&isPartOne, "part-one", true, "Whether part 1 (division of new worry level by 3) or part 2 (newLevel % least common multiple of all divisors) is executed")
	flag.IntVar(&rounds, "n", 20, "Number of rounds")
	flag.Parse()
}

// part 1: default values for flags
// part 2: -n=10000 -part-one=false

func main() {
	monkeys, err := util.ReadInput(helpers.GetInputFilePath())
	if err != nil {
		log.Fatalln(err)
	}

	var inspectionCount = util.InspectionCount{}

	var lcm = 1
	if !isPartOne {
		for _, m := range monkeys {
			lcm *= m.Test.Divisor
		}
	}

	for i := 0; i < rounds; i++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				inspectionCount[i]++
				newWorryLevel := util.RunOperation(item, monkey.Operation)

				if isPartOne {
					newWorryLevel = newWorryLevel / 3
				} else {
					newWorryLevel = newWorryLevel % lcm
				}

				targetMonkey := monkey.Test.TargetIfFalse
				if newWorryLevel%monkey.Test.Divisor == 0 {
					targetMonkey = monkey.Test.TargetIfTrue
				}

				monkeys[targetMonkey].Items = append(monkeys[targetMonkey].Items, newWorryLevel)
			}
			monkeys[i].Items = []int{}
		}
	}

	sortedKeys := inspectionCount.GetSortedKeysByValue()
	for i := len(sortedKeys) - 1; i >= 0; i-- {
		monkeyNumber := sortedKeys[i]
		log.Printf("monkey %d has inspected %d items", monkeyNumber, inspectionCount[monkeyNumber])
	}

	// level of monkey business for the two monkeys with the most inspections
	var levelOfMonkeyBusiness = inspectionCount[sortedKeys[len(sortedKeys)-1]] * inspectionCount[sortedKeys[len(sortedKeys)-2]]
	log.Printf("level of monkey business: %d", levelOfMonkeyBusiness)
}
