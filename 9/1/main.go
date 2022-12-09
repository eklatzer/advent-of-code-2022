package main

import (
	"flag"
	"log"

	"advent-of-code-2022/9/util"
	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

func main() {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	headPosition := &util.Position{}
	tailPosition := &util.Position{}

	field := helpers.Set[util.Position]{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		cmd := util.ExtractCommand(line)

		for ; cmd.Distance > 0; cmd.Distance-- {
			switch cmd.Direction {
			case util.Up:
				headPosition.Y++
			case util.Down:
				headPosition.Y--
			case util.Right:
				headPosition.X++
			case util.Left:
				headPosition.X--
			}
			tailPosition.Follow(headPosition)
			field[*tailPosition] = struct{}{}
		}
	}

	log.Printf("%d positions are visited by the tail at least once", len(field))
}
