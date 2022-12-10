package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

func main() {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var cycle = 0
	var registerValue = 1
	var sumOfSignificantSignalStrenghts = 0
	var screen = ""

	for scanner.Scan() {
		command := scanner.Text()
		addedValue := 0
		cyclesForCommand := 1
		if strings.HasPrefix(command, "addx") {
			commandParts := strings.Split(command, " ")
			addedValue = helpers.ParseInt(commandParts[1])
			cyclesForCommand = 2
		}

		for ; cyclesForCommand > 0; cyclesForCommand-- {
			if (cycle)%40 == 0 {
				screen += "\n"
			}
			if spriteIsDrawn(registerValue, cycle%40) {
				screen += "#"
			} else {
				screen += "."
			}

			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				sumOfSignificantSignalStrenghts += calcSignalStrength(registerValue, cycle)
			}
		}
		registerValue += addedValue
	}

	fmt.Println(screen)
	fmt.Println("---------------------")
	fmt.Printf("sum of significant signal strenghts: %d\n", sumOfSignificantSignalStrenghts)
}

func calcSignalStrength(registerValue, cycle int) int {
	return registerValue * cycle
}

func spriteIsDrawn(spritePosition, currentDrawnPosition int) bool {
	return helpers.Abs(spritePosition-currentDrawnPosition) <= 1
}
